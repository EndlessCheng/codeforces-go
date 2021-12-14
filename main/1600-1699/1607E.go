package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1607E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	dir4 := []struct{ x, y int }{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}}

	var T, n, m int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &s)
		var x, y, l, r, d, u int
		ansX, ansY := 1, 1
		for _, b := range s {
			x += dir4[b].x
			y += dir4[b].y
			if x < l {
				l = x
			} else if x > r {
				r = x
			}
			if y < d {
				d = y
			} else if y > u {
				u = y
			}
			if r-l >= m || u-d >= n {
				break
			}
			ansX, ansY = 1+u, 1-l
		}
		Fprintln(out, ansX, ansY)
	}
}

//func main() { CF1607E(os.Stdin, os.Stdout) }
