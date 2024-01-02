package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf1761E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]string, n)
		deg := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			deg[i] = strings.Count(a[i], "1")
		}

		vis := make([]bool, n)
		vs := []int{}
		var dfs func(int)
		dfs = func(v int) {
			vis[v] = true
			vs = append(vs, v)
			for w, b := range a[v] {
				if b == '1' && !vis[w] {
					dfs(w)
				}
			}
		}
		comps := [][]int{}
		for st, b := range vis {
			if b {
				continue
			}
			vs = []int{}
			dfs(st)
			m := len(vs)
			if m == n {
				Fprintln(out, 0)
				continue o
			}
			if m == 1 {
				Fprintln(out, 1)
				Fprintln(out, vs[0]+1)
				continue o
			}
			mn := vs[0]
			for _, v := range vs {
				if deg[v] < deg[mn] {
					mn = v
				}
			}
			if deg[mn] < m-1 {
				Fprintln(out, 1)
				Fprintln(out, mn+1)
				continue o
			}
			comps = append(comps, vs)
		}
		if len(comps) > 2 {
			Fprintln(out, 2)
			Fprintln(out, comps[0][0]+1, comps[1][0]+1)
			continue
		}
		ans := comps[0]
		if len(comps[1]) < len(ans) {
			ans = comps[1]
		}
		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprint(out, v+1, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1761E(os.Stdin, os.Stdout) }
