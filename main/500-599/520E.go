package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF520E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	const mx int = 1e5
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}
	pow := func(x, n int64) (res int64) {
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

	var n, k int
	var s string
	Fscan(in, &n, &k, &s)
	sum := make([]int64, n+1)
	for i, b := range s {
		sum[i+1] = sum[i] + int64(b&15)
	}
	ans := int64(0)
	for i, p10 := 0, int64(1); i < n-k; i++ { // 枚举 10^i，计算 10^i 可以算到哪些数字的贡献上
		if k > 0 {
			ans = (ans + sum[n-i-1]*p10%mod*C(n-i-2, k-1)) % mod
		}
		ans = (ans + (sum[n-i]-sum[n-i-1])*p10%mod*C(n-i-1, k)) % mod // 作为最后一个数的部分，后面无法跟加号
		p10 = p10 * 10 % mod
	}
	Fprint(out, ans)
}

//func main() { CF520E(os.Stdin, os.Stdout) }
