package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1777E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		type nb struct{ to, wt int }
		g := make([][]nb, n)
		wt := make([]int, m+1)
		for i := 1; i <= m; i++ {
			Fscan(in, &v, &w, &wt[i])
			v--
			w--
			g[v] = append(g[v], nb{w, 0})
			g[w] = append(g[w], nb{v, wt[i]})
		}
		sort.Ints(wt)
		i := sort.Search(m+1, func(i int) bool {
			limit := wt[i]
			dfn := make([]int, n)
			dfsClock := 0
			stk := []int{}
			inStk := make([]bool, n)
			last := 0
			var f func(int) int
			f = func(v int) int {
				dfsClock++
				dfn[v] = dfsClock
				lowV := dfsClock
				stk = append(stk, v)
				inStk[v] = true
				for _, e := range g[v] {
					if e.wt > limit {
						continue
					}
					if w := e.to; dfn[w] == 0 {
						lowV = min(lowV, f(w))
					} else if inStk[w] {
						lowV = min(lowV, dfn[w])
					}
				}
				if dfn[v] == lowV {
					for {
						last, stk = stk[len(stk)-1], stk[:len(stk)-1]
						inStk[last] = false
						if last == v {
							break
						}
					}
				}
				return lowV
			}
			for v, t := range dfn {
				if t == 0 {
					f(v)
				}
			}
			vis := make([]bool, n)
			left := n
			var f2 func(int)
			f2 = func(v int) {
				vis[v] = true
				left--
				for _, e := range g[v] {
					if e.wt <= limit && !vis[e.to] {
						f2(e.to)
					}
				}
			}
			f2(last)
			return left == 0
		})
		if i > m {
			Fprintln(out, -1)
		} else {
			Fprintln(out, wt[i])
		}
	}
}

//func main() { CF1777E(os.Stdin, os.Stdout) }
