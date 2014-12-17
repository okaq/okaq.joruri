// okaq.joruri web server
// waka image render and sampling tool
// AQ <aq@okaq.com>
// 2014-12-16
package main

import (
    "fmt"
    "net/http"
)

const (
    // HTML = "aomi.html"
    HTML = "bome.html"
    PORT = ":8080"
)

func WakaServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, HTML)
}

func main() {
    fmt.Printf("okaq.joruri waka starting on port%s\n", PORT)
    http.HandleFunc("/waka", WakaServer)
    http.Handle("/fonts", http.FileServer(http.Dir("/home/ahmad/Documents/gira/okaq.joruri/fonts")))
    // http.Handle("/", http.FileServer(http.Dir("/home/ahmad/Documents/gira/okaq.joruri")))
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
