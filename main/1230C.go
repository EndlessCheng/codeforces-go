package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol1230C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	edges := make([][]int, n+1)
	shown := map[int]int{}
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		edges[v] = append(edges[v], w)
		edges[w] = append(edges[w], v)
		shown[v] = 1
		shown[w] = 1
	}
	if len(shown) <= 6 {
		Fprintln(out, m)
		return
	}

	ans := 0
	for v, e := range edges {
		if len(e) == 0 {
			continue
		}
		// treat v as w
		for w, e2 := range edges {
			if w == v || len(e2) == 0 {
				continue
			}
			delta := 0
			for _, u := range e2 {
				if u == v {
					delta = 1
					break
				}
			}
			mp := map[int]int{}
			for _, u := range append(e, e2...) {
				if u == v {
					u = w
				}
				mp[u] = 1
			}
			delta = len(e) + len(e2) - delta - len(mp)
			if newAns := m - delta; newAns > ans {
				ans = newAns
			}
		}
	}
	Fprintln(out, ans)
}

//func main() {
//	Sol1230C(os.Stdin, os.Stdout)
//}
