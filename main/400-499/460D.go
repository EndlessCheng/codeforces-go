package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF460D(in io.Reader, out io.Writer) {
	var l, r, k, xor int64
	var a []int64
	Fscan(in, &l, &r, &k)
	xor1 := func() { xor, a = 1, []int64{l + l&1, l + l&1 + 1} }
	xor0 := func() {
		for b := int64(1); b*3 <= r; b <<= 1 {
			if x, y, z := b*2-1, b*3-1, b*3; l <= x && z <= r {
				a = []int64{x, y, z}
				return
			}
		}
		xor1()
	}
	if r-l+1 < k {
		k = r - l + 1
	}
	switch {
	case k == 1:
		xor, a = l, []int64{l}
	case k == 2:
		if l&1 == 0 || l+2 <= r {
			xor1()
		} else if l < l^r {
			xor, a = l, []int64{l}
		} else {
			xor, a = l^r, []int64{l, r}
		}
	case k == 3:
		xor0()
	case l&1 == 0:
		a = []int64{l, l + 1, l + 2, l + 3}
	case l+4 <= r:
		a = []int64{l + 1, l + 2, l + 3, l + 4}
	default:
		xor0()
	}
	Fprintln(out, xor)
	Fprintln(out, len(a))
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF460D(os.Stdin, os.Stdout) }
