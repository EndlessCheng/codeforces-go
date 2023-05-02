package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1707C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w int
	Fscan(in, &n, &m)
	mst := make([][]int, n)
	type edge struct{ x, y int }
	nonMstEdges := []edge{}
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
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		fv, fw := find(v), find(w)
		if fv != fw {
			mst[v] = append(mst[v], w)
			mst[w] = append(mst[w], v)
			fa[fv] = fw
		} else {
			nonMstEdges = append(nonMstEdges, edge{v, w})
		}
	}

	dep := make([]int, n)
	var build func(int, int, int)
	build = func(v, fa, d int) {
		dep[v] = d
		for _, w := range mst[v] {
			if w != fa {
				build(w, v, d+1)
			}
		}
	}
	build(0, -1, 0)

	check := make([][]int, n)
	for _, e := range nonMstEdges {
		v, w := e.x, e.y
		if dep[v] < dep[w] { // 保证 w 的深度不超过 v
			v, w = w, v
		}
		check[v] = append(check[v], w)
	}

	diff := make([]int, n)
	stk := make([]int, 0, n)
	inStk := make([]bool, n)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		stk = append(stk, v)
		inStk[v] = true
		for _, w := range check[v] {
			diff[v]++ // 子树 v 都是合法的
			if inStk[w] { // w 是 v 的祖先
				diff[0]++
				diff[stk[dep[w]+1]]-- // 从 w 到 v 的路径上的中间节点（不含 v 和 w）都是不合法的
			} else { // 横向边
				diff[w]++ // 子树 w 都是合法的
			}
		}
		for _, w := range mst[v] {
			if w != fa {
				dfs(w, v)
			}
		}
		stk = stk[:len(stk)-1]
		inStk[v] = false
	}
	dfs(0, -1)

	ans := bytes.Repeat([]byte{'0'}, n)
	build = func(v, fa, s int) {
		s += diff[v]
		if s == m-n+1 {
			ans[v] = '1'
		}
		for _, w := range mst[v] {
			if w != fa {
				build(w, v, s)
			}
		}
	}
	build(0, -1, 0)
	Fprintf(out, "%s", ans)
}

//func main() { CF1707C(os.Stdin, os.Stdout) }
