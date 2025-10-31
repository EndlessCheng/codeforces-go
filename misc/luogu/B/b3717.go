package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func b3717(in io.Reader, out io.Writer) {
	buf := make([]byte, 4096)
	_i, _n := 0, 0
	rc := func() byte {
		if _i == _n {
			_n, _ = in.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	const mod = 998244353
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	
	T := r()
	mx := r()
	F := make([]int, mx+1)
	F[0] = 1
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF := make([]int, mx+1)
	invF[mx] = pow(F[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int { return F[n] * invF[k] % mod * invF[n-k] % mod }
	xor := 0
	for ; T > 0; T-- {
		xor ^= C(r(), r())
	}
	Fprint(out, xor)
}

//func main() { b3717(os.Stdin, os.Stdout) }
