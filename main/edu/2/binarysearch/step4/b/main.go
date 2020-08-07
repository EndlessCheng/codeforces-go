package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	const eps = 1e-8 // 由于任何 ±1 带来的均值变动至少是 1/n，eps 取 1e-8 绰绰有余
	in := bufio.NewReader(_r)
	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	type to struct{ to, wt int }
	g := make([][]to, n)
	deg := make([]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], to{w, wt})
		deg[w]++
	}
	dd := make([]int, n)

	fa := make([]int, n)
	ff := make([]int, n)
	dp := make([]float64, n)
	l, r := 0.0, 101.0 // r 稍微取大点，保证 copy(ff, fa) 的逻辑能触发！
	for t := int(math.Log2((r - l) / eps)); t > 0; t-- {
		x := (l + r) / 2
		for i := range fa {
			fa[i] = -1
		}
		for i := range dp {
			dp[i] = 1e9
		}
		dp[0] = 0
		copy(dd, deg)
		q := []int{}
		for i, d := range dd {
			if d == 0 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			v, q = q[0], q[1:]
			for _, e := range g[v] {
				w, wt := e.to, float64(e.wt)-x
				if dp[v]+wt < dp[w] {
					dp[w] = dp[v] + wt
					fa[w] = v
				}
				if dd[w]--; dd[w] == 0 {
					q = append(q, w)
				}
			}
		}
		if dp[n-1] > 0 {
			l = x
		} else {
			r = x
			copy(ff, fa)
		}
	}
	path := []int{}
	for v := n - 1; v != -1; v = ff[v] {
		path = append(path, v+1)
	}
	Fprintln(out, len(path)-1)
	for i := len(path) - 1; i >= 0; i-- {
		Fprint(out, path[i], " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
