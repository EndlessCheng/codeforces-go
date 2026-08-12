package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1149E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	g := make([][]int, n+1)
	for range m {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
	}

	cnt := make([]int, n+1)
	vs := make([]int, n+1)
	var dfs func(int)
	dfs = func(v int) {
		if cnt[v] != 0 {
			return
		}
		for _, w := range g[v] {
			dfs(w)
		}
		for _, w := range g[v] {
			vs[cnt[w]] = v
		}
		cnt[v] = 1
		for vs[cnt[v]] == v {
			cnt[v]++
		}
	}

	for i := 1; i <= n; i++ {
		if cnt[i] == 0 {
			dfs(i)
		}
	}

	sg := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sg[cnt[i]] ^= a[i]
	}

	for i := n; i > 0; i-- {
		if sg[i] != 0 {
			Fprintln(out, "WIN")
			u := 0
			for j := 1; j <= n; j++ {
				if cnt[j] == i && (a[j]^sg[i]) < a[j] {
					u = j
					a[j] ^= sg[i]
					sg[i] = 0
					break
				}
			}
			for _, v := range g[u] {
				if sg[cnt[v]] != 0 {
					a[v] ^= sg[cnt[v]]
					sg[cnt[v]] = 0
				}
			}
			for j := 1; j <= n; j++ {
				Fprint(out, a[j], " ")
			}
			return
		}
	}
	Fprintln(out, "LOSE")
}

//func main() { cf1149E(bufio.NewReader(os.Stdin), os.Stdout) }
