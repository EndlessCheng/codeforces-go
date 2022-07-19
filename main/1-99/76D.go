package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF76D(in io.Reader, out io.Writer) {
	var a, b uint64
	Fscan(in, &a, &b)
	if a < b || a&1 != b&1 {
		Fprint(out, -1)
	} else {
		x := (a - b) / 2
		Fprint(out, x, a-x)
	}
}

//func main() { CF76D(os.Stdin, os.Stdout) }
