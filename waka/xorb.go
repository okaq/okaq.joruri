// okaq.joruri web server
// waka card assembler tool
// AQ <aq@okaq.com>
// 2014-12-24
package main

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"
)

const (
    HTML = "duka.html"
    PORT = ":8080"
    FONT = "/home/ahmad/Documents/gira/okaq.joruri/fonts/Kalocsai_Flowers-webfont.woff"
    HANA = "/home/ahmad/Documents/gira/okaq.joruri/hana/"
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
    // get file name from command prompt
    r0 := bufio.NewReader(os.Stdin)
    fmt.Println("What file name should I assign? Please input below and hit <Enter>")
    t0, err := r0.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }
    t0 = t0[:len(t0)-1] // remove trailing newline
    n0 := fmt.Sprintf("%s%s.json", HANA, t0)
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
