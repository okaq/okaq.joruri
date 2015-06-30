// okaq.joruri web server
package main

import (
    "fmt"
    "net/http"
)

const (
    DONI = "doni.html"
    PORT = ":8008"
)

func DoniHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, DONI)
}

func main() {
    http.HandleFunc("/doni", DoniHandler)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
