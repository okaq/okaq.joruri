// golang web server
// unki - patch based texture synth
// okaq joruri bitmap draw tools
// AQ <aq@okaq.com>
// 2016-03-12
package main

import (
    "fmt"
    "net/http"
)

const (
    ANKI = "anki.html"
)

func AnkiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, ANKI)
}

func OpenHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ok bitmaps!"))
}

func main() {
    fmt.Println("starting unki web on localhost:8800")
    http.HandleFunc("/", AnkiHandler)
    http.HandleFunc("/open", OpenHandler)
    http.ListenAndServe(":8800", nil)
}
