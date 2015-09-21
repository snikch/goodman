# Goodman

[![Godoc Reference](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/snikch/goodman)

Goodman is a [Dredd](https://github.com/apiaryio/dredd) hook handler implementation in Go.

This is beta software.

## Usage

Write your own goodman server, and run it as an argument to Dredd. Note, you'll need to supply an arbitrary `--hookfiles` param, this is required, but the binary is the main one.

```
 dredd ./blueprint.apib http://localhost:4567 --server "./my-server" --language ./go-hook-server --hookfiles *.rb
 ```

Here is an example usage from the test `cucumber/execution_order.feature`, which consumes all the available callbacks. For more information on the `Transaction` object, see [`transaction.go`](https://github.com/snikch/goodman/blob/master/transaction.go).

```go
package main

import (
	"fmt"
	"log"

	"github.com/snikch/goodman"
)

func main() {
	fmt.Println("Starting")
	server := goodman.NewServer(NewRunner())
	log.Fatal(server.Run())
}

// NewTestRunner creates a runner
func NewRunner() *goodman.Runner {
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
```
