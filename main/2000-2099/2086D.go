package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2086D(in io.Reader, out io.Writer) {
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
	fac := func(n int) int {
		res := 1
		for i := 2; i <= n; i++ {
			res = res * i % mod
		}
		return res
	}

	var T int
	a := [26]int{}
	for Fscan(in, &T); T > 0; T-- {
		n := 0
		ans := 1
		for i := range a {
			Fscan(in, &a[i])
			n += a[i]
			ans = ans * fac(a[i]) % mod
		}
		m := n / 2
		ans = pow(ans, mod-2) * fac(m) % mod * fac(n-m) % mod

		f := make([]int, m+1)
		f[0] = 1
		s := 0
		for _, v := range a {
			if v == 0 {
				continue
			}
			s = min(s+v, m)
			for j := s; j >= v; j-- {
				f[j] = (f[j] + f[j-v]) % mod
			}
		}
		Fprintln(out, ans*f[m]%mod)
	}
}

//func main() { cf2086D(bufio.NewReader(os.Stdin), os.Stdout) }
