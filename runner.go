package goodman

import (
	"fmt"
	"net/rpc"
	"os/exec"

	"github.com/snikch/goodman/transaction"
)

func NewRunner(rpcService string, port int, cmd *exec.Cmd) (*Run, error) {
	client, err := rpc.DialHTTPPath("tcp", fmt.Sprintf(":%d", port), "/")
	if err != nil {
		return nil, fmt.Errorf("dialing tcp server: %s", err.Error())
	}

	runner := Run{
		cmd:        cmd,
		client:     client,
		rpcService: rpcService,
	}
	return &runner, nil
}

type Run struct {
	cmd        *exec.Cmd
	client     *rpc.Client
	rpcService string
}

func (r *Run) RunBeforeAll(t *[]*transaction.Transaction) {
	var reply []*transaction.Transaction
	err := r.client.Call(r.rpcService+".RunBeforeAll", t, &reply)

	if err != nil {
		panic("RPC client threw error " + err.Error())
	}
	*t = reply
}

func (r *Run) RunBeforeEach(t *transaction.Transaction) {
	var reply transaction.Transaction
	err := r.client.Call(r.rpcService+".RunBeforeEach", *t, &reply)

	if err != nil {
		panic("RPC client threw error " + err.Error())
	}
	*t = reply
}

func (r *Run) RunBefore(t *transaction.Transaction) {
	var reply transaction.Transaction
	err := r.client.Call(r.rpcService+".RunBefore", *t, &reply)

	if err != nil {
		panic("RPC client threw error " + err.Error())
	}
	*t = reply
}

func (r *Run) RunBeforeEachValidation(t *transaction.Transaction) {
	var reply transaction.Transaction
	err := r.client.Call(r.rpcService+".RunBeforeEachValidation", *t, &reply)

	if err != nil {
		panic("RPC client threw error " + err.Error())
	}
	*t = reply
}

func (r *Run) RunBeforeValidation(t *transaction.Transaction) {
	var reply transaction.Transaction
	err := r.client.Call(r.rpcService+".RunBeforeValidation", *t, &reply)

	if err != nil {
		panic("RPC client threw error " + err.Error())
	}
	*t = reply
}

func (r *Run) RunAfterAll(t *[]*transaction.Transaction) {
	var reply []*transaction.Transaction
	err := r.client.Call(r.rpcService+".RunAfterAll", t, &reply)

	if err != nil {
		panic("RPC client threw error " + err.Error())
	}
	*t = reply
}

func (r *Run) RunAfterEach(t *transaction.Transaction) {
	var reply transaction.Transaction
	err := r.client.Call(r.rpcService+".RunAfterEach", *t, &reply)

	if err != nil {
		panic("RPC client threw error " + err.Error())
	}
	*t = reply
}

func (r *Run) RunAfter(t *transaction.Transaction) {
	var reply transaction.Transaction
	err := r.client.Call(r.rpcService+".RunAfter", *t, &reply)

	if err != nil {
		panic("RPC client threw error " + err.Error())
	}
	*t = reply
}

func (r *Run) Close() (err error) {
	// Kill the underlying hook binary
	if cmdErr := r.cmd.Process.Kill(); cmdErr != nil {
		// What is dead may never die
		// TODO: this is a pretty bad idea, we should robustly detect if the binary is still running
		if cmdErr.Error() != "os: process already finished" {
			err = fmt.Errorf("Killing cmd: %s", cmdErr.Error())
		}
	}
	// terminate our connection listening to it
	if clientErr := r.client.Close(); clientErr != nil {
		err = fmt.Errorf("RPC client on Close() " + clientErr.Error())
	}
	return err
}

type Runner interface {
	RunBeforeAll(t *[]*transaction.Transaction)
	RunBeforeEach(t *transaction.Transaction)
	RunBefore(t *transaction.Transaction)
	RunBeforeEachValidation(t *transaction.Transaction)
	RunBeforeValidation(t *transaction.Transaction)
	RunAfterAll(t *[]*transaction.Transaction)
	RunAfterEach(t *transaction.Transaction)
	RunAfter(t *transaction.Transaction)
	Close() error
}

type DummyRunner struct{}

func (r *DummyRunner) RunBeforeAll(t *[]*transaction.Transaction) {}

func (r *DummyRunner) RunBeforeEach(t *transaction.Transaction) {}

func (r *DummyRunner) RunBefore(t *transaction.Transaction) {}

func (r *DummyRunner) RunBeforeEachValidation(t *transaction.Transaction) {}

func (r *DummyRunner) RunBeforeValidation(t *transaction.Transaction) {}

func (r *DummyRunner) RunAfterAll(t *[]*transaction.Transaction) {}

func (r *DummyRunner) RunAfterEach(t *transaction.Transaction) {}

func (r *DummyRunner) RunAfter(t *transaction.Transaction) {}

func (r *DummyRunner) Close() error {
	return nil
}
