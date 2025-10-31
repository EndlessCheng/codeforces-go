package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func p2895(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	const mx = 303
	g := [mx][mx]int{}
	for i := 0; i < mx; i++ {
		for j := 0; j < mx; j++ {
			g[i][j] = 1e9
		}
	}
	var n, x, y, t int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &x, &y, &t)
		g[x][y] = min(g[x][y], t)
		for _, d := range dir4 {
			if xx, yy := x+d.x, y+d.y; 0 <= xx && 0 <= yy {
				g[xx][yy] = min(g[xx][yy], t)
			}
		}
	}

	vis := [mx][mx]bool{}
	type point struct{ x, y, t int }
	q := []point{}
	if g[0][0] > 0 {
		vis[0][0] = true
		q = append(q, point{})
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		if g[p.x][p.y] == 1e9 {
			Fprint(out, p.t)
			return
		}
		for _, d := range dir4 {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && 0 <= y && !vis[x][y] && g[x][y] > p.t+1 {
				vis[x][y] = true
				q = append(q, point{x, y, p.t + 1})
			}
		}
	}
	Fprint(out, -1)
}

//func main() { p2895(os.Stdin, os.Stdout) }
