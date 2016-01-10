package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

const (
    INDEX = "eian.html"
)

type User struct {
    A int64 // browser rand
    B int64 // server rand
    C int64 // nano time stamp
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok fetch!"))
    // marshal request json
    var d0 []int64
    dec := json.NewDecoder(r.Body)
    err := dec.Decode(&d0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(d0)
    w.Write([]byte("ok json fetch!"))
    // generate server random, timestamp, hash
    // send to browser
    // enc := json.NewEncoder(w)
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
