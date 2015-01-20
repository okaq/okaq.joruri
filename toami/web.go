// okaq joruri web server
// toami unit cell pattern drawing tool
// AQ <aq@okaq.com>
// 2015-01-20
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    HTML = "bora.html"
    PORT = ":8080"
)

var (
    Now time.Time
)

func ToamiServer(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, HTML)
}

func main() {
    Now = time.Now()
    fmt.Printf("okaq.joruri toami starting on port %s\n", PORT)
    fmt.Printf("at: %s\n", Now.Format(time.RFC1123Z))
    http.HandleFunc("/toami", ToamiServer)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
