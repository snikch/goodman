// Include this code in your bootstrap code
// it receives messages from dredd and triggers the callbacks you registered
package hooks

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	trans "github.com/snikch/goodman/transaction"
)

type Server struct {
	// TODO: stop exposing this, have the user only use server.Close()
	// kept around to avoid breaking changes
	Listener net.Listener
}

func NewServerWithPortAndError(run RunnerRPC, port int) (*Server, error) {
	if port == 0 {
		return nil, errors.New("hook server must be provided with non-0 port")
	}

	serv := rpc.NewServer()
	// Publishes in the server the exported methods on our `run` interface
	serv.Register(run)
	// register these handlers in the default mux handler
	// calling http.Serve later expose these
	serv.HandleHTTP("/", "/debug")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("listen error: %s", err.Error())
	}

	server := Server{
		Listener: l,
	}
	return &server, nil
}

var port int

// Use the globally scoped port (which is parsed in init)
func NewServerWithError(run RunnerRPC) (*Server, error) {
	server, err := NewServerWithPortAndError(run, port)
	if err != nil {
		return nil, fmt.Errorf("Creating hook server: %s", err.Error())
	}
	return server, nil
}

func init() {
	flag.IntVar(&port, "port", 0, "The port that the hooks server will run on")
	flag.Parse()
}

// Legacy, compliant with existing usage
func NewServer(run RunnerRPC) *Server {
	server, err := NewServerWithError(run)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return server
}

// Legacy, compliant with existing usage
func (s *Server) Serve() {
	http.Serve(s.Listener, nil)
}

func (s *Server) ServeWithError() error {
	// Listen on the tcp connection we made on boot
	// serve handlers from the _default mux handler_
	return http.Serve(s.Listener, nil)
}

func (s *Server) Close() error {
	if err := s.Listener.Close(); err != nil {
		return fmt.Errorf("Closing listener: %s", err.Error())
	}
	return nil
}

type RunnerRPC interface {
	RunBeforeAll(args []*trans.Transaction, reply *[]*trans.Transaction) error
	RunBeforeEach(args trans.Transaction, reply *trans.Transaction) error
	RunBefore(args trans.Transaction, reply *trans.Transaction) error
	RunBeforeEachValidation(args trans.Transaction, reply *trans.Transaction) error
	RunBeforeValidation(args trans.Transaction, reply *trans.Transaction) error
	RunAfter(args trans.Transaction, reply *trans.Transaction) error
	RunAfterEach(args trans.Transaction, reply *trans.Transaction) error
	RunAfterAll(args []*trans.Transaction, reply *[]*trans.Transaction) error
}
