package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF675E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i]--
	}
	a[n-1] = n - 1

	type pair struct{ v, i int }
	st := make([][17]pair, n)
	for i, v := range a {
		st[i][0] = pair{v, i}
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			if a, b := st[i][j-1], st[i+1<<(j-1)][j-1]; a.v > b.v {
				st[i][j] = a
			} else {
				st[i][j] = b
			}
		}
	}
	query := func(l, r int) int {
		k := bits.Len(uint(r-l)) - 1
		a, b := st[l][k], st[r-1<<k][k]
		if a.v > b.v {
			return a.i
		}
		return b.i
	}

	dp := make([]int64, n-1)
	for i := range dp {
		dp[i] = -1
	}
	var f func(int) int64
	f = func(i int) (res int64) {
		if i == n-1 {
			return
		}
		dv := &dp[i]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		j := query(i, a[i]+1) // 查询最值所处下标
		return f(j) + int64(n-1-i-(a[i]-j))
	}
	ans := int64(0)
	for i := range a {
		ans += f(i)
	}
	Fprint(out, ans)
}

//func main() { CF675E(os.Stdin, os.Stdout) }
