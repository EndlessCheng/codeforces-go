package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
const mod int64 = 1e9 + 7

type matrix [3][3]int64

func (a matrix) mul(b matrix) (c matrix) {
	for i, r := range a {
		for j := range b[0] {
			for k, v := range r {
				c[i][j] = (c[i][j] + v*b[k][j]) % (mod - 1)
			}
		}
	}
	return
}

func (a matrix) pow(n int64) (res matrix) {
	for i := range res {
		res[i][i] = 1
	}
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return
}

func CF1182E(in io.Reader, out io.Writer) {
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
	var n, f1, f2, f3, c int64
	Fscan(in, &n, &f1, &f2, &f3, &c)
	f1 = f1 * c % mod
	f2 = f2 * c % mod * c % mod
	f3 = f3 * c % mod * c % mod * c % mod
	m := matrix{
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
	}.pow(n - 3)
	Fprint(out, pow(c, mod-1-n%(mod-1))*
		pow(f3, m[0][0])%mod*
		pow(f2, m[0][1])%mod*
		pow(f1, m[0][2])%mod)
}

//func main() { CF1182E(os.Stdin, os.Stdout) }
