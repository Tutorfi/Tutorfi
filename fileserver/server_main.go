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

func main() {
	// create base storage folder within container
	err := os.Mkdir("storage", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	syscmd.Add_Group("group1")
	syscmd.Add_File("group1", "example.txt")
	fmt.Println(syscmd.Get("group1", "example.txt"))
	syscmd.Delete_File("group1", "example.txt")

	syscmd.Add_Group("group2")
	syscmd.Add_File("group2", "example.txt")
	fmt.Println(syscmd.Get("group2", "example.txt"))
	syscmd.Delete_Group("group2")

	syscmd.Add_Group("group3")
	syscmd.Add_Dir("group3", "folder")
	syscmd.Add_File("group3", "folder/example.txt")
	fmt.Println(syscmd.Get("group3", "folder/example.txt"))

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
