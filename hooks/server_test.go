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
var serv = NewServer(&run, port)

func TestServerRunBeforeAllRPC(t *testing.T) {
	// client, err := rpc.DialHTTPPath("tcp", addr, "/")
	// if err != nil {
	// 	log.Fatal("dialing:", err)
	// }
	// tss := trans.Transaction{}
	// args := []*trans.Transaction{
	// 	&tss,
	// }
	// var reply []*trans.Transaction
	// err = client.Call("DummyRunner.RunBeforeAll", args, &reply)

	// if err != nil {
	// 	t.Errorf("rpc client failed to connect to server: %s", err.Error())
	// }

	// // Verify that method was invoked
	// if !run.RunBeforeAllCalled {
	// 	t.Errorf("RunBeforeAll was never invoked")
	// }
}

func TestServerRunBeforeEachRPC(t *testing.T) {
	client, err := rpc.DialHTTPPath("tcp", addr, "/")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := trans.Transaction{}
	var reply trans.Transaction
	err = client.Call("DummyRunner.RunBeforeEach", args, &reply)

	if err != nil {
		t.Errorf("rpc client failed to connect to server: %s", err.Error())
	}

	// Verify that method was invoked
	if !run.RunBeforeEachCalled {
		t.Errorf("RunBeforeEach was never invoked")
	}
}

func TestServerRunBeforeRPC(t *testing.T) {
	client, err := rpc.DialHTTPPath("tcp", addr, "/")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := trans.Transaction{}
	var reply trans.Transaction
	err = client.Call("DummyRunner.RunBefore", args, &reply)

	if err != nil {
		t.Errorf("rpc client failed to connect to server: %s", err.Error())
	}

	// Verify that method was invoked
	if !run.RunBeforeCalled {
		t.Errorf("RunBefore was never invoked")
	}
}

func TestServerRunBeforeEachValidationRPC(t *testing.T) {
	client, err := rpc.DialHTTPPath("tcp", addr, "/")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := trans.Transaction{}
	var reply trans.Transaction
	err = client.Call("DummyRunner.RunBeforeEachValidation", args, &reply)

	if err != nil {
		t.Errorf("rpc client failed to connect to server: %s", err.Error())
	}

	// Verify that method was invoked
	if !run.RunBeforeEachValidationCalled {
		t.Errorf("RunBeforeEachValidation was never invoked")
	}
}

func TestServerRunBeforeValidationRPC(t *testing.T) {
	client, err := rpc.DialHTTPPath("tcp", addr, "/")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := trans.Transaction{}
	var reply trans.Transaction
	err = client.Call("DummyRunner.RunBeforeValidation", args, &reply)

	if err != nil {
		t.Errorf("rpc client failed to connect to server: %s", err.Error())
	}

	// Verify that method was invoked
	if !run.RunBeforeValidationCalled {
		t.Errorf("RunBeforeValidation was never invoked")
	}
}

func TestServerRunAfterRPC(t *testing.T) {
	client, err := rpc.DialHTTPPath("tcp", addr, "/")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := trans.Transaction{}
	var reply trans.Transaction
	err = client.Call("DummyRunner.RunAfter", args, &reply)

	if err != nil {
		t.Errorf("rpc client failed to connect to server: %s", err.Error())
	}

	// Verify that method was invoked
	if !run.RunAfterCalled {
		t.Errorf("RunAfter was never invoked")
	}
}

func TestServerRunAfterEachRPC(t *testing.T) {
	client, err := rpc.DialHTTPPath("tcp", addr, "/")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := trans.Transaction{}
	var reply trans.Transaction
	err = client.Call("DummyRunner.RunAfterEach", args, &reply)

	if err != nil {
		t.Errorf("rpc client failed to connect to server: %s", err.Error())
	}

	// Verify that method was invoked
	if !run.RunAfterEachCalled {
		t.Errorf("RunAfterEach was never invoked")
	}
}

func TestServerRunAfterAllRPC(t *testing.T) {
	// client, err := rpc.DialHTTPPath("tcp", addr, "/")
	// if err != nil {
	// 	log.Fatal("dialing:", err)
	// }
	// args := trans.Transaction{}
	// var reply trans.Transaction
	// err = client.Call("DummyRunner.RunAfterAll", args, &reply)

	// if err != nil {
	// 	t.Errorf("rpc client failed to connect to server: %s", err.Error())
	// }

	// // Verify that method was invoked
	// if !run.RunAfterAllCalled {
	// 	t.Errorf("RunAfterAll was never invoked")
	// }
}
