package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1523E(in io.Reader, out io.Writer) {
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

	const mx int = 1e5
	var f, invF [mx + 1]int
	f[0] = 1
	for i := 1; i <= mx; i++ {
		f[i] = f[i-1] * i % mod
	}
	invF[mx] = pow(f[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		ans := 1
		for t := 1; t < (n-1)/k+2; t++ {
			m := n - (k-1)*(t-1)
			ans = (ans + f[m]*f[n-t]%mod*invF[n]%mod*invF[m-t]) % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1523E(os.Stdin, os.Stdout) }
