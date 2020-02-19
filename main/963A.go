package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF963A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 9
	pow := func(x int64, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	inv := func(a int64) int64 { return pow(a, mod-2) }
	frac := func(a, b int64) int64 { return a * inv(b) % mod }

	var n, a, b, k, ans int64
	var s []byte
	Fscan(in, &n, &a, &b, &k, &s)
	q := frac(b, a)
	for i := len(s) - 1; i >= 0; i-- {
		ans *= q
		if s[i] == '+' {
			ans++
		} else {
			ans += mod - 1
		}
		ans %= mod
	}
	ans = ans * pow(a, n) % mod
	n = (n + 1) / k
	if q = pow(q, k); q == 1 {
		ans = ans * n % mod
	} else {
		ans = ans * frac(pow(q, n)-1, q-1) % mod
	}
	Fprint(out, ans)
}

//func main() { CF963A(os.Stdin, os.Stdout) }
