// simple random bitmap generator
// AQ <aq@okaq.com>
// 2015-08-08
package main

import (
    "fmt"
)

var (
    N [16]string
)

func Names() {
    N = [16]string{"jake","nancy","star","pluto",
    "dance","macabre","violin","picolo",
    "destiny","fate","cosmic","dream",
    "chili","chile","chill","velociraptor"}
}

func main() {
    fmt.Println("Bitmap generation program starting..")
    Names()
    for i, n := range N {
        fmt.Printf("Name %d is %s.\n", i, n)
    }
}
