package main

import (
	. "fmt"
	"io"
)

// 注：O(1) space 做法 https://codeforces.com/problemset/submission/1109/120376830

// https://github.com/EndlessCheng
func cf1109D(in io.Reader, out io.Writer) {
	const M = 1_000_000_007
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
	const mx int = 1e6
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % M
	}
	invF := [...]int{mx: pow(F[mx], M-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % M
	}
	P := func(n, k int) int { return F[n] * invF[n-k] % M }
	C := func(n, k int) int { return F[n] * invF[k] % M * invF[n-k] % M }

	var n, m, ans int
	Fscan(in, &n, &m)
	u := min(n-1, m)
	powN := pow(n, n-u+M-3)
	powM := pow(m, n-u-1)
	for k := u; k > 0; k-- {
		ans = (ans + P(n-2, k-1)*(k+1)%M*powN%M*C(m-1, k-1)%M*powM) % M
		powN = powN * n % M
		powM = powM * m % M
	}
	Fprint(out, ans)
}

//func main() { cf1109D(os.Stdin, os.Stdout) }
