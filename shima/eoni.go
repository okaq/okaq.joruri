// okaq.joruri web server
package main

import (
    "fmt"
    "net/http"
)

const (
    FONI = "foni.html"
    PORT = ":8008"
)

func FoniHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, FONI)
}

func main() {
    fmt.Println("starting okaq joruri shima web server...")
    http.HandleFunc("/", FoniHandler)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
