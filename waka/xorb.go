// okaq.joruri web server
// waka card assembler tool
// AQ <aq@okaq.com>
// 2014-12-24
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    HTML = "duka.html"
    PORT = ":8080"
)

var (
    Now time.Time
)

func WakaServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, HTML)
}

func main() {
    Now = time.Now()
    fmt.Printf("okaq.joruri waka starting on port%s\n", PORT)
    fmt.Printf("Started at: %s.\n", Now.Format(time.RFC1123Z))
    http.HandleFunc("/waka", WakaServer)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
