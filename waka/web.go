// okaq.joruri web server
// waka image render and sampling tool
// AQ <aq@okaq.com>
// 2014-12-16
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    // "log"
    "math/rand"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

const (
    // HTML = "aomi.html"
    HTML = "bome.html"
    PORT = ":8080"
    FONT = "/home/ahmad/Documents/gira/okaq.joruri/fonts/Kalocsai_Flowers-webfont.woff"
)

var (
    Now time.Time
    Source rand.Source
    Rng *rand.Rand
)

func WakaServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, HTML)
}

func FontServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    w.Header().Set("Content-Type", "application/font-woff")
    // w.Header().Set("Content-Type", "text/css")
    http.ServeFile(w, req, FONT)
}

// save - xhr conn json disk save
func SaveHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    fmt.Println(req.Body)
    /*
    type ByteString struct {
        Index, Value string
    }
    */
    // dec := json.NewDecoder(req.Body)
    /*
    bs0 := make([]ByteString, 0)
    for {
        var bs1 ByteString
        err := dec.Decode(&bs1)
        bs0 = append(bs0, bs1)
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
    }
    fmt.Println(bs0)
    type ByteString map[string]string
    var bs0 ByteString
    err := dec.Decode(&bs0)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(bs0)
    */
    b0 := new(bytes.Buffer)
    b0.ReadFrom(req.Body)
    fmt.Println(b0.Bytes())
    j0, err := json.Marshal(b0.Bytes())
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(j0))
    /*
    dec := json.NewDecoder(req.Body)
    type BmpData struct {
        Data []interface{}
    }
    var bd0 map[string]interface{}
    err := dec.Decode(&bd0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(bd0)
    */
    s0 := string(j0)
    s1 := strings.NewReader(s0)
    f0, err := os.Create("test.json")
    if err != nil {
        fmt.Println(err)
    }
    defer f0.Close()
    n0, err := io.Copy(f0, s1)
    if err != nil {
        fmt.Println(n0, err)
    }
    Shaker(w, "waka")
}

// handshake - xhr connection test handler
func ShakeServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    // w.Header().Set("Content-Type", "application/json")
    Shaker(w, "hanafuda")
}

func Shaker(w http.ResponseWriter, m0 string) {
    w.Header().Set("Content-Type", "application/json")
    type Shake struct {
        Id, Now, Cipher string
    }
    s0 := new(Shake)
    r0 := Rng.Int()
    s0.Id = "user_" + strconv.Itoa(r0)
    t0 := time.Now().UnixNano()
    s0.Now = strconv.FormatInt(t0, 16)
    s0.Cipher = m0
    fmt.Println(s0)
    b, err := json.Marshal(s0)
    if err != nil {
        fmt.Println(err)
    }
    w.Write(b)
}

func main() {
    Now = time.Now()
    fmt.Printf("okaq.joruri waka starting on port%s\n", PORT)
    fmt.Printf("Started at: %s.\n", Now.Format(time.RFC1123Z))
    Source = rand.NewSource(Now.UnixNano())
    Rng = rand.New(Source)
    http.HandleFunc("/waka", WakaServer)
    http.HandleFunc("/fonts", FontServer)
    // http.Handle("/fonts", http.FileServer(http.Dir("/home/ahmad/Documents/gira/okaq.joruri/fonts")))
    // http.Handle("/", http.FileServer(http.Dir("/home/ahmad/Documents/gira/okaq.joruri")))
    http.HandleFunc("/shake", ShakeServer)
    http.HandleFunc("/save", SaveHandler)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
