package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF723F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, s, t, ds, dt, sv, tv, both int
	Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	Fscan(in, &s, &t, &ds, &dt)

	ans := make([][2]int, 0, n-1)
	vis := make([]bool, n+1)
	var f func(int)
	f = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if w == s {
				sv = v
			} else if w == t {
				tv = v
			} else if !vis[w] {
				ans = append(ans, [2]int{v, w}) // 不与 s t 相邻的点，可以直接连边（DFS 树即生成树）
				f(w)
			}
		}
	}
	type st struct{ s, t int }
	nb := []st{}
	for i := 1; i <= n; i++ {
		if !vis[i] && i != s && i != t { // 对删去 s 和 t 的图 DFS
			sv, tv = -1, -1
			f(i)
			// 若该连通分量只与 s t 其中一个相连，可以任取一相邻点直接连边
			if sv < 0 {
				ans = append(ans, [2]int{tv, t})
				dt--
			} else if tv < 0 {
				ans = append(ans, [2]int{sv, s})
				ds--
			} else {
				both++
			}
			nb = append(nb, st{sv, tv})
		}
	}

	// 此时 s 和 t 还未连通，s 和 t 还需要连边
	if ds < 1 || dt < 1 || ds+dt <= both {
		Fprint(out, "No")
		return
	}

	if both == 0 { // s 和 t 相邻
		ans = append(ans, [2]int{s, t})
	} else {
		conn := false
		for _, p := range nb {
			if p.s >= 0 && p.t >= 0 {
				if conn {
					if ds > 0 {
						ans = append(ans, [2]int{p.s, s})
						ds--
					} else {
						ans = append(ans, [2]int{p.t, t})
						dt--
					}
				} else {
					conn = true
					ans = append(ans, [2]int{p.s, s}, [2]int{p.t, t})
					ds--
					dt--
				}
			}
		}
	}
	Fprintln(out, "Yes")
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

//func main() { CF723F(os.Stdin, os.Stdout) }
