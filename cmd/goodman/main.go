// The package that the dredd cli calls
// When you run `dredd` from the cli dredd creates a new process and runs this binary with
// a list of hooks to instrument.
// This (goodman) binary brings up an rpc server to recieve commands from dredd
// This (goodman) binary the relays this across the hooks specified on boot
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/snikch/goodman"
)

const (
	// The default port for communicating with the dredd cli
	defaultServerPort = 61321
	// How long to wait for a hook binary to begin responding
	hookServerWait    = 100 * time.Millisecond
	hookServerRetries = 5
)

func closeHookRunners(runners []goodman.Runner) []error {
	errs := []error{}
	for _, runner := range runners {
		if err := runner.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// Determine if an err represents a refused connection
// https://github.com/snikch/goodman/pull/23/files
func isConnectionRefusedError(err error) bool {
	if noerr, ok := err.(*net.OpError); ok {
		if scerr, ok := noerr.Err.(*os.SyscallError); ok {
			if scerr.Err == syscall.ECONNREFUSED {
				return true
			}
		}
	}
	return false
}

func createHookRunners(hookPaths []string, errChan chan error, startingPort int) ([]goodman.Runner, error) {
	runners := make([]goodman.Runner, len(hookPaths))
	// For each hook specified by the user, call it and bring up a hook runner
	// to make calls to it
	for i, path := range hookPaths {
		// Each hook should communicate on a different port, use a "block" of ports
		port := startingPort + i
		cmd := exec.Command(path, fmt.Sprintf("-port=%d", port))
		// Propogate messages from the cmd into the stdout/err
		// H.C: is this threadsafe?
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// don't block on the hook being called
		go func() {
			// Sniffed in logging output by tests to assert it actually happened
			log.Printf("Starting hooks server in go routine")
			if err := cmd.Run(); err != nil {
				errChan <- fmt.Errorf("hook server on port `%s`: %s", port, err.Error())
			}
		}()
		// The server may not immediatly return, give it a few attempts
		// TODO: investigate cmd.Wait() to shortcut this
		for retries := hookServerRetries; retries > 0; retries-- {
			// Must sleep so go routine running hooks server has chance to startup
			time.Sleep(hookServerWait)
			// Bring up the runner that will call out to the hook
			runner, err := goodman.NewRunner("Hooks", port, cmd)
			if err != nil {
				// Connection refused errors can be retried
				if isConnectionRefusedError(err) {
					continue
				}
				// Any other error cannot
				return runners, fmt.Errorf("creating runner: %s", err.Error())
			}
			runners[i] = runner
			break
		}
	}

	return runners, nil
}

func main() {
	port := flag.Int("port", defaultServerPort, "The port that the dredd callback server will run on")
	flag.Parse()
	if port == nil {
		log.Fatal("must be provided with port")
	}

	args := os.Args
	// Arguments are the hook binaries to run
	hookPaths := args[1:len(args)]
	// Each hook can report an error, as can the server itself (+1)
	errChan := make(chan error, len(hookPaths)+1)
	defer close(errChan)

	// Due to legacy reasons dummy run if not provided any real hooks
	runners := []goodman.Runner{&goodman.DummyRunner{}}
	var err error
	if len(hookPaths) > 0 {
		// Begin the ports to communicate with our runners 1 after the port used to communicate
		// with dredd
		startingPort := *port + 1
		runners, err = createHookRunners(hookPaths, errChan, startingPort)
	}

	defer func() {
		if errs := closeHookRunners(runners); len(errs) != 0 {
			closingErrs := make([]string, len(errs))
			for i, err := range errs {
				closingErrs[i] = err.Error()
			}
			log.Fatalf("closing hook runners: %s", strings.Join(closingErrs, "\n"))
		}
	}()

	if err != nil {
		log.Fatalf("creating hook runners: %s", err.Error())
	}

	// Bring up a server that will communicate with the main dredd runner
	// (on the port specified by dredd)
	var (
		server *goodman.Server
	)
	// Server blocks so run in goroutine
	go func() {
		server, err = goodman.NewServer(runners, *port)
		if err != nil {
			log.Fatalf("creating server: %s", err.Error())
		}

		if err := server.Run(); err != nil {
			errChan <- fmt.Errorf("running server: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		// If the user bails out, attempt to shut down gracefully
		case <-c:
			server.Close()
			closeHookRunners(runners)
			os.Exit(0)
		// If something broke, let the user know and bail out
		case err := <-errChan:
			server.Close()
			closeHookRunners(runners)
			log.Fatalf("encountered an unrecoverable error: %s", err.Error())
			os.Exit(0)
		}
	}
}
