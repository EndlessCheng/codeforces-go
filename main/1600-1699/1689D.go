package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1689D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, m, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		var mx1, mx2, mn1, mn2, mnD int = -1e9, -1e9, 1e9, 1e9, 1e9
		a := make([]string, n)
		for i := range a {
			Fscan(in, &a[i])
			for j, b := range a[i] {
				if b == 'B' {
					mx1 = max(mx1, i+j)
					mn1 = min(mn1, i+j)
					mx2 = max(mx2, i-j)
					mn2 = min(mn2, i-j)
				}
			}
		}
		for i, r := range a {
			for j := range r {
				v, w := i+j, i-j
				d := max(max(abs(v-mx1), abs(v-mn1)), max(abs(w-mx2), abs(w-mn2)))
				if d < mnD {
					mnD = d
					x, y = i, j
				}
			}
		}
		Fprintln(out, x+1, y+1)
	}
}

//func main() { CF1689D(os.Stdin, os.Stdout) }
