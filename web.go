// okaq.joruri web server
// json bitmap draw tool
// execute: go run web.go
package main

import (
    "encoding/json"
	"fmt"
	"log"
	"net/http"
    "os"
    "strings"
    "time"
)

const (
    HTML = "mota.html"
	// HTML = "joruri.html"
	PORT = ":8080"
)

var (
    big []byte
)

func JoruriServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	http.ServeFile(w, req, HTML)
}

func OpenServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    // time simulate network lag
    // hi := []byte("hi there from AServer\n")
    // w.Write(hi)
    /*
    for i := 0; i < 8; i++ {
        time.Sleep(1 * time.Second)
        mess := []byte(fmt.Sprintf("this is message number: %d.", i))
        w.Write(mess)
    }
    */
    // w.Header().Set("Content-Length", "131072")
    // w.Header().Set("Doobie-Doo", "ItsaDoob!")
    w.WriteHeader(200)
    w.Write(big)
}

type Bitmap struct {
    Greyscales int
    Grids []int
}

func SaveHandler(w http.ResponseWriter, req *http.Request) {
    /*
    var b []byte
    var err error
    _, err = req.Body.Read(b)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(b))
    err = req.ParseForm()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(req.Form)
    for k := range req.Header {
        fmt.Println(k)
        fmt.Println(req.Form[k])
    }
    w.Write([]byte("json file stored"))
    */
    var err error
    dec := json.NewDecoder(req.Body)
    var bm Bitmap
    err = dec.Decode(&bm)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(bm)
    // timestamp
    now := time.Now()
    stamp := now.Format(time.RFC3339)
    filename := "data/" + stamp + ".json"
    fmt.Println(filename)
    // os file
    var f0 *os.File
    f0, err = os.Create(filename)
    if err != nil {
        fmt.Println(err)
    }
    // json encoder
    enc := json.NewEncoder(f0)
    enc.Encode(bm)
    // cleanup
    err = f0.Close()
    if err != nil {
        fmt.Println(err)
    }
    w.Write([]byte("json file decoded and stored"))
}

func TestHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    // w.Write([]byte("[\"test\",\"ok\"]"))
    const s0 = `{"Id": "test", "Data": "ok!"}`
    type Test struct {
        Id, Data string
    }
    s1 := strings.NewReader(s0)
    dec := json.NewDecoder(s1)
    enc := json.NewEncoder(w)
    var t0 Test
    var err error
    err = dec.Decode(&t0)
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println(t0)
    err = enc.Encode(&t0)
    if err != nil {
        fmt.Println(err)
    }
}

// view archive
// on init, pre load data filenames

func main() {
	fmt.Printf("okaq.jouri starting on port %s\n", PORT)
    big = make([]byte, 255)
    for i := 0; i < len(big); i++ {
        big[i] = byte(i % 255)
    }
	http.HandleFunc("/", JoruriServer)
    http.HandleFunc("/aa", OpenServer)
    http.HandleFunc("/ab", SaveHandler)
    http.HandleFunc("/test", TestHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
