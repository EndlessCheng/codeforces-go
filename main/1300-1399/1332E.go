package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1332E(in io.Reader, out io.Writer) {
	const mod = 998244353
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
	var n, m, l, r int64
	Fscan(in, &n, &m, &l, &r)
	n *= m
	r -= l - 1
	if n&1 > 0 {
		Fprint(out, pow(r, n))
	} else {
		Fprint(out, (pow(r, n)+r&1)*(mod+1)/2%mod)
	}
}

//func main() { CF1332E(os.Stdin, os.Stdout) }
