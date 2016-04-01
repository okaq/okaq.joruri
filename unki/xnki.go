// unki - patch based texture synth
// AQ <aq@okaq.com>
// 2016-03-26
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

const (
    CNKI = "cnki.html"
    SAMP = "sample"
)

var (
    S []string // samples list
)

func CnkiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, CNKI)
}

func Samples() {
    f0, err := ioutil.ReadDir(SAMP)
    if err != nil {
        fmt.Println(err)
    }
    S = make([]string, len(f0))
    for i, f1 := range f0 {
        // fmt.Println(f1.Name())
        S[i] = f1.Name()
    }
    fmt.Println(S)
}

func main() {
    fmt.Println("starting unki web on localhost:8800")
    Samples()
    http.HandleFunc("/", CnkiHandler)
    http.ListenAndServe(":8800", nil)
}

// populate and cache dir of sample patch images
// bitmap format is png, 256x256 or similar 2^N
// serve files directly with appropriate header
