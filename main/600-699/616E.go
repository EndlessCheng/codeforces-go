package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF616E(in io.Reader, out io.Writer) {
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	const mod int64 = 1e9 + 7
	const mod2 = (mod + 1) / 2
	var k, n int64
	Fscan(in, &k, &n)
	sum := n % mod * (k % mod)
	for l, r := int64(1), int64(0); l <= n; l = r + 1 {
		h := k / l
		if h > 0 {
			r = min(k/h, n)
		} else {
			r = n
		}
		w := (r - l + 1) % mod
		s := (l + r) % mod * w % mod * mod2 % mod
		sum -= h % mod * s % mod
	}
	Fprint(out, (sum%mod+mod)%mod)
}

//func main() { CF616E(os.Stdin, os.Stdout) }
