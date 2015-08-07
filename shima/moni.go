// simple random bitmap generator
// AQ <aq@okaq.com>
// 2015-08-08
package main

import (
    "fmt"
    "math/rand"
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
    for i, n := range N {
        fmt.Printf("Name %d is %s.\n", i, n)
        b0 := RB(128)
        fmt.Println(b0)
        // file path from name
        // pwd current path
        // bmp dir ../cata/bata
    }
    // ioutil.WriteFiie rand bytes
}
