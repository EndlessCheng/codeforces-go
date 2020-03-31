package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1005F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, k, v, w int
	Fscan(in, &n, &m, &k)
	edges := make([][2]int, m)
	g := make([][][2]int, n)
	for i := range edges {
		Fscan(in, &v, &w)
		v--
		w--
		edges[i] = [2]int{v, w}
		g[v] = append(g[v], [2]int{w, i})
		g[w] = append(g[w], [2]int{v, i})
	}

	fromEdges := make([][]int, n)
	dep := make([]int, n)
	for i := range dep {
		dep[i] = -1
	}
	dep[0] = 0
	q := []int{0}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		for _, e := range g[v] {
			if w := e[0]; dep[w] < 0 {
				dep[w] = dep[v] + 1
				q = append(q, w)
				fromEdges[w] = append(fromEdges[w], e[1])
			} else if dep[w] == dep[v]+1 {
				fromEdges[w] = append(fromEdges[w], e[1])
			}
		}
	}
	ways := 1
	for _, ids := range fromEdges[1:] {
		ways *= len(ids)
		if ways >= k {
			ways = k
			break
		}
	}
	Fprintln(out, ways)
	ans := make([]byte, m)
	for i := range ans {
		ans[i] = '0'
	}
	var printRoads func(p int)
	printRoads = func(p int) {
		if ways == 0 {
			return
		}
		if p == n {
			Fprintln(out, string(ans))
			ways--
			return
		}
		for _, eid := range fromEdges[p] {
			ans[eid] = '1'
			printRoads(p + 1)
			ans[eid] = '0'
		}
	}
	printRoads(1)
}

//func main() { CF1005F(os.Stdin, os.Stdout) }
