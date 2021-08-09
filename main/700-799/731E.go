package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF731E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, s int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}
	f := s
	for i := n - 1; i > 1; i-- {
		s -= a[i]
		f = max(f, s-f)
	}
	Fprint(out, f)
}

//func main() { CF731E(os.Stdin, os.Stdout) }
