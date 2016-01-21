// okaq joruri qian
// noto thumb render test
package main

import (
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
)

func NotoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func NetoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ok neto"))
}

func NitoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ok nito"))
}


func Pop() {
    // pre populate dir lists
    Neto = Paths("../neto")
    Nito = Paths("../nito")
    fmt.Println(Neto, Nito)
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
    http.ListenAndServe(":8000", nil)
}
