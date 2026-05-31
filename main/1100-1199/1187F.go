package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1187F(in io.Reader, out io.Writer) {
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

	var n int
	Fscan(in, &n)
	l := make([]int, n+1)
	r := make([]int, n+1)
	sz := make([]int, n+1)
	sz[0] = 1
	for i := 1; i <= n; i++ {
		Fscan(in, &l[i])
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &r[i])
		sz[i] = pow(r[i]-l[i]+1, mod-2)
	}

	ans := 0
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		d := max(min(r[i], r[i-1])-max(l[i], l[i-1])+1, 0)
		p[i] = d * sz[i] % mod * sz[i-1] % mod
		ans = (ans + 1 - p[i]) % mod
	}

	s := 0
	for i := 3; i <= n; i++ {
		s = (s + 1 - p[i-2]) % mod
		ans = (ans + s*(1-p[i])*2) % mod
	}

	for i := 1; i < n; i++ {
		d := max(min(r[i+1], r[i], r[i-1])-max(l[i+1], l[i], l[i-1])+1, 0)
		v := 1 - p[i] - p[i+1] + d*sz[i+1]%mod*sz[i]%mod*sz[i-1]
		ans = (ans + v*2) % mod
	}
	Fprint(out, (ans+mod)%mod)
}

//func main() { cf1187F(bufio.NewReader(os.Stdin), os.Stdout) }
