// logo bitmap renderer
package main

import (
    "encoding/json"
	"fmt"
	"log"
	// "image"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

const (
	JSON_PATH = "okaq_logo_01.json"
	HTML_PATH = "unaki.html"
)

var (
	Html_Data []byte
    Json_Data []byte
)

type Unaki struct {
	B []byte // input data from JSON
	// pipe?
}

func (unaki *Unaki) Read([]byte) (n int, err error) {
	return 0, nil
}

func (unaki *Unaki) ReadInto() (n int, err error) {
	// read into Unaki.B
	return 0, nil
}

func genHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s,%s,%s\n", req.Method, req.URL, req.Host)
	// io.WriteString(w, "ok unaki genHandler")
	w.Write(Html_Data)
}

func disHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s,%s,%s\n", req.Method, req.URL, req.Host)
	if req.Method == "POST" {
		err := req.ParseMultipartForm(1024)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", req.Form["data01"])
		fmt.Printf("%s\n", req.RequestURI)
		h0 := req.Header
		for k, v := range h0 {
			fmt.Printf("%s:%s\n", k, v)
		}
		f0 := req.Form
		for k, v := range f0 {
			fmt.Printf("%s:%s\n", k, v)
		}
        b0, err := json.Marshal(f0)
        if err != nil {
            fmt.Printf("json err: %s\n", err)
        }
        fmt.Println(b0)
        Json_Data = b0
        Json()
		io.WriteString(w, "ok dis post\n")
		return
	}
	io.WriteString(w, "ok unaki disHandler\n")
}

func Servers() {
	http.HandleFunc("/gen", genHandler)
	http.HandleFunc("/dis", disHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Html() {
	var err error
	Html_Data, err = ioutil.ReadFile(HTML_PATH)
	if err != nil {
		log.Fatal(err)
	}
}

func Json() {
    err := ioutil.WriteFile(JSON_PATH, Json_Data, 0644)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	fmt.Println("ok unaki. bitmap renderer")
	n0 := runtime.NumCPU()
	fmt.Printf("Number CPUs: %d\n", n0)
	runtime.GOMAXPROCS(n0 + 1)
	fmt.Printf("GOMAXPROCS: %s\n", os.Getenv("GOMAXPROCS"))
	Html()
	Servers()
}

// read png, write json datauri
// gen png from html5 canvas2d

// websocket
// read input from chrome

// web input
// bitmap generator
// color bar
// click to save
// message tcp
// net.TCPConn
// port 8808
// paging
// bitmap sequences
// bitmap = [], len = 2^N
// bitseq = {0:bm0,1:bm1,2:bm2}

// wire format
// xhr send ArrayBufferView
// server side byte array (RGBA)
