package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"math/bits"
	"reflect"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func CF611D(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, ans int
	var s []byte
	Fscan(bufio.NewReader(in), &n, &s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[rk] = h
	}
	const mx = 13
	st := make([][mx]int, n)
	for i, v := range height {
		st[i][0] = v
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = min(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	_q := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return min(st[l][k], st[r-1<<k][k]) }
	lcp := func(i, j int) int {
		if i == j {
			return n - i
		}
		ri, rj := rank[i], rank[j]
		if ri > rj {
			ri, rj = rj, ri
		}
		return _q(ri+1, rj+1)
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			continue
		}
		for j, k, sum := i, i-1, 0; j < n; j++ {
			dp[i][j] = sum
			if k < 0 {
				continue
			}
			if s[k] > '0' && rank[k] < rank[i] && lcp(k, i) < i-k {
				dp[i][j] = (dp[i][j] + dp[k][i-1]) % mod
			}
			sum = (sum + dp[k][i-1]) % mod
			k--
		}
	}
	for _, d := range dp {
		ans = (ans + d[n-1]) % mod
	}
	Fprint(out, ans)
}

//func main() { CF611D(os.Stdin, os.Stdout) }
