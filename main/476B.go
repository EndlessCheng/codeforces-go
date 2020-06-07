package main

import (
	. "fmt"
	"io"
	. "strings"
)

// github.com/EndlessCheng/codeforces-go
func CF476B(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(in, &s, &t)
	p := Count(s, "+") - Count(s, "-") - Count(t, "+") + Count(t, "-")
	n := Count(t, "?")
	c := 0
	S := 1 << n
	for sub := 0; sub < S; sub++ {
		s := 0
		for i := 0; i < n; i++ {
			if sub>>i&1 == 0 {
				s--
			} else {
				s++
			}
		}
		if s == p {
			c++
		}
	}
	Fprintf(out, "%.12f", float64(c)/float64(S))
}

//func main() { CF476B(os.Stdin, os.Stdout) }
