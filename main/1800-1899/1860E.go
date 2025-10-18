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
	for x := range 26 {
		for y := range 26 {
			if g[x][y] == nil {
				continue
			}
			dis := make([]int, n)
			now++
			vis[x][y] = now
			q := slices.Clone(g[x][y])
			for step := 0; len(q) > 0; step++ {
				tmp := q
				q = nil
				for _, i := range tmp {
					dis[i] = step
					for j := i - 1; j < i+2; j += 2 {
						if 0 < j && j < n {
							a, b := s[j-1]-'a', s[j]-'a'
							if vis[a][b] < now {
								vis[a][b] = now
								q = append(q, g[a][b]...)
							}
						}
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
