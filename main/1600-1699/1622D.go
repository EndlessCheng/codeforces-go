package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1622D(in io.Reader, out io.Writer) {
	const mod, mx = 998244353, 5000
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
		return
	}
	invF := [...]int64{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % mod
	}
	C := func(n, k int) int64 { return F[n] * invF[k] % mod * invF[n-k] % mod }

	var n, k int
	var s string
	Fscan(bufio.NewReader(in), &n, &k, &s)
	pos := []int{}
	for i, b := range s {
		if b == '1' {
			pos = append(pos, i)
		}
	}
	if k == 0 || len(pos) < k {
		Fprint(out, 1)
		return
	}
	pos = append(pos, n)
	ans := C(pos[k], k)
	for i, p := range pos[:len(pos)-k-1] {
		ans += C(pos[i+k+1]-p-1, k) - C(pos[i+k]-p-1, k-1)
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { CF1622D(os.Stdin, os.Stdout) }
