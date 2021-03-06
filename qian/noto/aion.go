// okaq joruri qian
// noto thumb render test
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const (
    INDEX = "zion.html"
)

var (
    Neto []string
    Nito []string
    NetoB []byte
    NitoB []byte
)

func NotoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func NetoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok neto"))
    w.Header().Set("Content-type", "application/json")
    w.Write(NetoB)
}

func NitoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok nito"))
    w.Header().Set("Content-type", "application/json")
    w.Write(NitoB)
}
// handlers for bitmap json files
// direct load into byte array and
// write as byte array json mime type
// no need for enc/decode

func Pop() {
    // pre populate dir lists
    Neto = Paths("../neto")
    Nito = Paths("../nito")
    fmt.Println(Neto, Nito)
    // json encode
    var err error
    NetoB, err = json.Marshal(Neto)
    if err != nil {
        fmt.Println(err)
    }
    NitoB, err = json.Marshal(Nito)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(NetoB, NitoB)
}

func Paths(d0 string) []string {
    f0, err := ioutil.ReadDir(d0)
    if err != nil {
        fmt.Println(err)
    }
    s0 := make([]string, 1)
    s0[0] = "files"
    for _, f1 := range f0 {
        if f1.IsDir() {
            continue
        }
        s0 = append(s0, f1.Name())
    }
    return s0[1:]
}


func main() {
    fmt.Println("okaq joruri qian noto thumb render test is on...")
    Pop()
    http.HandleFunc("/", NotoHandler)
    http.HandleFunc("/neto", NetoHandler)
    http.HandleFunc("/nito", NitoHandler)
    http.Handle("/draw/", http.StripPrefix("/draw/", http.FileServer(http.Dir("../neto"))))
    http.Handle("/import/", http.StripPrefix("/import/", http.FileServer(http.Dir("../nito"))))
    http.ListenAndServe(":8000", nil)
}
