package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2074C(in io.Reader, out io.Writer) {
	var T, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		if x&1 == 0 {
			y = x&-x | 1
		} else {
			y = (x+1)&^x | 1
		}
		if y >= x {
			y = -1
		}
		Fprintln(out, y)
	}
}

//func main() { cf2074C(os.Stdin, os.Stdout) }
