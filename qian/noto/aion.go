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

func main() {
    fmt.Println("okaq joruri qian noto thumb render test is on...")
    http.HandleFunc("/", NotoHandler)
    http.ListenAndServe(":8000", nil)
}
