// okaq.joruri shima web server
// AQ <aq@okaq.com>
// 2015-07-26
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "net/http"
)

const (
    HONI = "honi.html"
    PORT = ":8008"
    CATA = "../cata"
)

var (
    Fi []os.FileInfo
)

func HoniHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, HONI)
}

func FileHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    w.Write([]byte("ok select file"))
    // list available files in dir
    // input select via command line
}

func Load() (int, error) {
    var err error
    Fi, err = ioutil.ReadDir(CATA)
    if err != nil {
        return 0, err
    }
    return len(Fi), nil
}

func main() {
    fmt.Println("okaq joruri shima goni web server is live...")
    fmt.Println("loading cata local bitmap file store from disk")
    i0, err0 := Load()
    if err0 != nil {
        fmt.Println(err0)
    }
    fmt.Printf("%d files loaded from %s.\n", i0, CATA)
    http.HandleFunc("/", HoniHandler)
    http.HandleFunc("/file", FileHandler)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
