package rpc

import (
	trans "github.com/snikch/goodman/transaction"
)

// DummyRunner is an implementation of the hooks.Runner interface
// it is strictly used for testing to ensure that the hooks.server
// serves its rpc correctly.
type DummyRunner struct {
	RunBeforeAllCalled            bool
	RunBeforeEachCalled           bool
	RunBeforeCalled               bool
	RunBeforeEachValidationCalled bool
	RunBeforeValidationCalled     bool
	RunAfterCalled                bool
	RunAfterEachCalled            bool
	RunAfterAllCalled             bool
}

func (run *DummyRunner) RunBeforeAll(args []*trans.Transaction, reply *[]*trans.Transaction) error {
	run.RunBeforeAllCalled = true
	return nil
}

func (run *DummyRunner) RunBeforeEach(args trans.Transaction, reply *trans.Transaction) error {
	run.RunBeforeEachCalled = true
	return nil
}

func (run *DummyRunner) RunBefore(args trans.Transaction, reply *trans.Transaction) error {
	run.RunBeforeCalled = true
	return nil
}

func (run *DummyRunner) RunBeforeEachValidation(args trans.Transaction, reply *trans.Transaction) error {
	run.RunBeforeEachValidationCalled = true
	return nil
}

func (run *DummyRunner) RunBeforeValidation(args trans.Transaction, reply *trans.Transaction) error {
	run.RunBeforeValidationCalled = true
	return nil
}
func (run *DummyRunner) RunAfter(args trans.Transaction, reply *trans.Transaction) error {
	run.RunAfterCalled = true
	return nil
}

func (run *DummyRunner) RunAfterEach(args trans.Transaction, reply *trans.Transaction) error {
	run.RunAfterEachCalled = true
	return nil
}

func (run *DummyRunner) RunAfterAll(args []*trans.Transaction, reply *[]*trans.Transaction) error {
	run.RunAfterAllCalled = true
	return nil
}
