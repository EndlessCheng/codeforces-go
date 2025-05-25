package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1542E2(in io.Reader, out io.Writer) {
	var n, mod, ans int
	Fscan(in, &n, &mod)

	perm := 1
	perms := make([]int, n+1)
	for i := n; i > 3; i-- {
		perms[i] = perm
		perm = perm * i % mod
	}

	bias := n * (n - 1) / 2
	m := bias*2 + n
	f := make([]int, m)
	s := make([]int, m+1)
	for i := 1; i < n; i++ {
		k := i * (i - 1) / 2
		f[bias-k] = 1
		for j := bias - k + 1; j <= bias+k; j++ {
			f[j] = (f[j-1] + s[j+i] - s[j]*2 + s[max(j-i, 0)]) % mod
		}
		for j, v := range f {
			s[j+1] = s[j] + v
		}

		res := 0
		for delta := 1; delta <= i; delta++ {
			res = (res + (s[m]-s[bias+delta+1])%mod*(i+1-delta)) % mod
		}
		ans = (ans + perms[i+1]*res) % mod
	}
	Fprint(out, (ans+mod)%mod)
}

//func main() { cf1542E2(os.Stdin, os.Stdout) }
