package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF242C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ x, y int }
	type pd struct{ x, y, d int }
	dir8 := [...]pair{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

	var sx, sy, tx, ty, n, i, l, r int
	g := map[int]map[int]bool{}
	vis := map[int]map[int]bool{}
	for Fscan(in, &sx, &sy, &tx, &ty, &n); n > 0; n-- {
		Fscan(in, &i, &l, &r)
		if g[i] == nil {
			g[i] = map[int]bool{}
			vis[i] = map[int]bool{}
		}
		for c := l; c <= r; c++ {
			g[i][c] = true
		}
	}
	q := []pd{{sx, sy, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		if p.x == tx && p.y == ty {
			Fprint(out, p.d)
			return
		}
		for _, d := range dir8 {
			if x, y := p.x+d.x, p.y+d.y; g[x] != nil && g[x][y] && !vis[x][y] {
				vis[x][y] = true
				q = append(q, pd{x, y, p.d + 1})
			}
		}
	}
	Fprint(out, -1)
}

//func main() { CF242C(os.Stdin, os.Stdout) }
