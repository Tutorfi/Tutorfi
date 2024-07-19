package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "time"
)

func fooHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}

func main() {
	// create base storage folder within container
	err := os.Mkdir("storage", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	add_group("group1")
	add_file("group1", "example.txt")
	fmt.Println(get("group1", "example.txt"))

	http.HandleFunc("/foo", fooHandler)

	s := &http.Server{
		Addr: ":8080",
		// Handler:        fooHandler,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
