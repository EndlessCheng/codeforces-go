package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2056C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprint(out, "1 1 ")
		for i := 2; i < n-1; i++ {
			Fprint(out, i, " ")
		}
		Fprintln(out, 1)
	}
}

//func main() { cf2056C(os.Stdin, os.Stdout) }
