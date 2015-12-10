// favicon encoder: 32x32 png and json base64 data uri formats
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

var (
	Png      string
	Json     string
	PngFile  *os.File
	JsonFile *os.File
	Pixels   *image.RGBA
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

func Paint() {
	// c0 := color.RGBA{200, 0, 0, 255}
	// royal blue - bitstarter
    // c0 := color.RGBA{65,105,255,255}
    // white - okaq-funka - no alpha<255
    // c0 := color.RGBA{242,242,242,255}
    // okaq unaki : logo bitmap renderer
    // c0 := color.RGBA{233,0,223,255}
    // okaq vaki : font engine
    // c0 := color.RGBA{135,200,240,255}
    // c0 := color.RGBA{190,0,200,255}
    // okaq vobu boru
    // c0 := color.RGBA{9,232,176,255}
    // okaq vobu foti
    // c0 := color.RGBA{176,9,232,255}
    // okaq toru - grayish pink
    // c0 := color.RGBA{221,177,232,255}
    // okaq gira bibi
    // c0 := color.RGBA{84,167,255,255}
    // okaq gira cobu
    // c0 := color.RGBA{167,84,255,255}
    // okaq gira daki
    // c0 := color.RGBA{205,255,128,255}
    // okaq.joruri
    // c0 := color.RGBA{255,230,30,255}
    // okaq.github.io
    // c0 := color.RGBA{255,89,140,255}
    // okaq gira fure
    // c0 := color.RGBA{100,255,200,255}
    // okaq joruri qian
    c0 := color.RGBA{127,196,245,255}
    s0 := Pixels.Bounds().Dx() * Pixels.Bounds().Dy()
	for i := 0; i < s0; i++ {
		x0 := i % Pixels.Bounds().Dx()
		y0 := i / Pixels.Bounds().Dy()
		Pixels.SetRGBA(x0, y0, c0)
	}
	fmt.Println(Pixels.At(0, 0))
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
	// fmt.Println(fi0.Size())
	d1 := make([]byte, fi0.Size())
	_, err = PngFile.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	_, err = PngFile.Read(d1)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(d1)
	// s0 := base64.URLEncoding.EncodeToString(Pixels.Pix)
	s0 := base64.StdEncoding.EncodeToString(d1)
	// fmt.Println(s0)
	s1 := "data:image/png;base64,"
	s2 := fmt.Sprintf("%s%s", s1, s0)
	// fmt.Println(s2)
	d0 := Favio{DataUri: s2}
	b0, err := json.Marshal(d0)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(b0)
	_, err = JsonFile.Write(b0)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	Path()
	Open()
	defer PngFile.Close()
	defer JsonFile.Close()
	Image()
	Paint()
	Write()
}
