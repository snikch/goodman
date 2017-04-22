package goodman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	t "github.com/snikch/goodman/transaction"
)

const (
	defaultMessageDelimiter = "\n"
)

// Server is responsible for starting a server and running lifecycle callbacks.
type Server struct {
	Runner           []Runner
	MessageDelimeter []byte
	listener         net.Listener
	conn             net.Conn
}

// NewServer returns a new server instance with the supplied runner. If no
// runner is supplied, a new one will be created.
func NewServer(runners []Runner, port int) (*Server, error) {
	log.Printf("trying to listen on %d", port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return nil, fmt.Errorf("listening on tcp: %s", err.Error())
	}

	server := Server{
		Runner:           runners,
		MessageDelimeter: []byte(defaultMessageDelimiter),
		listener:         ln,
	}
	return &server, nil
}

// Run starts the server listening for events from dredd.
func (server *Server) Run() error {
	fmt.Println("Starting")
	conn, err := server.listener.Accept()
	if err != nil {
		return fmt.Errorf("accepting on tcp: %s", err.Error())
	}

	server.conn = conn

	for {
		body, err := bufio.
			NewReader(server.conn).
			ReadString('\n')
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		body = body[:len(body)-1]
		m := &message{}
		if err := json.Unmarshal([]byte(body), m); err != nil {
			return fmt.Errorf("unmarshaling body: %s", err.Error())
		}
		if err := server.ProcessMessage(m); err != nil {
			return fmt.Errorf("processing message: %s", err.Error())
		}
	}
}

func (s Server) Close() (err error) {
	if listenerErr := s.listener.Close(); listenerErr != nil {
		err = fmt.Errorf("closing listener: %s", listenerErr.Error())
	}
	if s.conn != nil {
		if connErr := s.conn.Close(); connErr != nil {
			err = fmt.Errorf("closing server tcp connection: %s", connErr.Error())
		}
	}
	return err
}

// ProcessMessage handles a single event message.
func (server *Server) ProcessMessage(m *message) error {
	switch m.Event {
	case "beforeAll", "afterAll":
		m.transactions = []*t.Transaction{}
		err := json.Unmarshal(m.Data, &m.transactions)
		if err != nil {
			return err
		}
	default:
		m.transaction = &t.Transaction{}
		err := json.Unmarshal(m.Data, m.transaction)
		if err != nil {
			return err
		}
	}

	switch m.Event {
	case "beforeAll":
		server.RunBeforeAll(&m.transactions)
		break
	case "beforeEach":
		// before is run after beforeEach, as no separate event is fired.
		server.RunBeforeEach(m.transaction)
		server.RunBefore(m.transaction)
		break
	case "beforeEachValidation":
		// beforeValidation is run after beforeEachValidation, as no separate event
		// is fired.
		server.RunBeforeEachValidation(m.transaction)
		server.RunBeforeValidation(m.transaction)
		break
	case "afterEach":
		// after is run before afterEach as no separate event is fired.
		server.RunAfter(m.transaction)
		server.RunAfterEach(m.transaction)
		break
	case "afterAll":
		server.RunAfterAll(&m.transactions)
		break
	default:
		return fmt.Errorf("Unknown event '%s'", m.Event)
	}

	switch m.Event {
	case "beforeAll", "afterAll":
		return server.sendResponse(m, m.transactions)
	default:
		return server.sendResponse(m, m.transaction)
	}
}

func (server *Server) RunBeforeAll(trans *[]*t.Transaction) {
	for _, runner := range server.Runner {
		runner.RunBeforeAll(trans)
	}
}

func (server *Server) RunBeforeEach(trans *t.Transaction) {
	for _, runner := range server.Runner {
		runner.RunBeforeEach(trans)
	}
}

func (server *Server) RunBefore(trans *t.Transaction) {
	for _, runner := range server.Runner {
		runner.RunBefore(trans)
	}
}

func (server *Server) RunBeforeEachValidation(trans *t.Transaction) {
	for _, runner := range server.Runner {
		runner.RunBeforeEachValidation(trans)
	}
}

func (server *Server) RunBeforeValidation(trans *t.Transaction) {
	for _, runner := range server.Runner {
		runner.RunBeforeValidation(trans)
	}
}

func (server *Server) RunAfterEach(trans *t.Transaction) {
	for _, runner := range server.Runner {
		runner.RunAfterEach(trans)
	}
}

func (server *Server) RunAfter(trans *t.Transaction) {
	for _, runner := range server.Runner {
		runner.RunAfter(trans)
	}
}

func (server *Server) RunAfterAll(trans *[]*t.Transaction) {
	for _, runner := range server.Runner {
		runner.RunAfterAll(trans)
	}
}

// sendResponse submits the transaction(s) back to dredd.
func (server *Server) sendResponse(m *message, dataObj interface{}) error {
	data, err := json.Marshal(dataObj)
	if err != nil {
		return fmt.Errorf("marshaling data json: %s", err)
	}

	m.Data = json.RawMessage(data)
	response, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("marshaling message json: %s", err)
	}
	if _, err := server.conn.Write(response); err != nil {
		return fmt.Errorf("writing error response: %s", err.Error())
	}
	if _, err := server.conn.Write(server.MessageDelimeter); err != nil {
		return fmt.Errorf("writing message delimeter: %s", err.Error())
	}
	return nil
}

// message represents a single event received over the connection.
type message struct {
	UUID  string          `json:"uuid"`
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`

	transaction  *t.Transaction
	transactions []*t.Transaction
}
