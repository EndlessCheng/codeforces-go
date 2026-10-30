package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf204D(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, k int
	var s string
	Fscan(in, &n, &k, &s)
	s = " " + s
	f := make([]int, n+2)
	g := make([]int, n+2)
	p := make([]int, n+2)
	m, r, sum, ans := n, n, 1, 0
	p[n+1] = 1
	f[n] = 1
	for i := n; i > 0; i-- {
		p[i] = p[i+1]
		if s[i] == 'X' {
			p[i] = p[i+1] * 2 % mod
		}
	}

	for i := 1; i <= n; i++ {
		if s[i] == 'B' {
			m--
		}
		if s[i] == 'W' {
			r = m
			f[r] = sum
		}
		if s[i] == 'X' {
			m--
			f[m] = sum
			sum = sum * 2 % mod
		}
		if r-m == k {
			g[i] = f[r]
			sum = (sum - f[r] + mod) % mod
			r--
		}
	}

	clear(f)
	sum = 0
	m, r = n, n

	for i := 1; i <= n; i++ {
		if s[i] == 'W' {
			m--
		}
		if s[i] == 'B' {
			r = m
			f[r] = sum
		}
		if s[i] == 'X' {
			m--
			f[m] = sum
			sum = sum * 2 % mod
		}
		if r-m == k {
			ans = (ans + p[i+1]*f[r]) % mod
			sum = (sum - f[r] + mod) % mod
			r--
		}
		f[m] = (f[m] + g[i]) % mod
		sum = (sum + g[i]) % mod
	}

	Fprint(out, (ans+mod)%mod)
}

//func main() { cf204D(bufio.NewReader(os.Stdin), os.Stdout) }
