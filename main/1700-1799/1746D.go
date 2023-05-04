package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1746D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		g := make([][]int, n)
		for w := 1; w < n; w++ {
			Fscan(in, &p)
			g[p-1] = append(g[p-1], w)
		}
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		type pair struct{ x, y int }
		dp := map[pair]int64{}
		var f func(int, int) int64
		f = func(v, k int) (res int64) {
			if k == 0 {
				return
			}
			res = int64(a[v]) * int64(k)
			if g[v] == nil {
				return
			}
			p := pair{v, k}
			if v, ok := dp[p]; ok {
				return v
			}
			if m := len(g[v]); k%m == 0 {
				for _, w := range g[v] {
					res += f(w, k/m)
				}
			} else {
				delta := make([]int64, m)
				for i, w := range g[v] {
					r := f(w, k/m)
					res += r
					delta[i] = f(w, k/m+1) - r
				}
				sort.Slice(delta, func(i, j int) bool { return delta[i] > delta[j] })
				for _, v := range delta[:k%m] {
					res += v
				}
			}
			dp[p] = res
			return
		}
		Fprintln(out, f(0, k))
	}
}

//func main() { CF1746D(os.Stdin, os.Stdout) }
