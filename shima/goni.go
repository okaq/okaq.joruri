// okaq.joruri shima web server
// AQ <aq@okaq.com>
// 2015-07-26
package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "net/http"
    "strings"
)

const (
    HONI = "honi.html"
    PORT = ":8008"
    CATA = "../cata"
)

var (
    Fi []os.FileInfo
    R *bufio.Reader
    B string
)

func HoniHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, HONI)
}

func FileHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    // w.Write([]byte("ok select file"))
    // list available files in dir
    for _, f := range Fi {
        fmt.Printf("Name: %s. Size: %d. ModTime: %v.\n", f.Name(), f.Size(), f.ModTime())
    }
    // input select via command line
    fmt.Println("Enter file name to download: ")
    t, err := R.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }
    // create path to send to req
    B = "/cata/"
    // B += CATA
    // B += "/"
    t = strings.TrimSpace(t)
    // t = t[:len(t)-4]
    B += t
    fmt.Printf("Sending %s to request.\n", B)
    w.Write([]byte(B))
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
    R = bufio.NewReader(os.Stdin)
    i0, err0 := Load()
    if err0 != nil {
        fmt.Println(err0)
    }
    fmt.Printf("%d files loaded from %s.\n", i0, CATA)
    http.HandleFunc("/", HoniHandler)
    http.HandleFunc("/file", FileHandler)
    fs := http.FileServer(http.Dir(CATA))
    http.Handle("/cata/", http.StripPrefix("/cata/",fs))
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
