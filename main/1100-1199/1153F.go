package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1153F(in io.Reader, out io.Writer) {
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

	var n, k, l int
	Fscan(in, &n, &k, &l)
	f := make([]int, n*2+2)
	g := make([]int, n*2+2)
	f[0] = 1
	for i := 1; i <= n*2; i++ {
		for j := i & 1; j <= i; j += 2 {
			ex := 0
			if j > 0 {
				ex = f[j-1]
			}
			f[j] = (f[j+1]*(j+1) + ex) % mod
			ex = 0
			if j > 0 {
				ex = g[j-1]
			}
			ex2 := 0
			if j >= k {
				ex2 = f[j]
			}
			g[j] = (g[j+1]*(j+1) + ex + ex2) % mod
		}
	}

	Fprint(out, l*pow(n*2+1, mod-2)%mod*g[0]%mod*pow(f[0], mod-2)%mod)
}

//func main() { cf1153F(bufio.NewReader(os.Stdin), os.Stdout) }
