Feature: Execution order

  Background:
    Given I have "go" command installed
    And I have "dredd" command installed
    And a file named "server.rb" with:
      """
      require 'sinatra'
      get '/message' do
        "Hello World!\n\n"
      end
      """

    And a file named "apiary.apib" with:
      """
      # My Api
      ## GET /message
      + Response 200 (text/html;charset=utf-8)
          Hello World!
      """

  @announce
  Scenario:
    Given a file named "../../hookfile.go" with:
    """
    package main
    import "github.com/snikch/goodman"
    // NewTestRunner creates a test runner
    func NewTestRunner() *goodman.Runner {
    	runner := goodman.NewRunner()
    	runner.BeforeAll(func(t []*goodman.Transaction) {
    		t[0].AddTestOrderPoint("before all modification")
    	})
    	runner.BeforeEach(func(t *goodman.Transaction) {
    		t.AddTestOrderPoint("before each modification")
    	})
    	runner.Before("/message > GET", func(t *goodman.Transaction) {
    		t.AddTestOrderPoint("before modification")
    	})
    	runner.BeforeEachValidation(func(t *goodman.Transaction) {
    		t.AddTestOrderPoint("before each validation modification")
    	})
    	runner.BeforeValidation("/message > GET", func(t *goodman.Transaction) {
    		t.AddTestOrderPoint("before validation modification")
    	})
    	runner.After("/message > GET", func(t *goodman.Transaction) {
    		t.AddTestOrderPoint("after modification")
    	})
    	runner.AfterEach(func(t *goodman.Transaction) {
    		t.AddTestOrderPoint("after each modification")
    	})
    	runner.AfterAll(func(t []*goodman.Transaction) {
    		t[0].AddTestOrderPoint("after all modification")
    	})
    	return runner
    }
    """
    When I compile to "hookfile-go"
    And I set the environment variables to:
      | variable                       | value      |
      | TEST_DREDD_HOOKS_HANDLER_ORDER | true       |

    When I run `dredd ./apiary.apib http://localhost:4567 --server "ruby server.rb" --language ../../hookfile-go --hookfiles *.rb`
    Then the exit status should be 0
    Then the output should contain:
      """
      0 before all modification
      1 before each modification
      2 before modification
      3 before each validation modification
      4 before validation modification
      5 after modification
      6 after each modification
      7 after all modification
      """
