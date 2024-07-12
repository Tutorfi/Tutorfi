package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	// "time"
)

func fooHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}

func main() {
	err := os.Mkdir("storage", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("storage/testfile.txt", []byte("Hello, World!\n"), 0660)
	if err != nil {
		log.Fatal(err)
	}

	// data, err := os.ReadFile("testdir/testfile.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// os.Stdout.Write(data)

	cmd := exec.Command("cat", "storage/testfile.txt")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))

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
