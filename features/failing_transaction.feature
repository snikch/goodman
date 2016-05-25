Feature: Failing a transaction

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
          h.Before("/message > GET", func(t *trans.Transaction) {
              t.Fail = "Yay! Failed!"
              fmt.Println("Yay! Failed!")
          })
          <-ch
        }
      """
      Given I compile to "hookfile"
      # When I run `dredd ./apiary.apib http://localhost:4567 --server "ruby server.rb" --language "dredd-hooks-go" --hookfiles ./hookfile.go`
    When I run `../../node_modules/.bin/dredd ./apiary.apib http://localhost:4567 --server "ruby server.rb" --language ../../bin/dredd-hooks-go --hookfiles ./aruba`
    Then the exit status should be 1
    And the output should contain:
      """
      Yay! Failed!
      """
