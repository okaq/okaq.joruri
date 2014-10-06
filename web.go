// okaq.joruri web server
// json bitmap draw tool
// execute: go run web.go
package main

import (
	"fmt"
	"log"
	"net/http"
    // "time"
)

const (
	HTML = "joruri.html"
	PORT = ":8008"
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
    w.Header().Set("Content-Length", "131072")
    // w.Header().Set("Doobie-Doo", "ItsaDoob!")
    w.WriteHeader(200)
    w.Write(big)
}

func main() {
	fmt.Printf("okaq.jouri starting on port %s\n", PORT)
    big = make([]byte, 1024*128)
    for i := 0; i < len(big); i++ {
        big[i] = byte(i % 255)
    }
	http.HandleFunc("/", JoruriServer)
    http.HandleFunc("/aa", OpenServer)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
