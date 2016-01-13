package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "os"
    "time"
)

const (
    INDEX = "eian.html"
)

var (
    U *User
    // map of user hashes to data
    V *bufio.Reader
)

type User struct {
    A int64 // browser rand
    B int64 // server rand
    C int64 // nano time stamp
    R *rand.Rand
    S rand.Source
    T int64 // seed from time now
    // hash
}

func NewUser() *User {
    // init rand
    u := new(User)
    u.T = time.Now().UnixNano()
    u.S = rand.NewSource(u.T)
    u.R = rand.New(u.S)
    return u
}

func (u *User) Next(a int64) {
    u.A = a
    u.B = u.R.Int63()
    u.C = time.Now().UnixNano()
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok fetch!"))
    // marshal request json
    var d0 []int64
    dec := json.NewDecoder(r.Body)
    err := dec.Decode(&d0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(d0)
    // user data
    U.Next(d0[0])
    fmt.Println(U.B,U.C)
    w.Write([]byte("ok json fetch!"))
    // generate server random, timestamp, hash
    // send to browser
    // enc := json.NewEncoder(w)
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {
    // respond with www path to img src 
    // use fileserver to serve img
    fmt.Println(r)
    fmt.Println("Enter img src: ");
    s0, err := V.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }
    // w.Write([]byte("ok load!"))
    w.Write([]byte(s0))
}

func VianHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Println("okaq joruri bitmap draw tool started on localhost:8008")
    U = NewUser()
    V = bufio.NewReader(os.Stdin)
    http.Handle("/nato/", http.StripPrefix("/nato", http.FileServer(http.Dir("nato/"))))
    http.HandleFunc("/", VianHandler)
    http.HandleFunc("/user", UserHandler)
    http.HandleFunc("/load", LoadHandler)
    http.ListenAndServe(":8008", nil)
}
