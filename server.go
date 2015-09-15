package goodman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

const (
	defaultPort             = "61321"
	defaultMessageDelimiter = "\n"
)

// Server is responsible for starting a server and running lifecycle callbacks.
type Server struct {
	Runner           *Runner
	Port             string
	MessageDelimeter []byte
	conn             net.Conn
}

// NewServer returns a new server instance with the supplied runner. If no
// runner is supplied, a new one will be created.
func NewServer(runner *Runner) *Server {
	if runner == nil {
		runner = NewRunner()
	}
	return &Server{
		Runner:           runner,
		Port:             defaultPort,
		MessageDelimeter: []byte(defaultMessageDelimiter),
	}
}

// Run starts the server listening for events from dredd.
func (server *Server) Run() error {
	fmt.Println("Starting")
	ln, err := net.Listen("tcp", ":"+server.Port)
	if err != nil {
		return err
	}
	conn, err := ln.Accept()
	if err != nil {
		return err
	}

	defer conn.Close()
	server.conn = conn

	for {
		body, err := bufio.
			NewReader(conn).
			ReadString(byte(server.MessageDelimeter[0]))
		if err != nil {
			return err
		}

		body = body[:len(body)-1]
		m := &message{}
		err = json.Unmarshal([]byte(body), m)
		if err != nil {
			return err
		}
		err = server.ProcessMessage(m)
		if err != nil {
			return err
		}
	}
}

// ProcessMessage handles a single event message.
func (server *Server) ProcessMessage(m *message) error {
	switch m.Event {
	case "beforeAll":
		fallthrough
	case "afterAll":
		m.transactions = []*Transaction{}
		err := json.Unmarshal(m.Data, &m.transactions)
		if err != nil {
			return err
		}
	default:
		m.transaction = &Transaction{}
		err := json.Unmarshal(m.Data, m.transaction)
		if err != nil {
			return err
		}
	}

	switch m.Event {
	case "beforeAll":
		server.Runner.RunBeforeAll(m.transactions)
		break
	case "beforeEach":
		server.Runner.RunBeforeEach(m.transaction)
		break
	case "before":
		server.Runner.RunBefore(m.transaction)
		break
	case "beforeEachValidation":
		server.Runner.RunBeforeEachValidation(m.transaction)
		break
	case "beforeValidation":
		server.Runner.RunBeforeValidation(m.transaction)
		break
	case "after":
		server.Runner.RunAfter(m.transaction)
		break
	case "afterEach":
		server.Runner.RunAfterEach(m.transaction)
		break
	case "afterAll":
		server.Runner.RunAfterAll(m.transactions)
		break
	default:
		return fmt.Errorf("Unknown event '%s'", m.Event)
	}

	switch m.Event {
	case "beforeAll":
		fallthrough
	case "afterAll":
		return server.sendResponse(m, m.transactions)
	default:
		return server.sendResponse(m, m.transaction)
	}
}

// sendResponse submits the transaction(s) back to dredd.
func (server *Server) sendResponse(m *message, dataObj interface{}) error {
	data, err := json.Marshal(dataObj)
	if err != nil {
		return err
	}

	m.Data = json.RawMessage(data)
	response, err := json.Marshal(m)
	if err != nil {
		return err
	}
	server.conn.Write(response)
	server.conn.Write(server.MessageDelimeter)
	return nil
}

// message represents a single event received over the connection.
type message struct {
	UUID  string          `json:"uuid"`
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`

	transaction  *Transaction
	transactions []*Transaction
}
