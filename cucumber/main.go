package main

import (
	"fmt"
	"log"

	"github.com/snikch/goodman"
)

func main() {
	fmt.Println("Starting")
	server := goodman.NewServer(NewTestRunner())
	log.Fatal(server.Run())
}
