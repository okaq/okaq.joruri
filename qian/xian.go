package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "cian.html"
)

func XianHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Println("okaq joruri bitmap draw tool started on localhost:8008")
    http.HandleFunc("/", XianHandler)
    http.ListenAndServe(":8008", nil)
}
