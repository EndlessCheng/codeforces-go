package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1891F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, q, op, v, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &q)
		g := make([][]int, q+1)
		addT := make([]int, q+1)
		type pair struct{ t, x int }
		rec := make([][]pair, q+1)
		n := 1
		for t := range q {
			Fscan(in, &op, &v)
			v--
			if op == 1 {
				g[v] = append(g[v], n)
				addT[n] = t
				n++
			} else {
				Fscan(in, &x)
				rec[v] = append(rec[v], pair{t, x})
			}
		}

		tree := make([]int, q+1)
		add := func(i, x int) {
			for ; i <= q; i += i & -i {
				tree[i] += x
			}
		}
		pre := func(i int) (res int) {
			for ; i > 0; i &= i - 1 {
				res += tree[i]
			}
			return
		}

		ans := make([]any, n)
		var dfs func(int)
		dfs = func(v int) {
			for _, p := range rec[v] {
				add(q-p.t, p.x)
			}
			ans[v] = pre(q - addT[v])
			for _, w := range g[v] {
				dfs(w)
			}
			for _, p := range rec[v] {
				add(q-p.t, -p.x)
			}
		}
		dfs(0)
		Fprintln(out, ans...)
	}
}

//func main() { cf1891F(bufio.NewReader(os.Stdin), os.Stdout) }
