package _00_199

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF191C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, q int
	Fscan(in, &n)
	type edge struct{ to, id int }
	g := make([][]edge, n)
	for i := 0; i < n-1; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], edge{w, i})
		g[w] = append(g[w], edge{v, i})
	}
	pa := make([]int, n)
	for i := range pa {
		pa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if pa[x] != x {
			pa[x] = find(pa[x])
		}
		return pa[x]
	}

	d := make([]int, n)
	qs := make([][]int, n)
	vis := make([]int8, n)
	var _f func(int)
	_f = func(v int) {
		vis[v] = 1
		for _, e := range g[v] {
			if w := e.to; vis[w] == 0 {
				_f(w)
				pa[w] = v
			}
		}
		for _, w := range qs[v] {
			if vis[w] == 2 {
				d[v]++
				d[w]++
				d[find(w)] -= 2
			}
		}
		vis[v] = 2
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &v, &w)
		if v != w {
			v--
			w--
			qs[v] = append(qs[v], w)
			qs[w] = append(qs[w], v)
		}
	}
	_f(0)

	ans := make([]interface{}, n-1)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		sum := d[v]
		for _, e := range g[v] {
			if w := e.to; w != fa {
				s := f(w, v)
				ans[e.id] = s
				sum += s
			}
		}
		return sum
	}
	f(0, -1)
	Fprint(out, ans...)
}

//func main() { CF191C(os.Stdin, os.Stdout) }
