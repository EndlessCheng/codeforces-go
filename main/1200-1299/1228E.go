package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1228E(in io.Reader, out io.Writer) {
	const mod = 1e9 + 7
	pow := func(x int64, n int) (res int64) {
		res = 1
		for x %= mod; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}
	const mx = 250
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}
	invF := [...]int64{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % mod
	}
	C := func(n, k int) int64 { return F[n] * invF[k] % mod * invF[n-k] % mod }

	var n int
	var k, ans int64
	Fscan(in, &n, &k)
	// 容斥：枚举至少 i 列不合法，相当于每行中钦定了有 i 个数最小值大于 1
	for i := 0; i <= n; i++ {
		ans += int64(1-i&1*2) * C(n, i) * pow(pow(k, n-i)*pow(k-1, i)-pow(k-1, n), n) % mod
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { CF1228E(os.Stdin, os.Stdout) }
