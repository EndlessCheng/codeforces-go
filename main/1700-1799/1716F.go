package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1716F(in io.Reader, out io.Writer) {
	const M = 998244353
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % M
			}
			x = x * x % M
		}
		return res
	}
	const mx = 2001
	s2 := [mx][mx]int{}
	s2[0][0] = 1
	for i := 1; i < mx; i++ {
		for j := 1; j <= i; j++ {
			s2[i][j] = (s2[i-1][j-1] + s2[i-1][j]*j) % M
		}
	}

	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		o := (m + 1) / 2
		invM := pow(m, M-2)
		ans, p, po, pm := 0, 1, 1, pow(m, n)
		for i := 1; i <= min(k, n); i++ {
			p = p * (n - i + 1) % M
			po = po * o % M
			pm = pm * invM % M
			ans = (ans + s2[k][i]*p%M*po%M*pm) % M
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1716F(bufio.NewReader(os.Stdin), os.Stdout) }
