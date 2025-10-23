package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p162(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
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

	const mod = 998244352
	x := rd()
	const B = 1 << 16
	px := [B]int{1}
	for i := 1; i < B; i++ {
		px[i] = px[i-1] * x % mod
	}
	x2 := px[B-1] * x % mod
	px2 := [B]int{1}
	for i := 1; i < B; i++ {
		px2[i] = px2[i-1] * x2 % mod
	}
	for n := rd(); n > 0; n-- {
		e := rd()
		Fprint(out, px2[e/B]*px[e%B]%mod, " ")
	}
}

//func main() { p162(os.Stdin, os.Stdout) }
