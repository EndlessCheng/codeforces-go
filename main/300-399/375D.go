package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf375D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, k, dfn int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	type pair struct{ i, k int }
	qs := make([][]pair, n)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &k)
		if k <= n {
			qs[v-1] = append(qs[v-1], pair{i, k})
		}
	}

	nodes := make([]struct{ l, r, hson int }, n) // [l,r)
	nodeVals := make([]int, 0, n)
	var build func(int, int) int
	build = func(v, fa int) int {
		nodes[v].l = dfn
		dfn++
		nodeVals = append(nodeVals, a[v])
		size, hsz, hson := 1, 0, -1
		for _, w := range g[v] {
			if w != fa {
				sz := build(w, v)
				size += sz
				if sz > hsz {
					hsz, hson = sz, w
				}
			}
		}
		nodes[v].r = nodes[v].l + size
		nodes[v].hson = hson
		return size
	}
	build(0, -1)

	ans := make([]int, m)
	cnt := [1e5 + 1]int{}
	cc := make([]int, n+1)
	var f func(int, int)
	f = func(v, fa int) {
		hson := nodes[v].hson
		for _, w := range g[v] {
			if w == fa || w == hson {
				continue
			}
			f(w, v)
			// 恢复现场，这样下一棵子树不会受到影响
			for _, x := range nodeVals[nodes[w].l:nodes[w].r] {
				cc[cnt[x]]--
				cnt[x]--
			}
		}
		if hson >= 0 {
			f(hson, v)
			// 此时重儿子的数据已经添加
		}

		// 添加根节点的数据
		cnt[a[v]]++
		cc[cnt[a[v]]]++
		// 添加非重儿子的数据
		for _, w := range g[v] {
			if w == fa || w == hson {
				continue
			}
			for _, x := range nodeVals[nodes[w].l:nodes[w].r] {
				cnt[x]++
				cc[cnt[x]]++
			}
		}

		// 子树 v 的所有数据添加完毕，回答询问
		for _, q := range qs[v] {
			ans[q.i] = cc[q.k]
		}
	}
	f(0, -1)

	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf375D(os.Stdin, os.Stdout) }
