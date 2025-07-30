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
	cc := [mx + 1]int{}
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i % mod
		Fscan(in, &c)
		cc[c]++
		tot += c
	}

	ans := 0
	ways := 1
	for c := mx; c > 1; c-- {
		k := cc[c]
		if k == 0 {
			continue
		}
		// k 种元素，每种选两个一左一右，剩余的放入 cc[c-2]
		ans = (ans + (tot-k)%mod*k%mod*(c-1)) % mod
		ways = ways * fac[k] % mod * fac[k] % mod // 左右独立全排列
		tot -= k * 2
		cc[c-2] += k
	}
	ways = ways * fac[cc[1]] % mod // 1 个数无贡献，只有全排列
	Fprint(out, (ans+mod)%mod, ways)
}

//func main() { cf1612G(bufio.NewReader(os.Stdin), os.Stdout) }
