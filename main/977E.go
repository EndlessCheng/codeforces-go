package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF977E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, ans int
	Fscan(in, &n, &m)
	deg := make([]int, n)
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		deg[v]++
		deg[w]++
	}

	vis := make([]bool, n)
	var comp []int
	var f func(int)
	f = func(v int) {
		vis[v] = true
		comp = append(comp, v)
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !vis[i] {
			comp = []int{}
			f(i)
			ans++
			for _, v := range comp {
				if deg[v] != 2 {
					ans--
					break
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() {
//	CF977E(os.Stdin, os.Stdout)
//}
