// okaq.joruri web server
// json bitmap draw tool
// execute: go run web.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	HTML = "joruri.html"
	PORT = ":8008"
)

func JoruriServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	http.ServeFile(w, req, HTML)
}

func OpenServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    // time simulate network lag
    hi := []byte("hi there from AServer")
    w.Write(hi)
}

func main() {
	fmt.Printf("okaq.jouri starting on port %s\n", PORT)
	http.HandleFunc("/", JoruriServer)
    http.HandleFunc("/aa", OpenServer)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
