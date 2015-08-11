// okaq.joruri shima web server
// AQ <aq@okaq.com>
// 2015-08-04
package main

import (
    // "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path"
)

const (
    LONI = "loni.html"
    PORT = ":8008"
)

var (
    P string
    L []byte
)

func LoniHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, LONI)
}

func ListHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    w.Write(L)
}

func BitmapHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    // fmt.Println(req.Body)
    // s0 := bufio.NewScanner(req.Body)
    // b0 := s0.Bytes()
    b0 := new(bytes.Buffer)
    b0.ReadFrom(req.Body)
    b1 := b0.Bytes()
    s0 := string(b1)
    fmt.Println(b1, s0)
    w.Write([]byte("bitmap ok"))
    // read base64 file json and write to response
    // byte stream pipe direct to writer
}

func Path() {
    wd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }
    p0 := "../cata/bata/"
    P = path.Join(wd, p0)
}

func List() {
    f0, err := ioutil.ReadDir(P)
    if err != nil {
        fmt.Println(err)
    }
    s0 := make([]string, len(f0))
    for i, f1 := range f0 {
        // fmt.Printf("File %d is: %s.\n", i, f1.Name())
        s0[i] = f1.Name()
    }
    b0, err := json.Marshal(s0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(b0)
    L = b0
}

func main() {
    fmt.Println("okaq joruri shima koni web server ok...")
    http.HandleFunc("/", LoniHandler)
    // list dir contents
    Path()
    fmt.Printf("Reading bitmap files from dir: %s.\n", P)
    List()
    fmt.Println("Bitmap file list created.")
    // file list json array [a.bin,b.bin]
    // bitmap json request matches name and serves binary
    http.HandleFunc("/a", ListHandler)
    http.HandleFunc("/b", BitmapHandler)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
