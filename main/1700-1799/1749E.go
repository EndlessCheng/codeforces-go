package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1749E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dir := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dirR := []pair{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}

	var T, n, m int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]byte, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ok := func(i, j int) bool {
			for _, d := range dir {
				x, y := i+d.x, j+d.y
				if 0 <= x && x < n && 0 <= y && y < m && a[x][y] == '#' {
					return false
				}
			}
			return true
		}

		dis := make([][]int, n)
		from := make([][]pair, n)
		for i := range dis {
			dis[i] = make([]int, m)
			from[i] = make([]pair, m)
			for j := range dis[i] {
				dis[i][j] = 1e9
				from[i][j].x = -1
			}
		}
		ql, qr := []pair{}, []pair{}
		for i, row := range a {
			if row[0] == '#' {
				dis[i][0] = 0
				ql = append(ql, pair{i, 0})
			} else if ok(i, 0) {
				dis[i][0] = 1
				qr = append(qr, pair{i, 0})
			}
		}

		for len(ql) > 0 || len(qr) > 0 {
			var p pair
			if len(ql) > 0 {
				ql, p = ql[:len(ql)-1], ql[len(ql)-1]
			} else {
				p, qr = qr[0], qr[1:]
			}
			if p.y == m-1 {
				x, y := p.x, p.y
				for x >= 0 {
					a[x][y] = '#'
					q := from[x][y]
					x, y = q.x, q.y
				}
				Fprintln(out, "YES")
				for _, row := range a {
					Fprintf(out, "%s\n", row)
				}
				continue o
			}
			for _, d := range dirR {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < m && ok(x, y) {
					wt := int(a[x][y]&1 ^ 1) // '.' = 1
					newD := dis[p.x][p.y] + wt
					if newD < dis[x][y] {
						dis[x][y] = newD
						from[x][y] = p
						if wt == 0 {
							ql = append(ql, pair{x, y})
						} else {
							qr = append(qr, pair{x, y})
						}
					}
				}
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1749E(os.Stdin, os.Stdout) }
