package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1612G(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, c, tot int
	Fscan(in, &n)
	fac := make([]int, n+1)
	fac[0] = 1
	const mx int = 1e6
	cnt := [mx + 1]int{}
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i % mod
		Fscan(in, &c)
		cnt[c]++
		tot += c
	}

	ans := 0
	ways := 1
	for i := mx; i > 1; i-- {
		cc := cnt[i]
		if cc == 0 {
			continue
		}
		ans = (ans + (tot-cc)%mod*cc%mod*(i-1)) % mod
		ways = ways * fac[cc] % mod * fac[cc] % mod
		tot -= cc * 2
		cnt[i-2] += cc
	}
	ways = ways * fac[cnt[1]] % mod
	Fprint(out, (ans+mod)%mod, ways)
}

//func main() { cf1612G(bufio.NewReader(os.Stdin), os.Stdout) }
