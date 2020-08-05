package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF978B(in io.Reader, out io.Writer) {
	var n, st, d int
	var s string
	Fscan(in, &n, &s)
	s = "$" + s + "$"
	for i, b := range s {
		if b != 'x' {
			st = i
		} else if s[i+1] != 'x' && i-st > 2 {
			d += i - st - 2
		}
	}
	Fprint(out, d)
}

//func main() { CF978B(os.Stdin, os.Stdout) }
