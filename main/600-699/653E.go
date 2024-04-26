package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf653E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, ts int
	Fscan(in, &n, &m, &k)
	ban := make([][]int, n)
	for ; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		ban[v] = append(ban[v], w)
		ban[w] = append(ban[w], v)
	}
	// 没法从 0 连出去 k 条边
	if len(ban[0])+k >= n {
		Fprint(out, "impossible")
		return
	}

	// 度数最大的点
	maxV := 1
	for v := 1; v < n; v++ {
		if len(ban[v]) < len(ban[maxV]) {
			maxV = v
		}
	}

	ccNum := n - 1
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		f, t := find(from), find(to)
		if f == t {
			return
		}
		ccNum--
		fa[f] = t
	}
	same := func(x, y int) bool { return find(x) == find(y) }

	has := make([]int, n)
	mergeInv := func(v int) {
		ts++
		has[v] = ts
		for _, w := range ban[v] {
			has[w] = ts
		}
		for i := 1; i < n; i++ {
			if has[i] != ts {
				merge(i, v)
			}
		}
	}
	mergeInv(maxV)
	for v := 1; v < n; v++ {
		if !same(v, maxV) {
			mergeInv(v)
		}
	}
	// 要从 0 连出去的边超过 k
	if ccNum > k {
		Fprint(out, "impossible")
		return
	}

	ts++
	for _, w := range ban[0] {
		has[w] = ts
	}
	vis := make([]bool, n)
	for i := 1; i < n; i++ {
		if has[i] != ts && !vis[find(i)] {
			vis[find(i)] = true
		}
	}
	for i := 1; i < n; i++ {
		if !vis[find(i)] {
			// 还有独立的连通块
			Fprint(out, "impossible")
			return
		}
	}
	Fprint(out, "possible")
}

//func main() { cf653E(os.Stdin, os.Stdout) }
