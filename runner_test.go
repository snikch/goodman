package goodman

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"testing"

	"github.com/snikch/goodman/hooks"
	"github.com/snikch/goodman/transaction"
)

func TestNewRunner(t *testing.T) {
	runner := NewRunner()

	if len(runner.beforeEach) != 0 {
		t.Errorf("New runner should have empty beforeEach hooks")
	}

	if len(runner.beforeAll) != 0 {
		t.Errorf("New runner should have empty beforeAll hooks")
	}

	if len(runner.beforeEachValidation) != 0 {
		t.Errorf("New runner should have empty beforeEachValidation hooks")
	}

	if len(runner.beforeValidation) != 0 {
		t.Errorf("New runner should have empty beforeValidation hooks")
	}

	if len(runner.before) != 0 {
		t.Errorf("New runner should have empty before hooks")
	}

	if len(runner.afterEach) != 0 {
		t.Errorf("New runner should have empty afterEach hooks")
	}

	if len(runner.afterAll) != 0 {
		t.Errorf("New runner should have empty afterAll hooks")
	}
}

// TODO: Clean up messy server implementation
func TestBeforeAll(t *testing.T) {
	runner := NewRunner()

	addr := "127.0.0.1:61322"
	server := rpc.NewServer()
	server.Register(runner)
	server.HandleHTTP("/1", "/debug1")
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		http.Serve(l, nil)
	}()

	cb := func(trans *transaction.Transaction) {
		fmt.Println("Running code")
	}

	fn := hooks.HookCallback{
		Fn: cb,
	}
	client, err := rpc.DialHTTPPath("tcp", "127.0.0.1:61322", "/1")
	defer client.Close()

	if err != nil {
		t.Errorf("Connection to HooksServer failed %s", err.Error())
	}

	if client == nil {
		fmt.Printf("%#v", client)
	}
	var reply bool
	err = client.Call("Runner.BeforeAll", fn, &reply)

	if err != nil {
		t.Errorf("Another Error %s", err.Error())
	}

	count := len(runner.beforeAll)

	if count != 1 {
		t.Errorf("Runner should have beforeAll callback")
	}
}

func TestBeforeEach(t *testing.T) {
	runner := NewRunner()

	addr := "127.0.0.1:61324"
	server := rpc.NewServer()
	server.Register(runner)
	server.HandleHTTP("/", "/debug")
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		http.Serve(l, nil)
	}()

	cb := func(trans *transaction.Transaction) {
		fmt.Println("Running code")
	}

	fn := hooks.HookCallback{
		Fn: cb,
	}
	client, err := rpc.DialHTTP("tcp", addr)
	defer client.Close()

	if err != nil {
		t.Errorf("Connection to HooksServer failed %s", err.Error())
	}

	if client == nil {
		fmt.Printf("%#v", client)
	}
	var reply bool
	err = client.Call("Runner.BeforeEach", fn, &reply)

	if err != nil {
		t.Errorf("Another Error %s", err.Error())
	}

	count := len(runner.beforeEach)

	if count != 1 {
		t.Errorf("Runner should have beforeEach callback")
	}
}

func TestBefore(t *testing.T) {
	runner := NewRunner()

	addr := "127.0.0.1:61325"
	server := rpc.NewServer()
	server.Register(runner)
	server.HandleHTTP("/2", "/debug2")
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		http.Serve(l, nil)
	}()

	cb := func(trans *transaction.Transaction) {
		fmt.Println("Running code")
	}

	fn := hooks.HookCallback{
		Name: "name",
		Fn:   cb,
	}
	client, err := rpc.DialHTTPPath("tcp", addr, "/2")
	defer client.Close()

	if err != nil {
		t.Errorf("Connection to HooksServer failed %s", err.Error())
	}

	if client == nil {
		fmt.Printf("%#v", client)
	}
	var reply bool
	err = client.Call("Runner.Before", fn, &reply)

	if err != nil {
		t.Errorf("Another Error %s", err.Error())
	}

	count := len(runner.before)

	if count != 1 {
		t.Errorf("Runner should have before callback")
	}

	if runner.before[fn.Name] == nil {
		t.Errorf("before map should have been set with the name key")
	}
}

func TestBeforeEachValidation(t *testing.T) {
	runner := NewRunner()

	addr := "127.0.0.1:61326"
	server := rpc.NewServer()
	server.Register(runner)
	server.HandleHTTP("/4", "/debug5")
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		http.Serve(l, nil)
	}()

	cb := func(trans *transaction.Transaction) {
		fmt.Println("Running code")
	}

	fn := hooks.HookCallback{
		Fn: cb,
	}
	client, err := rpc.DialHTTPPath("tcp", addr, "/4")
	defer client.Close()

	if err != nil {
		t.Errorf("Connection to HooksServer failed %s", err.Error())
	}

	if client == nil {
		fmt.Printf("%#v", client)
	}
	var reply bool
	err = client.Call("Runner.BeforeEachValidation", fn, &reply)

	if err != nil {
		t.Errorf("Another Error %s", err.Error())
	}

	count := len(runner.beforeEachValidation)

	if count != 1 {
		t.Errorf("Runner should have beforeEachValidation callback")
	}
}

func TestBeforeValidation(t *testing.T) {
	runner := NewRunner()

	addr := "127.0.0.1:61327"
	server := rpc.NewServer()
	server.Register(runner)
	server.HandleHTTP("/5", "/debug6")
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		http.Serve(l, nil)
	}()

	cb := func(trans *transaction.Transaction) {
		fmt.Println("Running code")
	}

	fn := hooks.HookCallback{
		Fn: cb,
	}
	client, err := rpc.DialHTTPPath("tcp", addr, "/5")
	defer client.Close()

	if err != nil {
		t.Errorf("Connection to HooksServer failed %s", err.Error())
	}

	if client == nil {
		fmt.Printf("%#v", client)
	}
	var reply bool
	err = client.Call("Runner.BeforeValidation", fn, &reply)

	if err != nil {
		t.Errorf("Another Error %s", err.Error())
	}

	count := len(runner.beforeValidation)

	if count != 1 {
		t.Errorf("Runner should have beforeValidation callback")
	}
}

func TestAfter(t *testing.T) {
	runner := NewRunner()

	addr := "127.0.0.1:61328"
	server := rpc.NewServer()
	server.Register(runner)
	server.HandleHTTP("/6", "/debug7")
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		http.Serve(l, nil)
	}()

	cb := func(trans *transaction.Transaction) {
		fmt.Println("Running code")
	}

	fn := hooks.HookCallback{
		Name: "After",
		Fn:   cb,
	}
	client, err := rpc.DialHTTPPath("tcp", addr, "/6")
	defer client.Close()

	if err != nil {
		t.Errorf("Connection to HooksServer failed %s", err.Error())
	}

	if client == nil {
		fmt.Printf("%#v", client)
	}
	var reply bool
	err = client.Call("Runner.After", fn, &reply)

	if err != nil {
		t.Errorf("Another Error %s", err.Error())
	}

	count := len(runner.after)

	if count != 1 || runner.after[fn.Name] == nil {
		t.Errorf("Runner should have after callback")
	}
}

func TestAfterEach(t *testing.T) {
	runner := NewRunner()

	addr := "127.0.0.1:61329"
	server := rpc.NewServer()
	server.Register(runner)
	server.HandleHTTP("/7", "/debug8")
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		http.Serve(l, nil)
	}()

	cb := func(trans *transaction.Transaction) {
		fmt.Println("Running code")
	}

	fn := hooks.HookCallback{
		Fn: cb,
	}
	client, err := rpc.DialHTTPPath("tcp", addr, "/7")
	defer client.Close()

	if err != nil {
		t.Errorf("Connection to HooksServer failed %s", err.Error())
	}

	if client == nil {
		fmt.Printf("%#v", client)
	}
	var reply bool
	err = client.Call("Runner.AfterEach", fn, &reply)

	if err != nil {
		t.Errorf("Another Error %s", err.Error())
	}

	count := len(runner.afterEach)

	if count != 1 {
		t.Errorf("Runner should have after callback")
	}
}

func TestAfterAll(t *testing.T) {
	runner := NewRunner()

	addr := "127.0.0.1:61330"
	server := rpc.NewServer()
	server.Register(runner)
	server.HandleHTTP("/8", "/debug9")
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		http.Serve(l, nil)
	}()

	cb := func(trans *transaction.Transaction) {
		fmt.Println("Running code")
	}

	fn := hooks.HookCallback{
		Fn: cb,
	}
	client, err := rpc.DialHTTPPath("tcp", addr, "/8")
	defer client.Close()

	if err != nil {
		t.Errorf("Connection to HooksServer failed %s", err.Error())
	}

	if client == nil {
		fmt.Printf("%#v", client)
	}
	var reply bool
	err = client.Call("Runner.AfterAll", fn, &reply)

	if err != nil {
		t.Errorf("Another Error %s", err.Error())
	}

	count := len(runner.afterAll)

	if count != 1 {
		t.Errorf("Runner should have afterAll callback")
	}
}
