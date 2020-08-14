package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1183A(_r io.Reader, _w io.Writer) {
	var n int
	for Fscan(_r, &n); ; n++ {
		s := 0
		for x := n; x > 0; x /= 10 {
			s += x % 10
		}
		if s%4 == 0 {
			Fprint(_w, n)
			return
		}
	}
}

//func main() { CF1183A(os.Stdin, os.Stdout) }
