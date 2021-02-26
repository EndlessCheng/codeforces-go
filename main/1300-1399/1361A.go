package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1361A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	vs := make([][]int, n)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		vs[v-1] = append(vs[v-1], i)
	}
	mp := make([]map[int]bool, n)
	for i := range mp {
		mp[i] = map[int]bool{}
	}
	mex := make([]int, n)
	ans := make([]int, 0, n)
	for i, vs := range vs {
		for _, v := range vs {
			v--
			for mp[v][mex[v]] {
				mex[v]++
			}
			if mex[v] != i {
				Fprint(out, -1)
				return
			}
			for _, w := range g[v] {
				mp[w][i] = true
			}
		}
		ans = append(ans, vs...)
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF1361A(os.Stdin, os.Stdout) }
