// simple random bitmap generator
// AQ <aq@okaq.com>
// 2015-08-08
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "os"
    "path"
    "time"
)

var (
    N [16]string
    S rand.Source
    R *rand.Rand
)

func Names() {
    N = [16]string{"jake","nancy","star","pluto",
    "dance","macabre","violin","picolo",
    "destiny","fate","cosmic","dream",
    "chili","chile","chill","velociraptor"}
}

func Random() {
    S = rand.NewSource(time.Now().UnixNano())
    R = rand.New(S)
}

func RB(n int) []byte {
    r0 := make([]byte, n)
    for i := 0; i < n; i++ {
        b0 := R.Intn(255)
        r0[i] = byte(b0)
    }
    return r0
}

func main() {
    fmt.Println("Bitmap generation program starting..")
    Names()
    Random()
    fmt.Printf("Random byte test: %v.\n", RB(1))
    fmt.Printf("Bitmap size %dx%d = %d bits = %d bytes.\n", 32, 32, 32*32, 32*32/8)
    wd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("Current dir: %s.\n", wd)
    p0 := "../cata/bata/"
    p1 := path.Join(wd, p0)
    fmt.Println("Write to dir: %s.\n", p1)
    for i, n := range N {
        fmt.Printf("Name %d is %s.\n", i, n)
        b0 := RB(128)
        fmt.Println(b0)
        // file path from name
        // pwd current path
        // bmp dir ../cata/bata
        p2 := path.Join(p1, n)
        p3 := p2 + ".bin"
        fmt.Printf("Writing data to: %s.\n", p3)
        ioutil.WriteFile(p3, b0, 0664)
        // change chmod 0664
        // analogous json base64 string
        p4 := p2 + ".json"
        fmt.Printf("Writing data to: %s.\n,", p4)
        j0, err := json.Marshal(b0)
        if err != nil {
            fmt.Println(err)
        }
        ioutil.WriteFile(p4, j0, 0644)
    }
    // ioutil.WriteFiie rand bytes
}
