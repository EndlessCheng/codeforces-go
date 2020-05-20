package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	unique := func(a []int) []int {
		n := len(a)
		sort.Ints(a)
		j := 0
		for i := 1; i < n; i++ {
			if a[j] != a[i] {
				j++
				a[j] = a[i]
			}
		}
		return a[:j+1]
	}
	discrete := func(a []int) map[int]int {
		ids := make(map[int]int, len(a))
		for i, v := range a {
			ids[v] = i
		}
		return ids
	}
	type line struct{ a, b, c int }
	type pair struct{ x, y int }

	var n, m, a, b, c, ans int
	Fscan(in, &n, &m)
	xs := make([]int, 0, 3+2*n+m)
	ys := make([]int, 0, 3+n+2*m)
	xs = append(xs, -2e9, 0, 2e9)
	ys = append(ys, -2e9, 0, 2e9)
	lr := make([]line, n)
	for i := range lr {
		Fscan(in, &a, &b, &c)
		lr[i] = line{a, b, c}
		xs = append(xs, a, b)
		ys = append(ys, c)
	}
	du := make([]line, m)
	for i := range du {
		Fscan(in, &a, &b, &c)
		du[i] = line{a, b, c}
		xs = append(xs, a)
		ys = append(ys, b, c)
	}
	xs = unique(xs)
	xi := discrete(xs)
	ys = unique(ys)
	yi := discrete(ys)

	lx, ly := len(xi), len(yi)
	glr := make([][]int, lx)
	gdu := make([][]int, lx)
	vis := make([][]bool, lx)
	for i := range glr {
		glr[i] = make([]int, ly)
		gdu[i] = make([]int, ly)
		vis[i] = make([]bool, ly)
	}
	for _, p := range lr {
		glr[xi[p.a]][yi[p.c]]++
		glr[xi[p.b]][yi[p.c]]--
	}
	for _, p := range du {
		gdu[xi[p.a]][yi[p.b]]++
		gdu[xi[p.a]][yi[p.c]]--
	}
	for i := 1; i < lx-1; i++ {
		for j := 1; j < ly-1; j++ {
			glr[i][j] += glr[i-1][j]
			gdu[i][j] += gdu[i][j-1]
		}
	}

	q := []pair{{xi[0], yi[0]}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y := p.x, p.y
		if x == 0 || x == lx-1 || y == 0 || y == ly-1 {
			Fprint(_w, "INF")
			return
		}
		if !vis[x][y] {
			vis[x][y] = true
			ans += (xs[x+1] - xs[x]) * (ys[y+1] - ys[y])
			if glr[x][y] == 0 {
				q = append(q, pair{x, y - 1})
			}
			if glr[x][y+1] == 0 {
				q = append(q, pair{x, y + 1})
			}
			if gdu[x][y] == 0 {
				q = append(q, pair{x - 1, y})
			}
			if gdu[x+1][y] == 0 {
				q = append(q, pair{x + 1, y})
			}
		}
	}
	Fprint(_w, ans)
}

func main() { run(os.Stdin, os.Stdout) }
