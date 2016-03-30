// unki - patch based texture synth
// AQ <aq@okaq.com>
// 2016-03-26
package main

import (
    "fmt"
    "net/http"
)

const (
    CNKI = "cnki.html"
)

func CnkiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, CNKI)
}

func main() {
    fmt.Println("starting unki web on localhost:8800")
    http.HandleFunc("/", CnkiHandler)
    http.ListenAndServe(":8800", nil)
}
