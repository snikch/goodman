package main

// Runner is responsible for storing and running lifecycle callbacks.
type Runner struct {
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
func NewRunner() *Runner {
	return &Runner{
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

type (
	// Callback is a func type that accepts a Transaction pointer.
	Callback func(*Transaction)
	// AllCallback is a func type that accepts a slice of Transaction pointers.
	AllCallback func([]*Transaction)
)

// BeforeAll adds a callback function to be called before the entire test suite.
func (runner *Runner) BeforeAll(fn AllCallback) {
	runner.beforeAll = append(runner.beforeAll, fn)
}

// BeforeEach adds a callback function to be called before each transaction.
func (runner *Runner) BeforeEach(fn Callback) {
	runner.beforeEach = append(runner.beforeEach, fn)
}

// Before adds a callback function to be called before a named transaction.
func (runner *Runner) Before(name string, fn Callback) {
	if _, ok := runner.before[name]; !ok {
		runner.before[name] = []Callback{}
	}
	runner.before[name] = append(runner.before[name], fn)
}

// BeforeEachValidation adds a callback function to be called before each transaction.
func (runner *Runner) BeforeEachValidation(fn Callback) {
	runner.beforeEachValidation = append(runner.beforeEachValidation, fn)
}

// BeforeValidation adds a callback function to be called before a named transaction.
func (runner *Runner) BeforeValidation(name string, fn Callback) {
	if _, ok := runner.beforeValidation[name]; !ok {
		runner.beforeValidation[name] = []Callback{}
	}
	runner.beforeValidation[name] = append(runner.beforeValidation[name], fn)
}

// After adds a callback function to be called before a named transaction.
func (runner *Runner) After(name string, fn Callback) {
	if _, ok := runner.after[name]; !ok {
		runner.after[name] = []Callback{}
	}
	runner.after[name] = append(runner.after[name], fn)
}

// AfterEach adds a callback function to be called before each transaction.
func (runner *Runner) AfterEach(fn Callback) {
	runner.afterEach = append(runner.afterEach, fn)
}

// AfterAll adds a callback function to be called before the entire test suite.
func (runner *Runner) AfterAll(fn AllCallback) {
	runner.afterAll = append(runner.afterAll, fn)
}

// RunBeforeAll runs all beforeAll callbacks.
func (runner *Runner) RunBeforeAll(transaction []*Transaction) {
	for _, fn := range runner.beforeAll {
		fn(transaction)
	}
}

// RunBeforeEach runs all beforeEach callbacks.
func (runner *Runner) RunBeforeEach(transaction *Transaction) {
	for _, fn := range runner.beforeEach {
		fn(transaction)
	}
}

// RunBefore runs matching before callbacks.
func (runner *Runner) RunBefore(transaction *Transaction) {
	for _, fn := range runner.before[transaction.Name] {
		fn(transaction)
	}
}

// RunBeforeEachValidation runs all beforeEachValidation callbacks.
func (runner *Runner) RunBeforeEachValidation(transaction *Transaction) {
	for _, fn := range runner.beforeEachValidation {
		fn(transaction)
	}
}

// RunBeforeValidation runs matching beforeValidation callbacks.
func (runner *Runner) RunBeforeValidation(transaction *Transaction) {
	for _, fn := range runner.beforeValidation[transaction.Name] {
		fn(transaction)
	}
}

// RunAfter runs matching after callbacks.
func (runner *Runner) RunAfter(transaction *Transaction) {
	for _, fn := range runner.after[transaction.Name] {
		fn(transaction)
	}
}

// RunAfterEach runs all afterEach callbacks.
func (runner *Runner) RunAfterEach(transaction *Transaction) {
	for _, fn := range runner.afterEach {
		fn(transaction)
	}
}

// RunAfterAll runs all afterAll callbacks.
func (runner *Runner) RunAfterAll(transaction []*Transaction) {
	for _, fn := range runner.afterAll {
		fn(transaction)
	}
}
