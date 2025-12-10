package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1726E(in io.Reader, out io.Writer) {
	const mod = 998244353
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

	const mx int = 3e5
	fac := [mx + 1]int{}
	fac[0] = 1
	for i := 1; i <= mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	invF := [mx + 1]int{}
	invF[mx] = pow(fac[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	f := [mx + 1]int{1, 1}
	for i := 2; i <= mx; i++ {
		f[i] = (f[i-1] + (i-1)*f[i-2]) % mod
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		for i := range n/4 + 1 {
			ans = (ans + fac[n-i*2]*invF[i]%mod*invF[n-i*4]%mod*f[n-i*4]) % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1726E(bufio.NewReader(os.Stdin), os.Stdout) }
