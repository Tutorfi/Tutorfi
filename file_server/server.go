package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// fs := http.FileServer(http.Dir("./data"))
	// http.Handle("/", fs)

	r := mux.NewRouter()
    r.HandleFunc("/provisions/{id}", Provisions)
	http.ListenAndServe(":8080", r)
	
	// http.HandleFunc("/hello", helloHandler);
	// http.HandleFunc("/provisions/", Provisions)

	// http.HandleFunc("/", indexHandler)

	// http.ListenAndServe(":8080", nil)
	// print("Server opened on port 8080")
}

func Provisions(responseWriter http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
    id, ok := vars["id"]
    if !ok {
        fmt.Println("id is missing in parameters")
    }
	files, err := os.ReadDir("./data/"+id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	 }
	 fmt.Println((files));
	 for _, file := range files {
		fmt.Println(file.Name())
	 }
    fmt.Println(`id := `, id)
	fs := http.FileServer((http.Dir("./data/"+id)))
	http.Handle("/", fs)
	// fmt.Fprintf(responseWriter, id)
}


// func helloHandler(responseWriter http.ResponseWriter, request *http.Request) {
// 	print(request)
// 	http.ServeFile(responseWriter, request, "./data/Crib_sheet5.pdf")
// }

// func indexHandler(responseWriter http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintf(responseWriter, "Hello!")
// }