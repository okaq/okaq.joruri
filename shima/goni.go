// okaq.joruri shima web server
// AQ <aq@okaq.com>
// 2015-07-26
package main

import (
    "fmt"
    "net/http"
)

const (
    HONI = "honi.html"
    PORT = ":8008"
)

func HoniHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, HONI)
}

func main() {
    fmt.Println("okaq joruri shima goni web server is live...")
    http.HandleFunc("/", HoniHandler)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
