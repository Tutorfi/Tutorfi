package make

import (
	"log"
	"net/http"
	"strings"
)

func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", noDirListing(http.FileServer(http.Dir("./static/"))))
	log.Println(http.ListenAndServe(":8080", nil))
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"path/filepath"
// )

// func main() {
//     mux := http.NewServeMux()

//     fileServer := http.FileServer(neuteredFileSystem{http.Dir("../../data")})
//     mux.Handle("/data/", http.NotFoundHandler())
//     mux.Handle("/data/", http.StripPrefix("/data", fileServer))

//     err := http.ListenAndServe(":4000", mux)
//     log.Fatal(err)
// }

// type neuteredFileSystem struct {
//     fs http.FileSystem
// }

// func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
//     f, err := nfs.fs.Open(path)
//     if err != nil {
//         return nil, err
//     }

//     s, err := f.Stat()
//     if s.IsDir() {
//         index := filepath.Join(path, "index.html")
//         if _, err := nfs.fs.Open(index); err != nil {
//             closeErr := f.Close()
//             if closeErr != nil {
//                 return nil, closeErr
//             }

//             return nil, err
//         }
//     }

//     return f, nil
// }    
// package main

// import (
// 	// "fmt"
// 	"net/http"
// 	// "os"
// 	"log"
// 	// "github.com/gorilla/mux"
// 	// "path/filepath"
// )

// // const serverPort = 3333

// func main() {

// 	mux := http.NewServeMux()
// 	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	fmt.Printf("server: %s /\n", r.Method)
// 	// })

// 	// server := http.Server{
// 	// 	Addr:    fmt.Sprintf(":%d", serverPort),
// 	// 	Handler: mux,
// 	// }

// 	// if err := server.ListenAndServe(); err != nil {
// 	// 	if !errors.Is(err, http.ErrServerClosed) {
// 	// 		fmt.Printf("error running http server: %s\n", err)
// 	// 	}
// 	// }

// 	// r := mux.NewRouter()
//     // r.HandleFunc("/{fileName}", sendFile).Methods("GET")
// 	// http.ListenAndServe(":8080", r)

// 	// http.Handle("/", fs)
// 	fs := http.FileServer(http.Dir("../../data"))
// 	mux.Handle("../../data", http.NotFoundHandler())
// 	mux.Handle("../../data", http.StripPrefix("../../data/", fs))

// 	err := http.ListenAndServe(":4000", mux)
//     log.Fatal(err)

// 	// http.Handle("/", fs)

// 	// log.Print("Listening on :3000...")
// 	// err := http.ListenAndServe(":3000", nil)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// }

// // func sendFile(responseWriter http.ResponseWriter, request *http.Request){
// // 	vars := mux.Vars(request)
// //     fileName, ok := vars["fileName"]
// //     if !ok {
// //         fmt.Println("fileName is missing in parameters")
// //     }
// // 	file, err := os.ReadFile("../../data/"+fileName)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		return
// // 	 }
// // 	 os.Stdout.Write(file)

// // 	//  fmt.Println((file));
// //     fmt.Println(`fileName := `, fileName)
// // 	fs := http.FileServer((http.Dir("../../data/"+fileName)))
// // 	http.Handle("/"+fileName, fs)
// // }