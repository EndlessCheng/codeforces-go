package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF388C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	a := []int{}
	var n, m, v, x, y int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &m)
		for i := 0; i < m/2; i++ {
			Fscan(in, &v)
			x += v
		}
		if m&1 > 0 {
			Fscan(in, &v)
			a = append(a, v)
			m--
		}
		for i := m / 2; i < m; i++ {
			Fscan(in, &v)
			y += v
		}
	}
	sort.Ints(a)
	for i := len(a) - 1; i >= 0; i -= 2 {
		x += a[i]
	}
	for i := len(a) - 2; i >= 0; i -= 2 {
		y += a[i]
	}
	Fprint(out, x, y)
}

//func main() { CF388C(os.Stdin, os.Stdout) }
