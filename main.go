package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
    "golang.org/x/tour/pic"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var x1, x2 int = 1, 1
	return func() int {
		x := x1
		x1 = x2
		x2 += x
		return x
	}
}

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(buf []byte) (int, error) {
	k, err := rot.r.Read(buf)
	for i := 0; i < k; i++ {
		c := buf[i]
		if 'a' <= c && c <= 'z' {
			c += 13
			if c > 'z' {
				c = 'a' + c - 'z' - 1
			}
			buf[i] = c
		}
		if 'A' <= c && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c = 'A' + c - 'Z' - 1
			}
			buf[i] = c
		}
	}
	return k, err
}

// Image implementation
type Image struct{
	w, h int    // Dimensions
	b [][] color.RGBAColor
}

func (Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
    return i.b[y][x]
}

func (i *Image) init(w, h int) {
    if w <= 0 || h <= 0 { return }
    i.w = w; i.h = h; i.b = make([][]color.RGBA, h)
    for y := 0 ; y < h ; y++ {
        i.b[y] = make([]color.RGBA, w)
        for x := 0 ; x < w ; x++ {
			i.b[y][x] = color.RGBA{uint8(3*x), uint8(5*y), uint8((x + y) / 2) | 32, 255}
        }
    }
}

func main() {
	fmt.Println("GOLang version: ", runtime.Version())
	fmt.Println("Două funcții (șiruri) Fibonacci:")
	f, g := fibonacci(), fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(), g())
	}
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

    m := Image{}
    m.init(86, 52)
    pic.showImage(m)
}
