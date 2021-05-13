package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF453C(_r io.Reader, _w io.Writer) {
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
	odd := make([]bool, n)
	for i := range odd {
		Fscan(in, &odd[i])
	}

	path := []int{}
	vis := make([]bool, n)
	var f func(int, int)
	f = func(v, fa int) {
		vis[v] = true
		path = append(path, v)
		odd[v] = !odd[v]
		var son int
		for _, w := range g[v] {
			if !vis[w] {
				f(w, v)
				path = append(path, v)
				odd[v] = !odd[v]
				son = w
			}
		}
		if odd[v] {
			if fa < 0 { // 必定有 son
				path = append(path, son, v, son)
			} else {
				path = append(path, fa, v)
				odd[fa] = !odd[fa]
			}
			odd[v] = false
		}
	}
	for i, b := range odd {
		if b {
			f(i, -1)
			break
		}
	}

	for i, b := range odd {
		if b && !vis[i] { // 还有没访问的 odd
			Fprint(out, -1)
			return
		}
	}

	Fprintln(out, len(path))
	for _, v := range path {
		Fprint(out, v+1, " ")
	}
}

//func main() { CF453C(os.Stdin, os.Stdout) }
