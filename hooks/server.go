package hooks

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	trans "github.com/snikch/goodman/transaction"
)

type Server struct {
}

func NewServer(run RunnerRPC, port int) *Server {
	serv := rpc.NewServer()
	serv.Register(run)
	serv.HandleHTTP("/", "/debug")
	l, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	return &Server{}
}

type RunnerRPC interface {
	RunBeforeAll(args []*trans.Transaction, reply *[]*trans.Transaction) error
	RunBeforeEach(args trans.Transaction, reply *trans.Transaction) error
	RunBefore(args trans.Transaction, reply *trans.Transaction) error
	RunBeforeEachValidation(args trans.Transaction, reply *trans.Transaction) error
	RunBeforeValidation(args trans.Transaction, reply *trans.Transaction) error
	RunAfter(args trans.Transaction, reply *trans.Transaction) error
	RunAfterEach(args trans.Transaction, reply *trans.Transaction) error
	RunAfterAll(args []*trans.Transaction, reply *[]*trans.Transaction) error
}
