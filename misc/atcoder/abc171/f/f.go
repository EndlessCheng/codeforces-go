package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
const mod int = 1e9 + 7

func run(in io.Reader, out io.Writer) {
	const mx int = 2e6
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF := [...]int{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int { return F[n] * invF[k] % mod * invF[n-k] % mod }

	var k, ans int
	var s string
	Fscan(bufio.NewReader(in), &k, &s)
	n := len(s)
	p26 := 1
	p25 := pow(25, k)
	const inv25 = 280000002 // pow(25, mod-2)
	for i := 0; i <= k; i++ {
		ans = (ans + p26*p25%mod*C(n-1+k-i, n-1)) % mod
		p26 = p26 * 26 % mod
		p25 = p25 * inv25 % mod
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
