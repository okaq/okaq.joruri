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
    FONT = "/home/ahmad/Documents/gira/okaq.joruri/fonts/Kalocsai_Flowers-webfont.woff"
)

func WakaServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, HTML)
}

func FontServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    w.Header().Set("Content-Type", "application/font-woff")
    // w.Header().Set("Content-Type", "text/css")
    http.ServeFile(w, req, FONT)
}

func main() {
    fmt.Printf("okaq.joruri waka starting on port%s\n", PORT)
    http.HandleFunc("/waka", WakaServer)
    http.HandleFunc("/fonts", FontServer)
    // http.Handle("/fonts", http.FileServer(http.Dir("/home/ahmad/Documents/gira/okaq.joruri/fonts")))
    // http.Handle("/", http.FileServer(http.Dir("/home/ahmad/Documents/gira/okaq.joruri")))
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
