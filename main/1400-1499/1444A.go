package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1444A(in io.Reader, out io.Writer) {
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}
	var T, p, q int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &p, &q)
		if p%q > 0 {
			Fprintln(out, p)
			continue
		}
		ans := int64(0)
		for i := int64(2); i*i <= q; i++ {
			if q%i > 0 {
				continue
			}
			pp := p
			for q /= i; q%i == 0; q /= i {
				pp /= i
			}
			x := p
			for ; pp%i == 0; pp /= i {
				x /= i
			}
			ans = max(ans, x)
		}
		if q > 1 {
			for p%q == 0 {
				p /= q
			}
			ans = max(ans, p)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1444A(os.Stdin, os.Stdout) }
