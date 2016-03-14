package main

import (
	"log"

	"github.com/apiaryio/goodman"
)

func main() {
	server := goodman.NewServer(NewTestRunner())
	log.Fatal(server.Run())
}
