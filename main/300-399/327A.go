package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf327A(in io.Reader, out io.Writer) {
	var n, v, f0 int
	f1, f2 := int(-1e9), int(-1e9)
	Fscan(in, &n)
	for range n {
		Fscan(in, &v)
		f2 = max(f2, f1) + v
		f1 = max(f1, f0) + 1 - v
		f0 += v
	}
	Fprint(out, max(f1, f2))
}

//func main() { cf327A(os.Stdin, os.Stdout) }
