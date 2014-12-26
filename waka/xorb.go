// okaq.joruri web server
// waka card assembler tool
// AQ <aq@okaq.com>
// 2014-12-24
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

const (
    HTML = "duka.html"
    PORT = ":8080"
    FONT = "/home/ahmad/Documents/gira/okaq.joruri/fonts/Kalocsai_Flowers-webfont.woff"
)

var (
    Now time.Time
)

func WakaServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, HTML)
}

func FontServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    // octet-stream more portable than font-woff 
    w.Header().Set("Content-Type", "application/octet-stream")
    http.ServeFile(w, req, FONT)
}

func SaveServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    b0 := new(bytes.Buffer)
    b0.ReadFrom(req.Body)
    j0, err := json.Marshal(b0.Bytes())
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(j0))
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("file saved!"))
}

func main() {
    Now = time.Now()
    fmt.Printf("okaq.joruri waka starting on port%s\n", PORT)
    fmt.Printf("Started at: %s.\n", Now.Format(time.RFC1123Z))
    http.HandleFunc("/waka", WakaServer)
    http.HandleFunc("/fonts/Kalocsai_Flowers-webfont.woff", FontServer)
    http.HandleFunc("/save", SaveServer)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
