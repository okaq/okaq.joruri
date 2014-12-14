// okaq vaki
// font engine
// bitmap font sampled from
// canvas fillText render
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	HTML_PATH = "vaki.html"
	MOTD      = "Hello. okaq vaki web app is live\nRunning on port: \n"
	PORT      = ":8082"
)

var (
	B []byte
)

func Cache() {
	var err error
	B, err = ioutil.ReadFile(HTML_PATH)
	if err != nil {
		log.Fatal(err)
	}
}

func Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.RemoteAddr, req.RequestURI, req.Method)
		w.Write(B)
	})
	http.ListenAndServe(PORT, nil)
}

func Greet() {
	fmt.Println(MOTD, PORT)
}

func main() {
	Greet()
	Cache()
	Serve()
}
