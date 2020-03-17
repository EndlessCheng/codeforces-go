package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1325C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	if n == 2 { // 注意特判 n=2 的情况，此时两个叶结点对应同一条边
		Fprint(out, 0)
		return
	}
	g := make([][]int, n)
	edgeI := make([]int, n)
	for i := 0; i < n-1; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		edgeI[v] = i
		edgeI[w] = i
	}
	ans := make([]int, n-1)
	label := 1
	for v, vs := range g {
		if len(vs) == 1 {
			ans[edgeI[v]] = label
			label++
		}
	}
	for i, v := range ans {
		if v == 0 {
			ans[i] = label
			label++
		}
	}
	for _, v := range ans {
		Fprintln(out, v-1)
	}
}

//func main() { CF1325C(os.Stdin, os.Stdout) }
