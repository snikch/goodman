package hooks

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/snikch/goodman/transaction"
)

type Hooks struct {
	Client *rpc.Client
}

type (
	// Callback is a func type that accepts a Transaction pointer.
	Callback func(*transaction.Transaction)
)

type HookCallback struct {
	Name string
	Fn   Callback
}

func (h *Hooks) Before(name string, fn Callback) {
	args := HookCallback{
		Name: name,
		Fn:   fn,
	}
	h.sendRpcMethod("Before", args)
}

func (h *Hooks) After(name string, fn Callback) {
	args := HookCallback{
		Name: name,
		Fn:   fn,
	}
	h.sendRpcMethod("After", args)
}

func (h *Hooks) BeforeAll(fn Callback) {
	args := HookCallback{
		Fn: fn,
	}
	h.sendRpcMethod("BeforeAll", args)
}

func (h *Hooks) AfterAll(fn Callback) {
	args := HookCallback{
		Fn: fn,
	}
	h.sendRpcMethod("AfterAll", args)
}

func (h *Hooks) BeforeEach(fn Callback) {
	args := HookCallback{
		Fn: fn,
	}
	h.sendRpcMethod("BeforeEach", args)
}

func (h *Hooks) AfterEach(fn Callback) {
	args := HookCallback{
		Fn: fn,
	}
	h.sendRpcMethod("AfterEach", args)
}

func (h *Hooks) BeforeEachValidation(fn Callback) {
	args := HookCallback{
		Fn: fn,
	}
	h.sendRpcMethod("BeforeEachValidation", args)
}

func (h *Hooks) BeforeValidation(name string, fn Callback) {
	args := HookCallback{
		Name: name,
		Fn:   fn,
	}
	h.sendRpcMethod("BeforeValidation", args)
}

func (h *Hooks) sendRpcMethod(method string, fn HookCallback) error {
	var reply bool
	err := h.Client.Call("RpcServer."+method, &fn, &reply)

	if err != nil {
		// TODO: Remove this, RPC method will not work since it is unable to
		// serialize functions

		// panic(err.Error())
	}
	return nil
}

func (h *Hooks) connect(serverAddress string, port int) error {
	addr := fmt.Sprintf("%s:%d", serverAddress, port)
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("dialing: ", err)
		return err
	}
	h.Client = client
	return nil
}

func Default() *Hooks {
	h := &Hooks{}
	h.connect("127.0.0.1", 61322)
	return h
}

type HooksServer interface {
	BeforeAll(fn HookCallback, reply *bool) error
	BeforeEach(fn HookCallback, reply *bool) error
	Before(fn HookCallback, reply *bool) error
	BeforeEachValidation(fn HookCallback, reply *bool) error
	BeforeValidation(fn HookCallback, reply *bool) error
	After(fn HookCallback, reply *bool) error
	AfterEach(fn HookCallback, reply *bool) error
	AfterAll(fn HookCallback, reply *bool) error
}
