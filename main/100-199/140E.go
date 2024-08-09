package main

import (
	. "fmt"
	"io"
)

func cf140E(in io.Reader, out io.Writer) {
	var n, m, mod, preL, l int
	Fscan(in, &n, &m, &mod)

	const mx = 5001
	var p, fac, g [mx]int
	p[0] = 1
	fac[0] = 1
	for i := 1; i < mx; i++ {
		p[i] = p[i-1] * (m + 1 - i) % mod
		fac[i] = fac[i-1] * i % mod
	}

	f := [mx][mx]int{}
	f[0][0] = 1
	for i := 1; i < mx; i++ {
		for j := 1; j <= i; j++ {
			f[i][j] = (f[i-1][j-1] + f[i-1][j]*(j-1)) % mod
		}
	}

	sum := 1
	for ; n > 0; n-- {
		Fscan(in, &l)
		s := 0
		for j := 1; j <= min(m, l); j++ {
			g[j] = (sum*p[j] - g[j]*fac[j]) % mod * f[l][j] % mod
			s += g[j]
		}
		for j := l + 1; j <= preL; j++ {
			g[j] = 0
		}
		sum = s % mod
		preL = l
	}
	Fprint(out, (sum+mod)%mod)
}

//func main() { cf140E(bufio.NewReader(os.Stdin), os.Stdout) }
