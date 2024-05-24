package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func cf1301F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int16 }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n, m, k int16
	Fscan(in, &n, &m, &k)
	a := make([][]int8, n)
	g := make([][]pair, k)
	for i := range a {
		a[i] = make([]int8, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			a[i][j]--
			g[a[i][j]] = append(g[a[i][j]], pair{int16(i), int16(j)})
		}
	}

	dis := make([][][]int8, k)
	for st := range dis {
		d := make([][]int8, n)
		for i, row := range a {
			d[i] = make([]int8, m)
			for j, x := range row {
				if x != int8(st) {
					d[i][j] = -1
				}
			}
		}
		visColor := make([]bool, k)
		visColor[st] = true
		q := slices.Clone(g[st])
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			curD := d[p.x][p.y]
			c := a[p.x][p.y]
			if !visColor[c] {
				visColor[c] = true
				for _, p2 := range g[c] {
					if d[p2.x][p2.y] < 0 {
						d[p2.x][p2.y] = curD + 1
						q = append(q, p2)
					}
				}
			}
			for _, dir := range dir4 {
				x, y := p.x+dir.x, p.y+dir.y
				if 0 <= x && x < n && 0 <= y && y < m && d[x][y] < 0 {
					d[x][y] = curD + 1
					q = append(q, pair{x, y})
				}
			}
		}
		dis[st] = d
	}

	var q, r1, c1, r2, c2 int
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &r1, &c1, &r2, &c2)
		ans := abs(r1-r2) + abs(c1-c2)
		for _, d := range dis {
			ans = min(ans, int(d[r1-1][c1-1])+1+int(d[r2-1][c2-1]))
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1301F(bufio.NewReader(os.Stdin), os.Stdout) }
