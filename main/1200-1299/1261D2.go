package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1261D2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	const mod2 = (mod + 1) / 2
	const mx int = 2e5
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}
	pow := func(x int64, n int) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	invF := [...]int64{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % mod
	}
	C := func(n, k int) int64 { return F[n] * invF[k] % mod * invF[n-k] % mod }

	var n, k, m int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if i > 0 && a[i] != a[i-1] {
			m++
		}
	}
	if a[0] != a[n-1] {
		m++
	}

	ans := int64(0)
	for i := 0; i <= m/2; i++ {
		ans = (ans + C(m, i)%mod*C(m-i, i)%mod*pow(int64(k-2), m-2*i)) % mod
	}
	ans = (pow(int64(k), m) + mod - ans) * mod2 % mod * pow(int64(k), n-m) % mod
	Fprint(out, ans)
}

//func main() { CF1261D2(os.Stdin, os.Stdout) }
