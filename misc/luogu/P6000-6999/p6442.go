package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func p6442(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
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
	rd := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n, w := rd(), rd()
	f := make([]int, 1<<w)
	pow2 := make([]int, n+1)
	pow2[0] = 1
	for i := 1; i <= n; i++ {
		s := 0
		for k := rd(); k > 0; k-- {
			s |= 1 << (rd() - 1)
		}
		f[s]++
		pow2[i] = pow2[i-1] * 2 % mod
	}

	for i := 0; i < w; i++ {
		for s := 0; s < 1<<w; s++ {
			s |= 1 << i
			f[s] += f[s^1<<i]
		}
	}

	ans := 0
	for i, fv := range f {
		sign := 1 - bits.OnesCount(uint(1<<w-1^i))%2*2
		ans += sign * pow2[fv]
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { p6442(os.Stdin, os.Stdout) }
