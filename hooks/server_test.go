package hooks

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"

	r "github.com/snikch/goodman/rpc"
	trans "github.com/snikch/goodman/transaction"
)

var run = r.DummyRunner{}
var port = 61322
var addr = fmt.Sprintf(":%d", port)

func TestServerRPC(t *testing.T) {
	server := NewServer(&run, port)
	go func() {
		server.Serve()
		defer server.Listener.Close()
	}()

	client, err := rpc.DialHTTPPath("tcp", addr, "/")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	testCases := []struct {
		Method string
		args   interface{}
		reply  interface{}
		// Pointer needed so that when value is accessed it will
		// be real value and not a copy.
		notCalled *bool
	}{
		{
			Method:    "RunBeforeEach",
			args:      trans.Transaction{},
			reply:     trans.Transaction{},
			notCalled: &run.RunBeforeEachCalled,
		},
		{
			Method:    "RunBefore",
			args:      trans.Transaction{},
			reply:     trans.Transaction{},
			notCalled: &run.RunBeforeCalled,
		},
		{
			Method:    "RunBeforeValidation",
			args:      trans.Transaction{},
			reply:     trans.Transaction{},
			notCalled: &run.RunBeforeValidationCalled,
		},
		{
			Method:    "RunBeforeEachValidation",
			args:      trans.Transaction{},
			reply:     trans.Transaction{},
			notCalled: &run.RunBeforeEachValidationCalled,
		},
		{
			Method:    "RunAfterEach",
			args:      trans.Transaction{},
			reply:     trans.Transaction{},
			notCalled: &run.RunAfterEachCalled,
		},
		{
			Method:    "RunAfter",
			args:      trans.Transaction{},
			reply:     trans.Transaction{},
			notCalled: &run.RunAfterCalled,
		},
	}

	for _, test := range testCases {
		args := test.args.(trans.Transaction)
		reply := test.reply.(trans.Transaction)
		method := test.Method
		notCalled := test.notCalled
		err = client.Call("DummyRunner."+method, args, &reply)
		if err != nil {
			t.Errorf("rpc client failed to connect to server: %s", err.Error())
		}

		if !*notCalled {
			t.Errorf("RPC method %s was never invoked", method)
		}
	}

	// Testing for RunBeforeAll and RunAfter All
	var allReply []*trans.Transaction
	allCases := []struct {
		Method string
		args   []*trans.Transaction
		reply  []*trans.Transaction
		// Pointer needed so that when value is accessed it will
		// be real value and not a copy.
		notCalled *bool
	}{
		{
			Method:    "RunBeforeAll",
			args:      []*trans.Transaction{&trans.Transaction{}},
			reply:     allReply,
			notCalled: &run.RunBeforeAllCalled,
		},
		{
			Method:    "RunAfterAll",
			args:      []*trans.Transaction{&trans.Transaction{}},
			reply:     allReply,
			notCalled: &run.RunAfterAllCalled,
		},
	}

	for _, test := range allCases {
		args := test.args
		reply := test.reply
		method := test.Method
		notCalled := test.notCalled
		fmt.Println("Running test for " + method)
		err = client.Call("DummyRunner."+method, args, &reply)
		if err != nil {
			t.Errorf("rpc client failed to connect to server: %s", err.Error())
		}

		if !*notCalled {
			t.Errorf("RPC method %s was never invoked", method)
		}
	}
}
