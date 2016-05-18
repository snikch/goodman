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
      ## Implement following in your language utilizing each hook declaring function
      ## from API in your language:
      ## - write to standard output name of hook + "hook handled" e.g: "after hook handled"
      ##
      ## So, replace following pseudo code with yours:
      #
      #require 'mylanguagehooks'
      #
      #before("/message > GET") { |transaction|
      #  echo "before hook handled"
      #}
      #
      #after("/message > GET") { |transaction|
      #  echo "after hook handled"
      #}
      #
      #before_validation("/message > GET") { |transaction|
      #  echo "before validation hook handled"
      #}
      #
      #before_all { |transaction|
      #  echo "before all hook handled"
      #}
      #
      #after_all { |transaction|
      #  echo "after all hook handled"
      #}
      #
      #before_each { |transaction|
      #  echo "before each hook handled"
      #}
      #
      #before_each_validation { |transaction|
      #  echo "before each validation hook handled"
      #}

      #after_each { |transaction|
      #  echo "after each hook handled"
      #}

      """

    When I run `dredd ./apiary.apib http://localhost:4567 --server "ruby server.rb" --language dredd-hooks-go --hookfiles ./hookfile.go`
    Then the exit status should be 0
    Then the output should contain:
      """
      before hook handled
      """
    And the output should contain:
      """
      before validation hook handled
      """
    And the output should contain:
      """
      after hook handled
      """
    And the output should contain:
      """
      before each hook handled
      """
    And the output should contain:
      """
      before each validation hook handled
      """
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
