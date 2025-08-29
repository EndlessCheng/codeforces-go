package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1550D(in io.Reader, out io.Writer) {
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

	const mx int = 2e5
	var F, invF [mx + 1]int
	F[0] = 1
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF[mx] = pow(F[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int {
		if k < 0 || k > n {
			return 0
		}
		return F[n] * invF[k] % mod * invF[n-k] % mod
	}

	var T, n, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r)
		mn := min(r-n, 1-l)
		ans := C(n, n/2) * mn * (n%2 + 1)
		d, u := 1-mn, n+mn
		for {
			d--
			u++
			m := n - max(l-d, 0) - max(u-r, 0)
			if m < 0 {
				break
			}
			pos := n/2 - max(l-d, 0)
			ans += C(m, pos) + n%2*C(m, pos+1)
		}
		Fprintln(out, ans%mod)
	}
}

//func main() { cf1550D(bufio.NewReader(os.Stdin), os.Stdout) }
