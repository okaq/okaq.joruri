// golang web server
// unki - patch based texture synth
// okaq joruri bitmap draw tools
// AQ <aq@okaq.com>
// 2016-03-12
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

const (
    ANKI = "anki.html"
    // sample and save dirs
    SAMPLE = "sample/"
    SAVE = "save/"
)

var (
    SamplePaths []string
    // cached json dirs
)

func AnkiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, ANKI)
}

func OpenHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ok bitmaps!"))
    // ioutil.Readdir returns []os.FileInfo
    // use len() to get dir size
    // json encoding to output
}

func Paths() []string {
    f0, err := ioutil.ReadDir(SAMPLE)
    if err != nil {
        fmt.Println(err)
    }
    s0 := make([]string, len(f0))
    for i0, f1 := range f0 {
        if f1.IsDir() {
            // dirs should not contain sub dirs
            continue
        }
        // s0 = append(s0, f1.Name())
        s0[i0] = f1.Name()
    }
    return s0
}

func main() {
    fmt.Println("starting unki web on localhost:8800")
    SamplePaths = Paths()
    fmt.Println(SamplePaths, len(SamplePaths))
    http.HandleFunc("/", AnkiHandler)
    http.HandleFunc("/open", OpenHandler)
    http.ListenAndServe(":8800", nil)
}
