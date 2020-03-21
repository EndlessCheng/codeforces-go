package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1307D(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n, m, k := read(), read(), read()
	sp := make([]int, k)
	for i := range sp {
		sp[i] = read() - 1
	}
	g := make([][]int, n)
	for ; m > 0; m-- {
		v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	bfs := func(st int) []int {
		dep := make([]int, n)
		for i := range dep {
			dep[i] = -1
		}
		dep[st] = 0
		q := []int{st}
		for len(q) > 0 {
			var v int
			v, q = q[0], q[1:]
			for _, w := range g[v] {
				if dep[w] == -1 {
					dep[w] = dep[v] + 1
					q = append(q, w)
				}
			}
		}
		return dep
	}
	depSt, depEnd := bfs(0), bfs(n-1)
	sort.Slice(sp, func(i, j int) bool { a, b := sp[i], sp[j]; return depSt[a]-depEnd[a] < depSt[b]-depEnd[b] })
	ans := 0
	maxDepSt := int(-1e9)
	for _, v := range sp {
		ans = max(ans, maxDepSt+depEnd[v]+1)
		maxDepSt = max(maxDepSt, depSt[v])
	}
	if depSt[n-1] < ans {
		ans = depSt[n-1]
	}
	Fprint(_w, ans)
}

//func main() { CF1307D(os.Stdin, os.Stdout) }
