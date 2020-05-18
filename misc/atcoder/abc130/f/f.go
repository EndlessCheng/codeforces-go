package main

import (
	"bufio"
	. "fmt"
	"io"
	. "math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ x, y, d int }
	dir4 := [...][2]float64{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}}
	search := func(l, r float64, f func(float64) float64) float64 {
		for i := 0; i < 100; i++ {
			m1 := l + (r-l)/3
			m2 := r - (r-l)/3
			v1, v2 := f(m1), f(m2)
			if v1 < v2 {
				r = m2
			} else {
				l = m1
			}
		}
		return f((l + r) / 2)
	}

	var n, x, y int
	var s []byte
	Fscan(in, &n)
	ps := make([]pair, n)
	for i := range ps {
		Fscan(in, &x, &y, &s)
		ps[i] = pair{x, y, int(s[0])}
	}
	// 1e8 秒后，最坏情况下的左右会重合
	ans := search(0, 1e8, func(t float64) float64 {
		xl, xr, yl, yr := 1e9, -1e9, 1e9, -1e9
		for _, p := range ps {
			x, y := float64(p.x)+t*dir4[p.d][0], float64(p.y)+t*dir4[p.d][1]
			xl = Min(xl, x)
			xr = Max(xr, x)
			yl = Min(yl, y)
			yr = Max(yr, y)
		}
		return (xr - xl) * (yr - yl)
	})
	Fprintf(_w, "%.10f", ans)
}

func main() { run(os.Stdin, os.Stdout) }
