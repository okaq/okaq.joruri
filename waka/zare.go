// okaq.joruri web server
// waka mahjong tile bmps
// AQ <aq@okaq.com>
// 2014-12-30
package main

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
    "time"
)

const (
    HTML = "eago.html"
    PORT = ":8080"
    // FONT = "/home/ahmad/Documents/gira/okaq.joruri/fonts/Kalocsai_Flowers-webfont.woff"
    FONT = "/home/ahmad/Documents/gira/okaq.joruri/fonts/oxygen-mono-regular.woff"
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
    f0, err := os.Create(n0)
    if err != nil {
        fmt.Println(err)
    }
    defer f0.Close()
    // write json base64 encoded arraybuffer string to file
    s0 := string(j0)
    s1 := strings.NewReader(s0)
    w0, err := io.Copy(f0, s1)
    if err != nil {
        fmt.Println(err)
    }
    // direct copy - 512 bytes to binary file
    /*
    n0 := fmt.Sprintf("%s%s.bin", HANA, "zcopy_test_0")
    f0, err := os.Create(n0)
    if err != nil {
        fmt.Println(err)
    }
    defer f0.Close()
    w0, err := io.Copy(f0, req.Body)
    if err != nil {
        fmt.Println(err)
    }
    */
    d0 := fmt.Sprintf("Wrote %v bytes to file %s!\n", w0, n0)
    d1 := []byte(d0)
    w.Header().Set("Content-Type", "text/plain")
    // w.Write([]byte("file saved!"))
    w.Write(d1)
}

func main() {
    Now = time.Now()
    fmt.Printf("okaq.joruri waka starting on port%s\n", PORT)
    fmt.Printf("Started at: %s.\n", Now.Format(time.RFC1123Z))
    http.HandleFunc("/waka", WakaServer)
    // http.HandleFunc("/fonts/Kalocsai_Flowers-webfont.woff", FontServer)
    http.HandleFunc("/fonts/oxygen-mono-regular.woff", FontServer)
    http.HandleFunc("/save", SaveServer)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
