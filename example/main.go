package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id    int
	Name  string
	Email string
}

func main() {
	http.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("sqlite3", "./foo.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		rows, err := db.Query(`select id, name, email from users where id = ?;`, 1)

		if err != nil {
			panic(err)
		}

		var id int
		var name string
		var email string
		for rows.Next() {
			err = rows.Scan(&id, &name, &email)

			if err != nil {
				log.Fatal(err)
			}

		}
		user := &User{
			Id:    id,
			Name:  name,
			Email: email,
		}
		data, err := json.Marshal(user)
		if err != nil {
			panic("Marshaling failed with message " + err.Error())
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})

	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("Hello World!\n"))
	})

	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
