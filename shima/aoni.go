// okaq.joruri web server
package main

import (
    "bufio"
    "fmt"
    "os"
    "net/http"
    "time"
)

const (
    BONI = "boni.html"
    PORT = ":8008"
)

func BoniHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req)
    http.ServeFile(w, req, BONI)
}

func motd() {
    fmt.Println("The joy of a summer sunrise glows within the soul for eternity! - Emerson (1838)")
    r0 := bufio.NewReader(os.Stdin)
    fmt.Println("Who are you? Please type in your name and press <Enter>")
    t0, err := r0.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }
    // strip trailing newline
    t0 = t0[:len(t0)-1]
    fmt.Printf("Greetings %s! Welcome to the okaq.joruri bitmap importer\n", t0)
}

func main() {
    now := time.Now()
    fmt.Printf("okaq.joruri shima aoni web app running on port %s\n", PORT)
    fmt.Printf("Started at: %s\n", now.Format(time.RFC1123Z))
    motd()
    http.HandleFunc("/boni", BoniHandler)
    err := http.ListenAndServe(PORT, nil)
    if err != nil {
        fmt.Println(err)
    }
}
