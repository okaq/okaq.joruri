// favicon encoder: 32x32 png and json base64 data uri formats
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
    "math/rand"
	"os"
	"time"
)

var (
	Png      string
	Json     string
	PngFile  *os.File
	JsonFile *os.File
	Pixels   *image.RGBA
    Source   rand.Source
    Rng      *rand.Rand
)

type Favio struct {
	DataUri string
}

func Filepath(p, s string) string {
	now := time.Now().UnixNano()
	r := fmt.Sprintf("%s%d%s", p, now, s)
	return r
}

func Path() {
	// create filepath string from UnixNow
	// pref := "favicon_"
	// suf := ".png"
	// now := time.Now().UnixNano()
	// Name = fmt.Sprintf("%s%d%s", pref, now, suf)
	// fmt.Println(Name)
	Png = Filepath("favicon_", ".png")
	Json = Filepath("dataUri_", ".json")
	// fmt.Println(Png, Json)
}

func Open() {
	var err error
	PngFile, err = os.Create(Png)
	if err != nil {
		fmt.Println(err)
	}
	JsonFile, err = os.Create(Json)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(PngFile, JsonFile)
}

func Image() {
	rect := image.Rect(0, 0, 32, 32)
	Pixels = image.NewRGBA(rect)
}

func PaintBG() {
    // c0 := color.RGBA{100,255,200,255}
    c0 := RC()
    s0 := Pixels.Bounds().Dx() * Pixels.Bounds().Dy()
	for i := 0; i < s0; i++ {
		x0 := i % Pixels.Bounds().Dx()
		y0 := i / Pixels.Bounds().Dy()
		Pixels.SetRGBA(x0, y0, c0)
	}
	// fmt.Println(Pixels.At(0, 0))
}

func PaintEyes() {
    c0 := RC()
    // block size for 32x32 image is 8x8
    s0 := 64
    for i := 0; i < s0; i++ {
        x0 := i % 8
        x1 := x0 + 4
        x2 := x0 + 20
        y0 := i / 8
        y1 := y0 + 4
        Pixels.SetRGBA(x1, y1, c0)
        Pixels.SetRGBA(x2, y1, c0)
    }
}

func PaintMouth() {
    c0 := RC()
    s0 := 64
    for i := 0; i < s0; i++ {
        x0 := i % 8
        x1 := x0 + 12
        y0 := i / 8
        y1 := y0 + 20
        Pixels.SetRGBA(x1, y1, c0)
    }
}

func Paint() {
    PaintBG()
    PaintEyes()
    PaintMouth()
}

// Random color
func RC() color.RGBA {
    r0 := RB()
    g0 := RB()
    b0 := RB()
    return color.RGBA{r0,g0,b0,255}
}

// Random byte
func RB() uint8 {
    // s0 := rand.NewSource(time.Now().UnixNano())
    // r0 := rand.New(s0)
    r1 := Rng.Intn(256)
    return uint8(r1)
}

func InitRand() {
    Source = rand.NewSource(time.Now().UnixNano())
    Rng = rand.New(Source)
}

func Write() {
	err := png.Encode(PngFile, Pixels)
	if err != nil {
		fmt.Println(err)
	}
	// following line encodes pixel data only, not full png file
	fi0, err := PngFile.Stat()
	if err != nil {
		fmt.Println(err)
	}
	d1 := make([]byte, fi0.Size())
	_, err = PngFile.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	_, err = PngFile.Read(d1)
	if err != nil {
		fmt.Println(err)
	}
	s0 := base64.StdEncoding.EncodeToString(d1)
	s1 := "data:image/png;base64,"
	s2 := fmt.Sprintf("%s%s", s1, s0)
	d0 := Favio{DataUri: s2}
	b0, err := json.Marshal(d0)
	if err != nil {
		fmt.Println(err)
	}
	_, err = JsonFile.Write(b0)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
    InitRand()
	Path()
	Open()
	defer PngFile.Close()
	defer JsonFile.Close()
	Image()
	Paint()
	Write()
}
