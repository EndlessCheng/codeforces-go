package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1296F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, q, min int
	Fscan(in, &n)
	g := make([][][2]int, n)
	for i := 0; i < n-1; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], [2]int{w, i})
		g[w] = append(g[w], [2]int{v, i})
	}
	Fscan(in, &q)
	qs := make([][3]int, q)
	for i := range qs {
		Fscan(in, &qs[i][0], &qs[i][1], &qs[i][2])
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i][2] > qs[j][2] })

	ans := make([]int, n-1)
	var ok bool
	var f func(v, fa int) bool
	f = func(v, fa int) bool {
		if v == w {
			return true
		}
		for _, e := range g[v] {
			if w, id := e[0], e[1]; w != fa {
				if f(w, v) {
					if ans[id] == 0 {
						ans[id] = min
						ok = true
					}
					return true
				}
			}
		}
		return false
	}
	for i, q := range qs {
		v, w, min = q[0]-1, q[1]-1, q[2]
		ok = false
		f(v, -1)
		if !ok && min != qs[i-1][2] {
			Fprint(out, -1)
			return
		}
	}
	for _, v := range ans {
		if v == 0 {
			v = 1
		}
		Fprint(out, v, " ")
	}
}

func main() { CF1296F(os.Stdin, os.Stdout) }
