package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1860E(in io.Reader, _w io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var s string
	var m, now int
	Fscan(in, &s, &m)
	g := [26][26][]int{}
	n := len(s)
	for i := 1; i < n; i++ {
		v, w := s[i-1]-'a', s[i]-'a'
		g[v][w] = append(g[v][w], i)
	}

	qs := make([]struct{ x, y int }, m)
	ans := make([]int, m)
	for i := range qs {
		Fscan(in, &qs[i].x, &qs[i].y)
		ans[i] = abs(qs[i].x - qs[i].y)
	}

	vis := [26][26]int{}
	dis := make([]int, n)
	for x := range 26 {
		for y := range 26 {
			if g[x][y] == nil {
				continue
			}
			for i := range dis {
				dis[i] = 1e9
			}
			now++
			vis[x][y] = now
			q := slices.Clone(g[x][y])
			for _, i := range q {
				dis[i] = 0
			}
			push := func(v, w int) {
				if dis[w] == 1e9 {
					dis[w] = dis[v] + 1
					q = append(q, w)
				}
			}

			for len(q) > 0 {
				i := q[0]
				q = q[1:]
				if i > 1 {
					push(i, i-1)
				}
				if i < n-1 {
					push(i, i+1)
				}
				a, b := s[i-1]-'a', s[i]-'a'
				if vis[a][b] < now {
					vis[a][b] = now
					for _, j := range g[a][b] {
						push(i, j)
					}
				}
			}
			for i, q := range qs {
				ans[i] = min(ans[i], dis[q.x]+dis[q.y]+1)
			}
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf1860E(bufio.NewReader(os.Stdin), os.Stdout) }
