package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf321E(in io.Reader, out io.Writer) {
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

	n, k := rd(), rd()
	s := make([][]int32, n+1)
	for i := range s {
		s[i] = make([]int32, n+1)
	}
	for i := range n {
		for j := range n {
			s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j] + int32(rd())
		}
	}

	f := make([]int32, n+1)
	opt := make([]int, n+2)
	opt[n+1] = n
	for range k - 1 {
		for i := n; i >= 1; i-- {
			f[i] = -1e9
			for j := opt[i]; j <= opt[i+1]; j++ {
				v := f[j] - s[j][j] + s[j][i]
				if v > f[i] {
					f[i] = v
					opt[i] = j
				}
			}
		}
	}
	Fprint(out, s[n][n]/2-f[n])
}

//func main() { cf321E(os.Stdin, os.Stdout) }
