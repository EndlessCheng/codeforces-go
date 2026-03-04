package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2174C2(in io.Reader, out io.Writer) {
	var T, n, m, mod int
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

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &mod)
		inv := make([]int, n+1)
		sum := make([]int, n+1)
		inv[0] = 1
		inv[1] = pow(m, mod-2)
		for i := 2; i <= n; i++ {
			inv[i] = inv[i-1] % mod * inv[1] % mod
		}
		sum[0] = 1
		for i := 1; i <= n; i++ {
			sum[i] = (sum[i-1] + inv[i]) % mod
		}

		ans := 0
		f := make([]int, n+1)
		for i := 1; i <= n; i++ {
			f[i] = (n + 1 - i) * inv[i/2] % mod
			ans += f[i]
		}
		ans %= mod
		ans *= ans
		for i := 2; i <= n; i++ {
			l, r := 1-i%2, i/2
			x := inv[r] * ((r-l+1)*2 - 1)
			y := (sum[r*2]-sum[r+l-1])*2 - inv[r+r]
			ans = (ans + (x-y)%mod*(n-i+1)) % mod
		}
		Fprintln(out, (ans+mod)%mod)
	}
}

//func main() { cf2174C2(os.Stdin, os.Stdout) }
