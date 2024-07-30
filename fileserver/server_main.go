package main

import (
	syscmd "fileserver/syscmd"
	"fmt"
	"log"
	"net/http"
	"os"
)

func fooHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}

func add_group_handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("add_group_handler")

	syscmd.Add_Group("group")
}

func main() {
	// create base storage folder within container
	err := os.Mkdir("storage", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	http.HandleFunc("/foo", fooHandler)

	http.HandleFunc("/add_group", add_group_handler)

	s := &http.Server{
		Addr: ":8080",
		// Handler:        fooHandler,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
