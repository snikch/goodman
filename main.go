package main

import (
	"log"

	"github.com/snikch/goodman"
)

func main() {
	server := goodman.NewServer(NewTestRunner())
	log.Fatal(server.Run())
}
