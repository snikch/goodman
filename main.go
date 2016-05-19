package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("%#v", os.Args)
	server := NewServer(NewRunner())
	log.Fatal(server.Run())
}
