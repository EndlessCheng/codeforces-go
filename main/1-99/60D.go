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

	const mx = 10_000_000
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

	n := rd()
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

	for i := 1; i*(i+2) <= mx; i += 2 {
		// 如果 a > mx 或 b > mx，那么 c > max(a, b) > mx
		for j := i + 2; i*j <= mx && j*j-i*i <= mx*2; j += 2 {
			if gcd(i, j) > 1 {
				continue
			}
			a := uint32(i * j)
			b := uint32(j*j-i*i) / 2
			c := uint32(i*i+j*j) / 2
			merge(a, b)
			if c <= mx {
				merge(a, c)
				merge(b, c)
			}
		}
	}
	Fprint(out, n)
}

//func main() { cf60D(bufio.NewReader(os.Stdin), os.Stdout) }
