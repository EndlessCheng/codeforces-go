package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2415(in io.Reader, out io.Writer) {
	var c, v, s int
	for {
		n, _ := Fscan(in, &v)
		if n == 0 {
			break
		}
		s += v
		c++
	}
	Fprint(out, s<<(c-1))
}

//func main() { p2415(os.Stdin, os.Stdout) }
