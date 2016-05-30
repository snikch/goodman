package hooks

import trans "github.com/snikch/goodman/transaction"

type (
	// Callback is a func type that accepts a Transaction pointer.
	Callback func(*trans.Transaction)
	// AllCallback is a func type that accepts a slice of Transaction pointers.
	AllCallback func([]*trans.Transaction)
)

// Runner is responsible for storing and running lifecycle callbacks.
type Hooks struct {
	beforeAll            []AllCallback
	beforeEach           []Callback
	before               map[string][]Callback
	beforeEachValidation []Callback
	beforeValidation     map[string][]Callback
	after                map[string][]Callback
	afterEach            []Callback
	afterAll             []AllCallback
}

// NewRunner returns a new Runner instance will all callback fields initialized.
func NewHooks() *Hooks {
	return &Hooks{
		beforeAll:            []AllCallback{},
		beforeEach:           []Callback{},
		before:               map[string][]Callback{},
		beforeEachValidation: []Callback{},
		beforeValidation:     map[string][]Callback{},
		after:                map[string][]Callback{},
		afterEach:            []Callback{},
		afterAll:             []AllCallback{},
	}
}

func (h *Hooks) RunBeforeAll(args []*trans.Transaction, reply *[]*trans.Transaction) error {
	reply = &args
	for _, cb := range h.beforeAll {
		cb(args)
	}
	return nil
}

func (h *Hooks) RunBeforeEach(args trans.Transaction, reply *trans.Transaction) error {
	reply = &args
	for _, cb := range h.beforeEach {
		cb(reply)
	}
	return nil
}
func (h *Hooks) RunBefore(args trans.Transaction, reply *trans.Transaction) error {
	name := args.Name
	reply = &args
	for _, cb := range h.before[name] {
		cb(reply)
	}
	return nil
}

func (h *Hooks) RunBeforeEachValidation(args trans.Transaction, reply *trans.Transaction) error {
	reply = &args
	for _, cb := range h.beforeEachValidation {
		cb(reply)
	}
	return nil
}
func (h *Hooks) RunBeforeValidation(args trans.Transaction, reply *trans.Transaction) error {
	name := args.Name
	reply = &args
	for _, cb := range h.beforeValidation[name] {
		cb(reply)
	}
	return nil
}

func (h *Hooks) RunAfter(args trans.Transaction, reply *trans.Transaction) error {
	name := args.Name
	reply = &args
	for _, cb := range h.after[name] {
		cb(reply)
	}
	return nil
}

func (h *Hooks) RunAfterEach(args trans.Transaction, reply *trans.Transaction) error {
	reply = &args
	for _, cb := range h.afterEach {
		cb(reply)
	}
	return nil
}

func (h *Hooks) RunAfterAll(args []*trans.Transaction, reply *[]*trans.Transaction) error {
	reply = &args
	for _, cb := range h.afterAll {
		cb(args)
	}
	return nil
}

// BeforeAll adds a callback function to be called before the entire test suite.
func (h *Hooks) BeforeAll(fn AllCallback) {
	h.beforeAll = append(h.beforeAll, fn)
}

// BeforeEach adds a callback function to be called before each transaction.
func (h *Hooks) BeforeEach(fn Callback) {
	h.beforeEach = append(h.beforeEach, fn)
}

// Before adds a callback function to be called before a named transaction.
func (h *Hooks) Before(name string, fn Callback) {
	if _, ok := h.before[name]; !ok {
		h.before[name] = []Callback{}
	}
	h.before[name] = append(h.before[name], fn)
}

// BeforeEachValidation adds a callback function to be called before each transaction.
func (h *Hooks) BeforeEachValidation(fn Callback) {
	h.beforeEachValidation = append(h.beforeEachValidation, fn)
}

// BeforeValidation adds a callback function to be called before a named transaction.
func (h *Hooks) BeforeValidation(name string, fn Callback) {
	if _, ok := h.beforeValidation[name]; !ok {
		h.beforeValidation[name] = []Callback{}
	}
	h.beforeValidation[name] = append(h.beforeValidation[name], fn)
}

// After adds a callback function to be called before a named transaction.
func (h *Hooks) After(name string, fn Callback) {
	if _, ok := h.after[name]; !ok {
		h.after[name] = []Callback{}
	}
	h.after[name] = append(h.after[name], fn)
}

// AfterEach adds a callback function to be called before each transaction.
func (h *Hooks) AfterEach(fn Callback) {
	h.afterEach = append(h.afterEach, fn)
}

// AfterAll adds a callback function to be called before the entire test suite.
func (h *Hooks) AfterAll(fn AllCallback) {
	h.afterAll = append(h.afterAll, fn)
}

// func (h *Hooks) Before(name string, fn Callback) {
// 	args := HookCallback{
// 		Name: name,
// 		Fn:   fn,
// 	}
// 	h.sendRpcMethod("Before", args)
// }

// func (h *Hooks) After(name string, fn Callback) {
// 	args := HookCallback{
// 		Name: name,
// 		Fn:   fn,
// 	}
// 	h.sendRpcMethod("After", args)
// }

// func (h *Hooks) BeforeAll(fn Callback) {
// 	args := HookCallback{
// 		Fn: fn,
// 	}
// 	h.sendRpcMethod("BeforeAll", args)
// }

// func (h *Hooks) AfterAll(fn Callback) {
// 	args := HookCallback{
// 		Fn: fn,
// 	}
// 	h.sendRpcMethod("AfterAll", args)
// }

// func (h *Hooks) BeforeEach(fn Callback) {
// 	args := HookCallback{
// 		Fn: fn,
// 	}
// 	h.sendRpcMethod("BeforeEach", args)
// }

// func (h *Hooks) AfterEach(fn Callback) {
// 	args := HookCallback{
// 		Fn: fn,
// 	}
// 	h.sendRpcMethod("AfterEach", args)
// }

// func (h *Hooks) BeforeEachValidation(fn Callback) {
// 	args := HookCallback{
// 		Fn: fn,
// 	}
// 	h.sendRpcMethod("BeforeEachValidation", args)
// }

// func (h *Hooks) BeforeValidation(name string, fn Callback) {
// 	args := HookCallback{
// 		Name: name,
// 		Fn:   fn,
// 	}
// 	h.sendRpcMethod("BeforeValidation", args)
// }

// func (h *Hooks) sendRpcMethod(method string, fn HookCallback) error {
// 	var reply bool
// 	err := h.Client.Call("RpcServer."+method, &fn, &reply)

// 	if err != nil {
// 		// TODO: Remove this, RPC method will not work since it is unable to
// 		// serialize functions

// 		// panic(err.Error())
// 	}
// 	return nil
// }

// func (h *Hooks) connect(serverAddress string, port int) error {
// 	addr := fmt.Sprintf("%s:%d", serverAddress, port)
// 	client, err := rpc.DialHTTP("tcp", addr)
// 	if err != nil {
// 		log.Fatal("dialing: ", err)
// 		return err
// 	}
// 	h.Client = client
// 	return nil
// }

// func Default() *Hooks {
// 	h := &Hooks{}
// 	h.connect("127.0.0.1", 61322)
// 	return h
// }
