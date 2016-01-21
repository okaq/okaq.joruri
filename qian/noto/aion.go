// okaq joruri qian
// noto thumb render test
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "zion.html"
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
}

func main() {
    fmt.Println("okaq joruri qian noto thumb render test is on...")
    Pop()
    http.HandleFunc("/", NotoHandler)
    http.HandleFunc("/neto", NetoHandler)
    http.HandleFunc("/nito", NitoHandler)
    http.ListenAndServe(":8000", nil)
}
