// unki - patch based texture synth
// AQ <aq@okaq.com>
// 2016-03-26
package main

import (
    "fmt"
    "net/http"
)

const (
    BNKI = "bnki.html"
)

func BnkiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, BNKI)
}

func main() {
    fmt.Println("starting unki web on localhost:8800")
    http.HandleFunc("/", BnkiHandler)
    http.ListenAndServe(":8800", nil)
}
