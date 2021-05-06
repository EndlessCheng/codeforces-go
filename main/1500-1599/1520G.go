package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1520G(_r io.Reader, out io.Writer) {
	_i, buf := 1<<12, make([]byte, 1<<12)
	rc := func() byte {
		if _i == 1<<12 {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		neg := false
		for ; '0' > b; b = rc() {
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	const inf int64 = 1e18
	type pair struct{ x, y int }
	var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	n, m, w := r(), r(), int64(r())
	a := make([][]int, n)
	dis := make([][]int64, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			a[i][j] = r()
		}
		dis[i] = make([]int64, m)
	}

	ans := inf
	bfsAll := func(sx, sy int) int64 {
		for i := range dis {
			for j := range dis[i] {
				dis[i][j] = -1
			}
		}
		dis[sx][sy] = 0
		q := []pair{{sx, sy}}
		for curD := w; len(q) > 0; curD += w {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dir4 {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && a[x][y] >= 0 && dis[x][y] < 0 {
						dis[x][y] = curD
						q = append(q, pair{x, y})
					}
				}
			}
		}
		if sx == 0 && dis[n-1][m-1] > 0 {
			ans = dis[n-1][m-1]
		}
		mi := inf
		for i, r := range a {
			for j, v := range r {
				if v > 0 && dis[i][j] >= 0 {
					mi = min(mi, dis[i][j]+int64(v))
				}
			}
		}
		return mi
	}
	ans = min(ans, bfsAll(0, 0)+bfsAll(n-1, m-1))
	if ans == inf {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF1520G(os.Stdin, os.Stdout) }
