package main

import (
	"bytes"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1344C(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	deg := make([]int, n)
	rg := make([][]int, n)
	for range m {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		deg[w]++
		rg[w] = append(rg[w], v)
	}

	q := make([]int, 0, n)
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for i := 0; i < len(q); i++ {
		for _, w := range g[q[i]] {
			deg[w]--
			if deg[w] == 0 {
				q = append(q, w)
			}
		}
	}
	if len(q) < n {
		Fprint(out, -1)
		return
	}

	nxt := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		v := q[i]
		nxt[v] = v
		for _, w := range g[v] {
			nxt[v] = min(nxt[v], nxt[w])
		}
	}

	ans := bytes.Repeat([]byte{'E'}, n)
	na := 0
	pre := make([]int, n)
	for _, v := range q {
		pre[v] = v
		for _, w := range rg[v] {
			pre[v] = min(pre[v], pre[w])
		}
		if pre[v] == v && nxt[v] == v {
			ans[v] = 'A'
			na++
		}
	}

	Fprintln(out, na)
	Fprintf(out, "%s", ans)
}

//func main() { cf1344C(bufio.NewReader(os.Stdin), os.Stdout) }
