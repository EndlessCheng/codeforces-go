package main

import (
	"bufio"
	. "fmt"
	"io"
)

func cf1899G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, l, r, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		nodes := make([]struct{ l, r int }, n)
		dfn := 1
		var dfs func(int, int)
		dfs = func(v, fa int) {
			nodes[v].l = dfn
			dfn++
			for _, w := range g[v] {
				if w != fa {
					dfs(w, v)
				}
			}
			nodes[v].r = dfn - 1
		}
		dfs(0, -1)

		type data struct{ x, d, qid int }
		qs := make([][]data, n)
		for i := 0; i < q; i++ {
			Fscan(in, &l, &r, &x)
			l -= 2
			r--
			x--
			if l >= 0 {
				qs[l] = append(qs[l], data{x, -1, i})
			}
			qs[r] = append(qs[r], data{x, 1, i})
		}

		ans := make([]int, q)
		t := make([]int, n+1)
		pre := func(i int) (res int) {
			for ; i > 0; i &= i - 1 {
				res += t[i]
			}
			return
		}
		for i, ps := range qs {
			for j := nodes[a[i]-1].l; j <= n; j += j & -j {
				t[j]++
			}
			for _, p := range ps {
				node := nodes[p.x]
				ans[p.qid] += p.d * (pre(node.r) - pre(node.l-1))
			}
		}
		for _, v := range ans {
			if v > 0 {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1899G(bufio.NewReader(os.Stdin), os.Stdout) }
