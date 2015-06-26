// okaq.joruri web server
package main

import (
    "fmt"
    "net/http"
)

const (
    BONI = "boni.html"
    PORT = ":8008"
)

func BoniHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, BONI)
}

func main() {
    fmt.Printf("okaq.joruri shima aoni web app started on port %s\n", PORT)
    http.HandleFunc("/boni", BoniHandler)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
