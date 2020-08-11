package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 另一种写法是从 1 走到 l[0], 从 l[0] 走到 l[1]...
// 统计经过的点，若长度超过 2*n-1 则不合法
// 计算每个节点的父亲，以及 l 中相邻两叶子的 LCA，利用这些数据来计算路径
// 采用 Tarjan 的离线写法可以做到 O(n) 的优秀复杂度

// github.com/EndlessCheng/codeforces-go
func CF29D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n+1)
	g[1] = append(g[1], 0)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	l := []int{}

	ms := make([]map[int]bool, n+1)
	var f func(v, fa int) map[int]bool
	f = func(v, fa int) map[int]bool {
		ms[v] = map[int]bool{}
		if len(g[v]) == 1 {
			ms[v][v] = true
			Fscan(in, &w)
			l = append(l, w)
		}
		for _, w := range g[v] {
			if w != fa {
				for u := range f(w, v) {
					ms[v][u] = true
				}
			}
		}
		return ms[v]
	}
	f(1, 0)

	vs := []int{}
	var f2 func(int, int, []int) bool
	f2 = func(v, fa int, l []int) bool {
		if len(g[v]) == 1 {
			vs = append(vs, v)
			return l[0] == v
		}
	o:
		for len(l) > 0 {
			for _, w := range g[v] {
				if w != fa && ms[w][l[0]] {
					vs = append(vs, v)
					if !f2(w, v, l[:len(ms[w])]) {
						return false
					}
					l = l[len(ms[w]):]
					continue o
				}
			}
			break
		}
		vs = append(vs, v)
		return len(l) == 0
	}
	if f2(1, 0, l) {
		for _, v := range vs {
			Fprint(out, v, " ")
		}
		return
	}
	Fprint(out, -1)
}

//func main() { CF29D(os.Stdin, os.Stdout) }
