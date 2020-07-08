package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1335F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	dir4 := ['U' + 1]int{'L': -1, 'R': 1}
	type pair struct{ v, d int }

	var t, n, m int
	var s string
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		dir4['D'] = m
		dir4['U'] = -m
		c := make([]byte, 0, n*m)
		for i := 0; i < n; i++ {
			Fscan(in, &s)
			c = append(c, s...)
		}
		deg := make([]int8, n*m)
		rg := make([][]int, n*m)
		g := make([]byte, 0, n*m)
		for i := 0; i < n; i++ {
			Fscan(in, &s)
			g = append(g, s...)
			for j, b := range s {
				v := i*m + j
				w := v + dir4[b]
				deg[w]++
				rg[w] = append(rg[w], v)
			}
		}

		vis := make([]int8, n*m)
		q := []int{}
		for i, d := range deg {
			if d == 0 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			vis[v] = 1
			v += dir4[g[v]]
			deg[v]--
			if deg[v] == 0 {
				q = append(q, v)
			}
		}

		ans1, ans2 := 0, 0
		for v, b := range vis {
			if b != 0 {
				continue
			}
			cycle := []int{}
			for ; vis[v] == 0; v += dir4[g[v]] {
				vis[v] = 2
				cycle = append(cycle, v)
			}
			sz := len(cycle)
			for i, cv := range cycle {
				for _, r := range rg[cv] {
					if vis[r] == 2 {
						continue
					}
					// 从树枝的根部开始遍历反图
					maxDep := 0
					for q := []pair{{r, 1}}; len(q) > 0; {
						p := q[0]
						q = q[1:]
						v, d := p.v, p.d
						if d > maxDep && c[v] == '0' {
							c[cycle[((i-d)%sz+sz)%sz]] = '0' // 从 cycle[i] 上倒着走 d 步
							maxDep = d
						}
						for _, w := range rg[v] {
							q = append(q, pair{w, d + 1})
						}
					}
				}
			}
			ans1 += sz
			for _, v := range cycle {
				if c[v] == '0' {
					ans2++
				}
			}
		}
		Fprintln(out, ans1, ans2)
	}
}

//func main() { CF1335F(os.Stdin, os.Stdout) }
