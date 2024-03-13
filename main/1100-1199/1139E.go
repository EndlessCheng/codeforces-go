package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1139E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, d, ts int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
		b[i]--
	}
	Fscan(in, &d)
	del := make([]int, d)
	inDel := make([]bool, n)
	for i := range del {
		Fscan(in, &del[i])
		del[i]--
		inDel[del[i]] = true
	}
	g := make([][]int, m+1)
	for i, in := range inDel {
		if !in && a[i] <= m {
			g[a[i]] = append(g[a[i]], b[i])
		}
	}

	matchR := make([]int, m)
	for i := range matchR {
		matchR[i] = -1
	}
	vis := make([]int, m+1)
	var dfs func(int) bool
	dfs = func(v int) bool {
		vis[v] = ts
		for _, w := range g[v] {
			lv := matchR[w]
			if lv == -1 || vis[lv] != ts && dfs(lv) {
				matchR[w] = v
				return true
			}
		}
		return false
	}

	ans := make([]int, n+1)
	for i := d - 1; i >= 0; i-- {
		for j := ans[i+1]; j <= m; j++ {
			ts++
			if !dfs(j) {
				ans[i] = j
				break
			}
		}
		v := a[del[i]]
		if v <= m {
			g[v] = append(g[v], b[del[i]])
		}
	}
	for _, v := range ans[:d] {
		Fprintln(out, v)
	}
}

//func main() { cf1139E(os.Stdin, os.Stdout) }
