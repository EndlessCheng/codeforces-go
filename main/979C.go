package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF979C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int64
	var flowrisa, beetopia, v, w int
	Fscan(in, &n, &flowrisa, &beetopia)
	g := make([][]int, n+1)
	for i := 1; i < int(n); i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	size := make([]int, n+1)
	fa := make([]int, n+1)
	var f func(int, int) int
	f = func(v, p int) int {
		fa[v] = p
		sz := 1
		for _, w := range g[v] {
			if w != p {
				sz += f(w, v)
			}
		}
		size[v] = sz
		return sz
	}
	f(flowrisa, 0)
	v = beetopia
	for ; fa[v] != flowrisa; v = fa[v] {
	}
	Fprint(out, n*(n-1)-(n-int64(size[v]))*int64(size[beetopia]))
}

//func main() {
//	CF979C(os.Stdin, os.Stdout)
//}
