package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1350E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dir4 := [...]pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var n, m, t, x, y int
	var p int64
	Fscan(in, &n, &m, &t)
	g := make([][]byte, n)
	dep := make([][]int, n)
	for i := range g {
		Fscan(in, &g[i])
		dep[i] = make([]int, m)
	}
	q := []pair{}
	for i, row := range g {
		for j, b := range row {
			for _, d := range dir4 {
				if x, y := i+d.x, j+d.y; x >= 0 && x < n && y >= 0 && y < m && g[x][y] == b {
					dep[i][j] = 1
					q = append(q, pair{i, j})
					break
				}
			}
		}
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, d := range dir4 {
			if x, y := p.x+d.x, p.y+d.y; x >= 0 && x < n && y >= 0 && y < m && dep[x][y] == 0 {
				dep[x][y] = dep[p.x][p.y] + 1
				q = append(q, pair{x, y})
			}
		}
	}
	for ; t > 0; t-- {
		Fscan(in, &x, &y, &p)
		c := g[x-1][y-1] - '0'
		d := dep[x-1][y-1]
		if d == 0 {
			Fprintln(out, c)
			continue
		}
		p -= int64(d - 1)
		if p <= 0 {
			Fprintln(out, c)
		} else {
			Fprintln(out, byte(p&1)^c)
		}
	}
}

//func main() { CF1350E(os.Stdin, os.Stdout) }
