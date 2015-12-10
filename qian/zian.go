package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "aian.html"
)

func ZianHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Println("okaq joruri bitmap tool started on localhost:8008")
    http.HandleFunc("/", ZianHandler)
    http.ListenAndServe(":8008", nil)
}
