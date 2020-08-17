package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1358E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	var n2, x, s int64
	Fscan(in, &n2)
	n := (n2 + 1) / 2
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &x)

	mins := make([]int64, n+1)
	for i, v := range a {
		s += x - v
		mins[i+1] = min(mins[i], s)
	}
	s = 0
	for _, v := range a {
		s += v
	}
	for k := n; k <= n2; k++ {
		if s+mins[n2-k] > 0 {
			Fprint(out, k)
			return
		}
		s += x
	}
	Fprint(out, -1)
}

//func main() { CF1358E(os.Stdin, os.Stdout) }
