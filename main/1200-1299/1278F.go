package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1278F(in io.Reader, out io.Writer) {
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
	var n, m, k, ans int
	Fscan(in, &n, &m, &k)
	s2 := make([]int, k+1)
	s2[0] = 1
	for i := 1; i <= k; i++ {
		for j := i; j > 0; j-- {
			s2[j] = (s2[j-1] + s2[j]*j) % mod
		}
		s2[0] = 0
	}

	m1 := pow(m, mod-2)
	perm, powM := 1, 1
	for i := range k + 1 {
		ans = (ans + s2[i]*perm%mod*powM) % mod
		perm = perm * (n - i) % mod
		powM = powM * m1 % mod
	}
	Fprint(out, ans)
}

//func main() { cf1278F(os.Stdin, os.Stdout) }
