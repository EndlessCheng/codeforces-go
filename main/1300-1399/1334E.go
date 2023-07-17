package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1334E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, q, v, w int64
	Fscan(in, &n, &q)
	ps := []int64{}
	x := n
	if x&1 == 0 {
		ps = append(ps, 2)
		x /= x & -x
	}
	for i := int64(3); i*i <= x; i += 2 {
		if x%i == 0 {
			for x /= i; x%i == 0; x /= i {
			}
			ps = append(ps, i)
		}
	}
	if x > 1 {
		ps = append(ps, x)
	}

	dp := map[int64]int{}
	var f func(n int64) int
	f = func(n int64) (res int) {
		if n == 1 {
			return 1
		}
		if v, ok := dp[n]; ok {
			return v
		}
		for _, p := range ps {
			if n%p == 0 {
				res = (res + f(n/p)) % mod
			}
		}
		dp[n] = res
		return
	}
	for ; q > 0; q-- {
		Fscan(in, &v, &w)
		g := gcd(v, w)
		Fprintln(out, int64(f(v/g))*int64(f(w/g))%mod)
	}
}

//func main() { CF1334E(os.Stdin, os.Stdout) }
