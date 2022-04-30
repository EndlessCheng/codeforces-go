package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1613C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var h int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &h)
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		l, r := int64(1), h
	o:
		for l < r {
			m := (l + r) / 2
			left := h
			for i, t := range a {
				if i < n-1 && a[i+1]-t < m {
					t = a[i+1] - t
				} else {
					t = m
				}
				if left -= t; left <= 0 {
					r = m
					continue o
				}
			}
			l = m + 1
		}
		Fprintln(out, l)
	}
}

//func main() { CF1613C(os.Stdin, os.Stdout) }
