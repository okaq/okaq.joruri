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

func main() {
	fmt.Printf("okaq.jouri starting on port %s\n", PORT)
	http.HandleFunc("/", JoruriServer)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
