package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2203D(in io.Reader, out io.Writer) {
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

	for range rd() {
		n := rd()
		m := rd()
		mx := n + m + 1
		cnt := make([]int32, mx)
		for range n {
			cnt[rd()]++
		}

		div := make([]int32, mx)
		for i, c := range cnt {
			if c > 0 {
				for j := i; j < mx; j += i {
					div[j] += c
				}
			}
		}

		alice, bob := 0, 0
		for range m {
			c := div[rd()]
			if c == int32(n) {
				alice++
			} else if c == 0 {
				bob++
			}
		}

		common := m - alice - bob
		if alice+common%2 > bob {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}
}

//func main() { cf2203D(os.Stdin, os.Stdout) }
