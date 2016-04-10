// unki - patch based texture synth
// AQ <aq@okaq.com>
// 2016-04-12
package main

import (
    "fmt"
    "net/http"
)

const (
    DNKI = "dnki.html"
    SAMP = "sample"
)


func DnkiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, DNKI)
}

func main() {
    fmt.Println("starting unki web on localhost:8800")
    http.HandleFunc("/", DnkiHandler)
    http.ListenAndServe(":8800", nil)
}
