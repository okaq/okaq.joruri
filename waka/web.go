// okaq.joruri web server
// waka image render and sampling tool
// AQ <aq@okaq.com>
// 2014-12-16
package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "strconv"
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

// handshake - xhr connection test handler
func ShakeServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    w.Header().Set("Content-Type", "application/json")
    type Shake struct {
        Id, Now, Cipher string
    }
    s0 := new(Shake)
    r0 := Rng.Int()
    s0.Id = "user_" + strconv.Itoa(r0)
    t0 := time.Now().UnixNano()
    s0.Now = strconv.FormatInt(t0, 16)
    s0.Cipher = "hanafuda"
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
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
