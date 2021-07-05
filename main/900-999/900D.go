package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF900D(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	pow := func(x int64, n int) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	var g, n int
	Fscan(in, &g, &n)
	if n%g > 0 {
		Fprint(out, 0)
		return
	}
	n /= g
	ps := []int{}
	x := n
	for i := 2; i <= x; i++ {
		if x%i == 0 {
			ps = append(ps, i)
			for x /= i; x%i == 0; x /= i {
			}
		}
	}
	if x > 1 {
		ps = append(ps, x)
	}
	ans := int64(0)
	var f func(p, mu, d int)
	f = func(p, mu, d int) {
		if p == len(ps) {
			ans += int64(mu)*pow(2, n/d-1)%mod + mod
			return
		}
		f(p+1, mu, d)
		f(p+1, -mu, d*ps[p])
	}
	f(0, 1, 1)
	Fprint(out, ans%mod)
}

//func main() { CF900D(os.Stdin, os.Stdout) }
