package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1444A(in io.Reader, out io.Writer) {
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
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
			e := 0
			for ; q%i == 0; q /= i {
				e++
			}
			if e > 0 {
				c := 0
				for v := p; v%i == 0; v /= i {
					c++
				}
				v := p
				for ; c >= e; c-- {
					v /= i
				}
				ans = max(ans, v)
			}
		}
		if q > 1 && p/q > ans {
			for ; p%q == 0; p /= q {
			}
			ans = max(ans, p)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1444A(os.Stdin, os.Stdout) }
