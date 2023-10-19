package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF750C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	
	var n, delta, div, s int
	mn, mx := int(-1e9), int(1e9)
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &delta, &div)
		if div == 1 {
			mn = max(mn, 1900-s)
		} else {
			mx = min(mx, 1899-s)
		}
		s += delta
	}
	if mn > mx {
		Fprint(out, "Impossible")
	} else if mx == 1e9 {
		Fprint(out, "Infinity")
	} else {
		Fprint(out, mx+s)
	}
}

//func main() { CF750C(os.Stdin, os.Stdout) }
