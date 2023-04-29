package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1806E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, v, w int
	Fscan(in, &n, &q)
	a := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	pa := make([]int, n+1)
	for i := 2; i <= n; i++ {
		Fscan(in, &pa[i])
	}
	sum := make([]int64, n+1)
	dep := make([]int, n+1)
	cnt := make([]int, n+1)
	idx := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[pa[i]] + a[i]*a[i]
		dep[i] = dep[pa[i]] + 1
		idx[i] = cnt[dep[i]]
		cnt[dep[i]]++
	}
	const B = 317
	dp := make([][B]int64, n+1)
	var f func(int, int) int64
	f = func(v, w int) (res int64) {
		if v == w {
			return sum[v]
		}
		if cnt[dep[v]] <= B { // 节点数大于 √n 的层至多有 √n 个，暴力算都没事
			if v > w {
				v, w = w, v
			}
			dv := &dp[v][idx[w]] // 避免使用哈希表的技巧：把 w 映射为 idx[w]
			if *dv > 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		return f(pa[v], pa[w]) + a[v]*a[w]
	}
	for ; q > 0; q-- {
		Fscan(in, &v, &w)
		Fprintln(out, f(v, w))
	}
}

//func main() { CF1806E(os.Stdin, os.Stdout) }
