package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 也可以用 01BFS 做：https://codeforces.com/contest/877/submission/31659168

// github.com/EndlessCheng/codeforces-go
func CF877D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var n, m, k, sx, sy, tx, ty int
	Fscan(in, &n, &m, &k)
	g := make([]string, n)
	for i := range g {
		Fscan(in, &g[i])
	}
	Fscan(in, &sx, &sy, &tx, &ty)
	sx--
	sy--
	tx--
	ty--

	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, m)
		for j := range dis[i] {
			dis[i][j] = 1e9
		}
	}
	dis[sx][sy] = 0
	q := []pair{{sx, sy}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		if p.x == tx && p.y == ty {
			Fprint(out, dis[tx][ty])
			return
		}
		for _, d := range dir4 {
			for i, x, y := 0, p.x, p.y; i < k; i++ {
				x += d.x
				y += d.y
				if !(0 <= x && x < n && 0 <= y && y < m && g[x][y] == '.' && dis[x][y] > dis[p.x][p.y]) {
					break
				}
				if newD := dis[p.x][p.y] + 1; newD < dis[x][y] {
					dis[x][y] = newD
					q = append(q, pair{x, y})
				}
			}
		}
	}
	Fprint(out, -1)
}

//func main() { CF877D(os.Stdin, os.Stdout) }
