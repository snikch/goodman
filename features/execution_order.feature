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
    Given a file named "hookfile.go" with:
    """
    package main
    import (
      "fmt"

      "github.com/snikch/goodman/hooks"
      trans "github.com/snikch/goodman/transaction"
    )

    func main() {

        h := hooks.Default()
        h.BeforeAll(func(t *trans.Transaction) {
          fmt.Println("before all modification")
        })
        h.BeforeEach(func(t *trans.Transaction) {
          fmt.Println("before each modification")
        })
        h.Before("/message > GET", func(t *trans.Transaction) {
          fmt.Println("before modification")
        })
        h.BeforeEachValidation(func(t *trans.Transaction) {
          fmt.Println("before each validation modification")
        })
        h.BeforeValidation("/message > GET", func(t *trans.Transaction) {
          fmt.Println("before validation modification")
        })
        h.After("/message > GET", func(t *trans.Transaction) {
          fmt.Println("after modification")
        })
        h.AfterEach(func(t *trans.Transaction) {
          fmt.Println("after each modification")
        })
        h.AfterAll(func(t *trans.Transaction) {
          fmt.Println("after all modification")
        })
    }
    """
    When I compile to "hookfile"
    And I set the environment variables to:
      | variable                       | value      |
      | TEST_DREDD_HOOKS_HANDLER_ORDER | true       |

    When I run `../../node_modules/.bin/dredd ./apiary.apib http://localhost:4567 --server "dredd-hooks-go" --hookfiles ../tmp/aruba/hookfile-go`
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
