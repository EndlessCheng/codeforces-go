package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
const mod = 1_000_000_007

func run(in io.Reader, out io.Writer) {
	const mx int = 2e5
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF := [...]int{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int { return F[n] * invF[k] % mod * invF[n-k] % mod }

	var n, k, ans int
	Fscan(in, &n, &k)
	for i := 0; i <= min(k, n-1); i++ {
		ans = (ans + C(n, i)*C(n-1, i)) % mod
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
