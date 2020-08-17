package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1303B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	searchRange := func(l, r int64, f func(int64) bool) int64 {
		i, j := l, r
		for i < j {
			h := (i + j) >> 1
			if f(h) {
				j = h
			} else {
				i = h + 1
			}
		}
		return i
	}

	solve := func() (ans int64) {
		var n, g, b int64
		Fscan(in, &n, &g, &b)
		sum := g + b
		half := (n-1)/2 + 1
		return searchRange(n, 1e18, func(days int64) bool {
			seg, left := days/sum, days%sum
			if left > g {
				left = g
			}
			return seg*g+left >= half
		})
	}
	var t int
	for Fscan(in, &t); t > 0; t-- {
		Fprintln(out, solve())
	}
}

//func main() { CF1303B(os.Stdin, os.Stdout) }
