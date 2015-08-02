// okaq.joruri shima web server
// AQ <aq@okaq.com>
// 2015-08-04
package main

import (
    "fmt"
    "net/http"
)

const (
    LONI = "loni.html"
    PORT = ":8008"
)

func LoniHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, LONI)
}

func main() {
    fmt.Println("okaq joruri shima koni web server ok...")
    http.HandleFunc("/", LoniHandler)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
