package main

import (
	"bufio"
	. "fmt"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF1305D() {
	in := bufio.NewReader(os.Stdin)
	var n int
	Fscan(in, &n)
	edges := map[[2]int]bool{}
	for i := 0; i < n-1; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		edges[[2]int{v, w}] = true
	}

	var lca int
	for {
		g := make([][]int, n)
		for e := range edges {
			v, w := e[0], e[1]
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		if len(g[lca]) == 0 {
			break
		}

		// 求直径
		var u, maxD int
		var f func(v, fa, d int)
		f = func(v, fa, d int) {
			if d > maxD {
				maxD = d
				u = v
			}
			for _, w := range g[v] {
				if w != fa {
					f(w, v, d+1)
				}
			}
		}
		maxD = -1
		f(lca, -1, 0)
		dv := u
		maxD = -1
		f(dv, -1, 0)
		dw := u
		Println("?", dv+1, dw+1)
		Fscan(in, &lca)
		lca--
		if lca == dv || lca == dw {
			break
		}

		// 删除 LCA 到直径上的边
		var del func(v, end, fa int) bool
		del = func(v, end, fa int) bool {
			if v == end {
				delete(edges, [2]int{v, fa})
				delete(edges, [2]int{fa, v})
				return true
			}
			for _, w := range g[v] {
				if w != fa {
					if del(w, end, v) {
						delete(edges, [2]int{v, w})
						delete(edges, [2]int{w, v})
						return true
					}
				}
			}
			return false
		}
		del(lca, dv, -1)
		del(lca, dw, -1)
	}
	Println("!", lca+1)
}

//func main() { CF1305D() }
