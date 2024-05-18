package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf451E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	invF := [20]int{1}
	m := 1
	for i := 1; i < 20; i++ {
		m = m * i % mod
		invF[i] = pow(m, mod-2)
	}
	comb := func(n, k int) int {
		if n < k {
			return 0
		}
		n %= mod
		p := 1
		for i := 1; i <= k; i++ {
			p = p * (n - i + 1) % mod
		}
		return p * invF[k] % mod
	}

	var n, s, tot, ans int
	Fscan(in, &n, &s)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		tot += a[i]
	}
	if tot < s {
		Fprint(out, 0)
		return
	}
	for i := uint(0); i < 1<<n; i++ {
		s2 := s
		for t := i; t > 0; t &= t - 1 {
			s2 -= a[bits.TrailingZeros(t)] + 1
		}
		res := comb(s2+n-1, n-1)
		if bits.OnesCount(i)%2 > 0 {
			res = -res
		}
		ans += res
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { cf451E(os.Stdin, os.Stdout) }
