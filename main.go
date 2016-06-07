package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice("s", s)

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[0:2]
	printSlice("c", c)

	d := c[2:4]
	printSlice("d", d)

    e := a[0:2]
    printSlice("e", e)

    f := e[2:4]
    printSlice("f", f)

    d[0] = 1; f[0] = 2;
    fmt.Println(a, b, c, d, e, f)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s %T len=%d cap=%d %v\n",
		s, x, len(x), cap(x), x)
}
