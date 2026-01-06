package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf60D(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
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
	rd := func() (x uint32) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + uint32(b&15)
		}
		return
	}

	const mx int = 1e7
	fa := [mx + 1]uint32{}
	find := func(x uint32) uint32 {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	n := int(rd())
	merge := func(x, y uint32) {
		if fa[x] == 0 || fa[y] == 0 {
			return
		}
		x = find(x)
		y = find(y)
		if x != y {
			fa[x] = y
			n--
		}
	}

	for range n {
		v := rd()
		fa[v] = v
	}

	for u := 3; u*u < mx*2; u += 2 {
		for v := 1; v < u && u*u+v*v <= mx*2; v += 2 {
			if gcd(u, v) > 1 {
				continue
			}
			a := uint32(u * v)
			b := uint32(u*u-v*v) / 2
			c := uint32(u*u+v*v) / 2
			merge(a, b)
			merge(a, c)
			merge(b, c)
		}
	}
	Fprint(out, n)
}

//func main() { cf60D(bufio.NewReader(os.Stdin), os.Stdout) }
