package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "eian.html"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ok fetch!"))
    // marshal request json
    // generate server random, timestamp, hash
    // send to browser
}

func VianHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Println("okaq joruri bitmap draw tool started on localhost:8008")
    http.HandleFunc("/", VianHandler)
    http.HandleFunc("/user", UserHandler)
    http.ListenAndServe(":8008", nil)
}
