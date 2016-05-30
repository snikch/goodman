package goodman

// var run = rpc.DummyRunner{}
// var port = 61323
// var addr = fmt.Sprintf(":%d", port)
// var serv = hooks.NewServer(&run, port)

// func Test

// func TestRunBeforeAll(t *testing.T) {
// 	runner := NewRunner("DummyRunner", port)
// 	ts := &trans.Transaction{}
// 	runner.RunBeforeAll([]*trans.Transaction{ts})

// 	if !run.RunBeforeAllCalled {
// 		t.Errorf("Runner failed to execute RPC call")
// 	}
// }

// func TestRunBeforeEach(t *testing.T) {
// 	runner := NewRunner("DummyRunner", port)
// 	ts := &trans.Transaction{}
// 	runner.RunBeforeEach(ts)

// 	if !run.RunBeforeEachCalled {
// 		t.Errorf("Runner failed to execute RPC call")
// 	}
// }

// func TestRunBefore(t *testing.T) {
// 	runner := NewRunner("DummyRunner", port)
// 	ts := &trans.Transaction{}
// 	runner.RunBefore(ts)

// 	if !run.RunBeforeCalled {
// 		t.Errorf("Runner failed to execute RPC call")
// 	}
// }

// func TestRunBeforeEachValidation(t *testing.T) {
// 	runner := NewRunner("DummyRunner", port)
// 	ts := &trans.Transaction{}
// 	runner.RunBeforeEachValidation(ts)

// 	if !run.RunBeforeEachValidationCalled {
// 		t.Errorf("Runner failed to execute RPC call")
// 	}
// }

// func TestRunBeforeValidation(t *testing.T) {
// 	runner := NewRunner("DummyRunner", port)
// 	ts := &trans.Transaction{}
// 	runner.RunBeforeValidation(ts)

// 	if !run.RunBeforeValidationCalled {
// 		t.Errorf("Runner failed to execute RPC call")
// 	}
// }

// func TestRunAfterAll(t *testing.T) {
// 	runner := NewRunner("DummyRunner", port)
// 	ts := &trans.Transaction{}
// 	runner.RunAfterAll([]*trans.Transaction{ts})

// 	if !run.RunAfterAllCalled {
// 		t.Errorf("Runner failed to execute RPC call")
// 	}
// }

// func TestRunAfterEach(t *testing.T) {
// 	runner := NewRunner("DummyRunner", port)
// 	ts := &trans.Transaction{}
// 	runner.RunAfterEach(ts)

// 	if !run.RunAfterEachCalled {
// 		t.Errorf("Runner failed to execute RPC call")
// 	}
// }

// func TestRunAfter(t *testing.T) {
// 	runner := NewRunner("DummyRunner", port)
// 	ts := &trans.Transaction{}
// 	runner.RunAfter(ts)

// 	if !run.RunAfterCalled {
// 		t.Errorf("Runner failed to execute RPC call")
// 	}
// }
