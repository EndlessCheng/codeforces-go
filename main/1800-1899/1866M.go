package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1866M(in io.Reader, out io.Writer) {
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

	var p, np, sum [100]int
	inv := pow(100, mod-2)
	for i := range p {
		p[i] = i * inv % mod
		np[i] = pow((1-p[i]+mod)%mod, mod-2)
	}

	var n, v, ans, e int
	Fscan(in, &n)
	for range n {
		for j := range sum {
			sum[j] = (sum[j] + e) * p[j] % mod
		}
		Fscan(in, &v)
		e = (1 + p[v]*(1+sum[v])%mod*np[v]) % mod
		ans += e
	}
	Fprint(out, ans%mod)
}

//func main() { cf1866M(bufio.NewReader(os.Stdin), os.Stdout) }
