package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1365D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dir4 := [...]pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var t, n, m int
	G := []byte{'G'}
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		cg := 0
		g := make([][]byte, n)
		for i := range g {
			Fscan(in, &g[i])
			cg += bytes.Count(g[i], G)
		}
		for i, row := range g {
			for j, c := range row {
				if c == 'B' {
					for _, d := range dir4 {
						if x, y := i+d.x, j+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] != 'B' {
							g[x][y] = '#'
						}
					}
				}
			}
		}
		cnt := 0
		if g[n-1][m-1] != '#' {
			g[n-1][m-1] = '#'
			for q := []pair{{n - 1, m - 1}}; len(q) > 0; {
				p := q[0]
				q = q[1:]
				for _, d := range dir4 {
					if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] != '#' {
						if g[x][y] == 'G' {
							cnt++
						}
						g[x][y] = '#'
						q = append(q, pair{x, y})
					}
				}
			}
		}
		if cnt == cg {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { CF1365D(os.Stdin, os.Stdout) }
