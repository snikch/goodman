package goodman

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/snikch/goodman/hooks"
	t "github.com/snikch/goodman/transaction"
)

// Runner is responsible for storing and running lifecycle callbacks.
type Runner struct {
	beforeAll            []hooks.Callback
	beforeEach           []hooks.Callback
	before               map[string][]hooks.Callback
	beforeEachValidation []hooks.Callback
	beforeValidation     map[string][]hooks.Callback
	after                map[string][]hooks.Callback
	afterEach            []hooks.Callback
	afterAll             []hooks.Callback
}

// NewRunner returns a new Runner instance will all callback fields initialized.
func NewRunner() *Runner {
	return &Runner{
		beforeAll:            []hooks.Callback{},
		beforeEach:           []hooks.Callback{},
		before:               map[string][]hooks.Callback{},
		beforeEachValidation: []hooks.Callback{},
		beforeValidation:     map[string][]hooks.Callback{},
		after:                map[string][]hooks.Callback{},
		afterEach:            []hooks.Callback{},
		afterAll:             []hooks.Callback{},
	}
}

func Start(server *Runner) {
	addr := "127.0.0.1:61322"
	if server == nil {
		server = new(Runner)
	}
	rpc.Register(server)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// BeforeAll adds a callback function to be called before the entire test suite.
func (runner *Runner) BeforeAll(fn hooks.HookCallback, reply *bool) error {
	runner.beforeAll = append(runner.beforeAll, fn.Fn)
	return nil
}

// BeforeEach adds a callback function to be called before each transaction.
func (runner *Runner) BeforeEach(fn hooks.HookCallback, reply *bool) error {
	runner.beforeEach = append(runner.beforeEach, fn.Fn)
	return nil
}

// Before adds a callback function to be called before a named transaction.
func (runner *Runner) Before(fn hooks.HookCallback, reply *bool) error {
	name := fn.Name
	if _, ok := runner.before[name]; !ok {
		runner.before[name] = []hooks.Callback{}
	}
	runner.before[name] = append(runner.before[name], fn.Fn)
	return nil
}

// BeforeEachValidation adds a callback function to be called before each transaction.
func (runner *Runner) BeforeEachValidation(fn hooks.HookCallback, reply *bool) error {
	runner.beforeEachValidation = append(runner.beforeEachValidation, fn.Fn)
	return nil
}

// BeforeValidation adds a callback function to be called before a named transaction.
func (runner *Runner) BeforeValidation(fn hooks.HookCallback, reply *bool) error {
	name := fn.Name
	if _, ok := runner.beforeValidation[name]; !ok {
		runner.beforeValidation[name] = []hooks.Callback{}
	}
	runner.beforeValidation[name] = append(runner.beforeValidation[name], fn.Fn)
	return nil
}

// After adds a callback function to be called before a named transaction.
func (runner *Runner) After(fn hooks.HookCallback, reply *bool) error {
	name := fn.Name
	if _, ok := runner.after[name]; !ok {
		runner.after[name] = []hooks.Callback{}
	}
	runner.after[name] = append(runner.after[name], fn.Fn)
	return nil
}

// AfterEach adds a callback function to be called before each transaction.
func (runner *Runner) AfterEach(fn hooks.HookCallback, reply *bool) error {
	runner.afterEach = append(runner.afterEach, fn.Fn)
	return nil
}

// AfterAll adds a callback function to be called before the entire test suite.
func (runner *Runner) AfterAll(fn hooks.HookCallback, reply *bool) error {
	runner.afterAll = append(runner.afterAll, fn.Fn)
	return nil
}

// RunBeforeAll runs all beforeAll callbacks.
func (runner *Runner) RunBeforeAll(transaction *t.Transaction) {
	for _, fn := range runner.beforeAll {
		fn(transaction)
	}
}

// RunBeforeEach runs all beforeEach callbacks.
func (runner *Runner) RunBeforeEach(transaction *t.Transaction) {
	for _, fn := range runner.beforeEach {
		fn(transaction)
	}
}

// RunBefore runs matching before callbacks.
func (runner *Runner) RunBefore(transaction *t.Transaction) {
	for _, fn := range runner.before[transaction.Name] {
		fn(transaction)
	}
}

// RunBeforeEachValidation runs all beforeEachValidation callbacks.
func (runner *Runner) RunBeforeEachValidation(transaction *t.Transaction) {
	for _, fn := range runner.beforeEachValidation {
		fn(transaction)
	}
}

// RunBeforeValidation runs matching beforeValidation callbacks.
func (runner *Runner) RunBeforeValidation(transaction *t.Transaction) {
	for _, fn := range runner.beforeValidation[transaction.Name] {
		fn(transaction)
	}
}

// RunAfter runs matching after callbacks.
func (runner *Runner) RunAfter(transaction *t.Transaction) {
	for _, fn := range runner.after[transaction.Name] {
		fn(transaction)
	}
}

// RunAfterEach runs all afterEach callbacks.
func (runner *Runner) RunAfterEach(transaction *t.Transaction) {
	for _, fn := range runner.afterEach {
		fn(transaction)
	}
}

// RunAfterAll runs all afterAll callbacks.
func (runner *Runner) RunAfterAll(transaction *t.Transaction) {
	for _, fn := range runner.afterAll {
		fn(transaction)
	}
}

type RunnerInterface interface {
	RunBeforeAll(transaction *t.Transaction)
	RunBeforeEach(transaction *t.Transaction)
	RunBefore(transaction *t.Transaction)
	RunBeforeValidation(transaction *t.Transaction)
	RunBeforeEachValidation(transaction *t.Transaction)
	RunAfterAll(transaction *t.Transaction)
	RunAfterEach(transaction *t.Transaction)
	RunAfter(transaction *t.Transaction)
}
