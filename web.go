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
)

func JoruriServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, HTML)
}

func main() {
	fmt.Println("okaq.jouri starting...")
    http.HandleFunc("/", JoruriServer)
    err := http.ListenAndServe(":8008", nil)
    if err != nil {
        log.Fatal("ListenAndServe: " , err)
    }
}
