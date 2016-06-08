package main

import (
	"io"
	"os"
    "fmt"
	"strings"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var x1, x2 int = 1, 1
	return func () int {
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
	for i := 0 ; i < k ; i++ {
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


func main() {
	f, g := fibonacci(), fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(), g())
	}
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
