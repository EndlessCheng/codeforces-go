package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1401D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	type nb struct{ to, i int }

	var T, n, v, w, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]nb, n+1)
		for i := 0; i < n-1; i++ {
			Fscan(in, &v, &w)
			g[v] = append(g[v], nb{w, i})
			g[w] = append(g[w], nb{v, i})
		}
		Fscan(in, &m)
		p := make([]int, m)
		for i := range p {
			Fscan(in, &p[i])
		}
		for len(p) < n-1 {
			p = append(p, 1)
		}
		sort.Ints(p)
		pp := int64(p[n-2])
		for _, v := range p[n-1:] {
			pp = pp * int64(v) % mod
		}
		p[n-2] = int(pp)
		cnt := make([]int64, n-1)
		var f func(v, fa int) int
		f = func(v, fa int) int {
			sz := 1
			for _, p := range g[v] {
				if p.to != fa {
					s := f(p.to, v)
					cnt[p.i] = int64(s) * int64(n-s)
					sz += s
				}
			}
			return sz
		}
		f(1, 0)
		sort.Slice(cnt, func(i, j int) bool { return cnt[i] < cnt[j] })
		ans := int64(0)
		for i, c := range cnt {
			ans += c * int64(p[i]) % mod
		}
		Fprintln(out, ans%mod)
	}
}

//func main() { CF1401D(os.Stdin, os.Stdout) }
