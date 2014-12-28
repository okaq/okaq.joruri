// okaq.joruri waka stitch
// piece together hana bmp data
// into single json object
// data is encoded as 
// base64 arraybuffer string
package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"
)

const (
    PATH = "../hana2"
    OUT = "../hana2/hana2d.js"
)

var (
    Now time.Time
    Output *os.File
    W *bufio.Writer
)

func OutputFile() {
    var err error
    Output, err = os.Create(OUT)
    if err != nil {
        fmt.Println(err)
    }
    W = bufio.NewWriter(Output)
}

func Header() {
    s0 := fmt.Sprintf("var hana2d = {\n")
    W.WriteString(s0)
}

func Footer() {
    s0 := fmt.Sprintf("};\n")
    W.WriteString(s0)
    W.Flush()
}

func StitchWalk(path string, info os.FileInfo, err error) error {
    fmt.Println(path, info, err)
    if info.IsDir() {
        fmt.Println("Is a dir")
    } else {
        /*
        f0, err0 := os.Open(path)
        if err0 != nil {
            return err0
        }
        */
        fmt.Printf("writing %s. %v bytes", path, info.Size())
        b0, err0 := ioutil.ReadFile(path)
        if err0 != nil {
            return err0
        }
        s0 := string(b0)
        // s1 := fmt.Sprintf("%s,\n", s0)
        s1 := info.Name()
        s1 = s1[:len(s1)-5]
        s2 := fmt.Sprintf("%s:%s,\n", s1, s0)
        W.WriteString(s2)
    }
    return err
}

func main() {
    Now = time.Now()
    fmt.Println("stitching okaq.joruri hana bmp data")
    fmt.Printf("started at: %s\n", Now.Format(time.RFC3339))
    fmt.Printf("creating output file: %s\n", OUT)
    OutputFile()
    defer Output.Close()
    Header()
    fmt.Printf("walking path at: %s\n", PATH)
    err := filepath.Walk(PATH, StitchWalk)
    if err != nil {
        fmt.Println(err)
    }
    Footer()
}
