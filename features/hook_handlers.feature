Feature: Hook handlers

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

  @debug
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
          ch := make(chan bool)
          h := hooks.NewHooks()
          hooks.NewServer(h, 61322)
          h.BeforeAll(func(t []*trans.Transaction) {
            fmt.Println("before all hook handled")
          })
          h.BeforeEach(func(t *trans.Transaction) {
            fmt.Println("before each hook handled")
          })
          h.Before("/message > GET", func(t *trans.Transaction) {
            fmt.Println("before hook handled")
          })
          h.BeforeEachValidation(func(t *trans.Transaction) {
            fmt.Println("before each validation hook handled")
          })
          h.BeforeValidation("/message > GET", func(t *trans.Transaction) {
            fmt.Println("before validation hook handled")
          })
          h.After("/message > GET", func(t *trans.Transaction) {
            fmt.Println("after hook handled")
          })
          h.AfterEach(func(t *trans.Transaction) {
            fmt.Println("after each hook handled")
          })
          h.AfterAll(func(t []*trans.Transaction) {
            fmt.Println("after all hook handled")
          })
          <-ch
      }


      """
      When I compile to "hookfile"

      # When I run `dredd ./apiary.apib http://localhost:4567 --server "ruby server.rb" --language dredd-hooks-go --hookfiles ./hookfile.go`
    When I run `../../node_modules/.bin/dredd ./apiary.apib http://localhost:4567 --server "ruby server.rb" --language ../../bin/dredd-hooks-go --hookfiles ./aruba --hooks-worker-connect-timeout 3000`
    # Then the exit status should be 0
    Then the output should contain:
      """
      before hook handled
      """
    # And the output should contain:
    #   """
    #   before validation hook handled
    #   """
    And the output should contain:
      """
      after hook handled
      """
    And the output should contain:
      """
      before each hook handled
      """
    # And the output should contain:
    #   """
    #   before each validation hook handled
    #   """
    And the output should contain:
      """
      after each hook handled
      """
    And the output should contain:
      """
      before all hook handled
      """
    And the output should contain:
      """
      after all hook handled
      """
