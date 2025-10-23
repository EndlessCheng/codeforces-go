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
	px := [1 << 16]int{1}
	for i := 1; i < len(px); i++ {
		px[i] = px[i-1] * x % mod
	}
	x2 := px[len(px)-1] * x % mod
	px2 := [1 << 16]int{1}
	for i := 1; i < len(px); i++ {
		px2[i] = px2[i-1] * x2 % mod
	}
	for n := rd(); n > 0; n-- {
		e := rd()
		Fprint(out, px2[e>>16]*px[e&0xffff]%mod, " ")
	}
}

//func main() { p162(os.Stdin, os.Stdout) }
