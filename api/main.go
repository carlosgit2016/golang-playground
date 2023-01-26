package main

import (
	"api/entities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        mux,
	}
	// https://pkg.go.dev/net/http
	// GET
	//resp, err := http.Get("http://example.com/")
	// POST /create (Body)
	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	// PUT /update (Body)
	// parameters /create/person
	// queries /get?name=carlos&age=25
	// resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})

	mux.HandleFunc("/foo", FooHandler)
	mux.HandleFunc("/home", HomeHandler)
	mux.HandleFunc("/", NotFound)

	log.Fatal(s.ListenAndServe())
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, "/home", http.StatusFound)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
	person := entities.NewPerson()
	error := json.NewEncoder(w).Encode(person)
	if error != nil {
		panic(error)
	}
}

func FooHandler(w http.ResponseWriter, r *http.Request) {

	type Example struct {
		Foo string "Bar"
	}

	// json.NewEncoder(w).Encode(Example{Foo: "Bar"})
	w.Write([]byte("Hello World"))
}
