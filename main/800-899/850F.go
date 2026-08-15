package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf850F(in io.Reader, out io.Writer) {
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
	a := make([]int, n+1)
	s := 0
	mx := 0
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		s += a[i]
		mx = max(mx, a[i])
	}

	f := make([]int, mx+1)
	f[1] = (s - 1) * (s - 1) % mod * pow(s, mod-2) % mod
	for i := 1; i < mx; i++ {
		f[i+1] = (2*f[i] - f[i-1] - (s-1)*pow(s-i, mod-2)) % mod
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += f[a[i]]
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { cf850F(bufio.NewReader(os.Stdin), os.Stdout) }
