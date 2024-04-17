package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1857G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	pow := func(x, n int64) (res int64) {
		res = 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	var T, n, s int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		es := make([]struct{ v, w, wt int }, n-1)
		for i := range es {
			Fscan(in, &es[i].v, &es[i].w, &es[i].wt)
		}
		sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })
		fa := make([]int, n+1)
		sz := make([]int, n+1)
		for i := range fa {
			fa[i] = i
			sz[i] = 1
		}
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				fa[x] = find(fa[x])
			}
			return fa[x]
		}
		ans := int64(1)
		for _, e := range es {
			v, w := find(e.v), find(e.w)
			fa[v] = w
			ans = ans * pow(int64(s-e.wt+1), int64(sz[v])*int64(sz[w])-1) % mod
			sz[w] += sz[v]
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1857G(os.Stdin, os.Stdout) }
