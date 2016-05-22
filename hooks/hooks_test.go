package hooks

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"testing"

	trans "github.com/snikch/goodman/transaction"
)

const host = "127.0.0.1"
const port = 1235

// This struct implements the HooksServer interface and is used to test the Hook client's behavior independent of the real HooksServer implementation
type RpcServer struct {
	Test bool
}

var beforeCalled = false
var afterCalled = false
var beforeAllCalled = false
var afterAllCalled = false
var beforeEachCalled = false
var afterEachCalled = false
var beforeEachValidationCalled = false
var beforeValidationCalled = false
var nullCallbackErrorMessage = "Callback passed to HooksServer is nil"

func (server *RpcServer) Before(args HookCallback, reply *bool) error {
	beforeCalled = true
	if args.Fn == nil {
		return errors.New(nullCallbackErrorMessage)
	}
	return nil
}

func (server *RpcServer) After(args HookCallback, reply *bool) error {
	afterCalled = true
	if args.Fn == nil {
		return errors.New(nullCallbackErrorMessage)
	}
	return nil
}

func (server *RpcServer) BeforeAll(args HookCallback, reply *bool) error {
	beforeAllCalled = true
	// out := fmt.Sprintf("HookCallback %#v", args)
	if args.Fn == nil {
		return errors.New(nullCallbackErrorMessage)
	}
	return nil
}

func (server *RpcServer) AfterAll(args HookCallback, reply *bool) error {
	afterAllCalled = true
	if args.Fn == nil {
		return errors.New(nullCallbackErrorMessage)
	}
	return nil
}

func (server *RpcServer) BeforeEach(args HookCallback, reply *bool) error {
	beforeEachCalled = true
	if args.Fn == nil {
		return errors.New(nullCallbackErrorMessage)
	}
	return nil
}

func (server *RpcServer) AfterEach(args HookCallback, reply *bool) error {
	afterEachCalled = true
	if args.Fn == nil {
		return errors.New(nullCallbackErrorMessage)
	}
	return nil
}

func (server *RpcServer) BeforeEachValidation(args HookCallback, reply *bool) error {
	beforeEachValidationCalled = true
	if args.Fn == nil {
		return errors.New(nullCallbackErrorMessage)
	}
	return nil
}

func (server *RpcServer) BeforeValidation(args HookCallback, reply *bool) error {
	beforeValidationCalled = true
	if args.Fn == nil {
		return errors.New(nullCallbackErrorMessage)
	}
	return nil
}

func Serve(serverAddress string, port int) {
	addr := fmt.Sprintf("%s:%d", serverAddress, port)
	server := new(RpcServer)
	rpc.Register(server)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// RpcServer must implement the HooksServer interface
func TestDummyHooksServer(t *testing.T) {
	var _ HooksServer = &RpcServer{}
}

// TODO: Does server need to be closed here?
func TestHooksConnect(t *testing.T) {
	Serve(host, port)
	hooks := new(Hooks)
	err := hooks.connect(host, port)

	if err != nil {
		t.Errorf("Hooks client failed to connect to server")
	}
}

func TestHooksBefore(t *testing.T) {
	hooks := new(Hooks)
	hooks.connect(host, port)
	hooks.Before("name", func(*trans.Transaction) {

	})

	if !beforeCalled {
		t.Errorf("Before callback was not sent to RPC server")
	}
}

func TestHooksAfter(t *testing.T) {
	hooks := new(Hooks)
	hooks.connect(host, port)
	hooks.After("name", func(*trans.Transaction) {

	})
	if !afterCalled {
		t.Errorf("After callback was not sent to RPC server")
	}
}

func TestHooksBeforeAll(t *testing.T) {
	hooks := new(Hooks)
	hooks.connect(host, port)
	hooks.BeforeAll(func(*trans.Transaction) {

	})

	if !beforeAllCalled {
		t.Errorf("BeforeAll callback was not sent to RPC Server")
	}
}

func TestHooksAfterAll(t *testing.T) {
	hooks := new(Hooks)
	hooks.connect(host, port)
	hooks.AfterAll(func(*trans.Transaction) {

	})

	if !afterAllCalled {
		t.Errorf("AfterAll callback was not sent to RPC Server")
	}
}

func TestHooksBeforeEach(t *testing.T) {
	hooks := new(Hooks)
	hooks.connect(host, port)
	hooks.BeforeEach(func(*trans.Transaction) {

	})

	if !beforeEachCalled {
		t.Errorf("BeforeEach callback was not sent to RPC Server")
	}
}

func TestHooksAfterEach(t *testing.T) {
	hooks := new(Hooks)
	hooks.connect(host, port)
	hooks.AfterEach(func(*trans.Transaction) {

	})

	if !afterEachCalled {
		t.Errorf("AfterEach callback was not sent to RPC Server")
	}
}

func TestHooksBeforeEachValidation(t *testing.T) {
	hooks := new(Hooks)
	hooks.connect(host, port)
	hooks.BeforeEachValidation(func(*trans.Transaction) {

	})

	if !beforeEachValidationCalled {
		t.Errorf("BeforeEachValidation callback was not sent to RPC Server")
	}
}

func TestHooksBeforeValidation(t *testing.T) {
	hooks := new(Hooks)
	hooks.connect(host, port)
	hooks.BeforeValidation("name", func(*trans.Transaction) {

	})

	if !beforeValidationCalled {
		t.Errorf("BeforeValidation callback was not sent to RPC Server")
	}
}
