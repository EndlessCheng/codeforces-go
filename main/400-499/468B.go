package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF468B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, y int
	Fscan(in, &n, &x, &y)
	a := make([]int, n)
	id := map[int]int{}
	for i := range a {
		Fscan(in, &a[i])
		id[a[i]] = i + 1
	}

	m := n * 2
	g := make([][]int, m)
	rg := make([][]int, m)
	add := func(v, w int) {
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
	}

	// 令「v 在集合 A」中为真，「v 在集合 B」中为假
	for i, v := range a {
		if j := id[x-v] - 1; j >= 0 {
			// i 为真则 j 为真
			add(i, j)
			add(j+n, i+n)
		} else {
			// i 为假
			add(i, i+n)
		}
		if j := id[y-v] - 1; j >= 0 {
			// i 为假则 j 为假
			add(i+n, j+n)
			add(j, i)
		} else {
			// i 为真
			add(i+n, i)
		}
	}
	// 注：若 v 无法在 A 中且无法在 B 中，即 i 即假又真，这会导致 i 和 i+n 在同一个 SCC 中，见后面的代码
	// 也可以在建图时直接判断出这种情况

	vs := make([]int, 0, m)
	vis := make([]bool, m)
	var dfs func(int)
	dfs = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				dfs(w)
			}
		}
		vs = append(vs, v)
	}
	for i, b := range vis {
		if !b {
			dfs(i)
		}
	}
	vis = make([]bool, m)
	sccIDs := make([]int, m)
	sid := 0
	var rdfs func(int)
	rdfs = func(v int) {
		sccIDs[v] = sid
		vis[v] = true
		for _, w := range rg[v] {
			if !vis[w] {
				rdfs(w)
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		if v := vs[i]; !vis[v] {
			rdfs(v)
			sid++
		}
	}

	ans := make([]interface{}, n)
	for i, id := range sccIDs[:n] {
		if id == sccIDs[i+n] {
			Fprint(out, "NO")
			return
		}
		ans[i] = 0
		if id < sccIDs[i+n] {
			ans[i] = 1
		}
	}
	Fprintln(out, "YES")
	Fprintln(out, ans...)
}

//func main() { CF468B(os.Stdin, os.Stdout) }
