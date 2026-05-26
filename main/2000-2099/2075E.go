package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2075E(in io.Reader, out io.Writer) {
	const mod = 998244353
	pow := func(x, n int) int {
		res := 1 % mod
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var T, n, m, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &a, &b)
		n = (pow(2, n) - 2 + mod) % mod
		m = (pow(2, m) - 2 + mod) % mod
		a++
		b++

		ans := (a*b + b*(b-1)/2%mod*m%mod*a + a*(a-1)/2%mod*n%mod*b) % mod
		for i := range 30 {
			mask := 1<<i - 1
			x := a >> (i + 1) << i
			if a>>i&1 > 0 {
				x += a & mask
			}
			y := b >> (i + 1) << i
			if b>>i&1 > 0 {
				y += b & mask
			}
			ans = (ans + x*y%mod<<i%mod*n%mod*m) % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2075E(bufio.NewReader(os.Stdin), os.Stdout) }
