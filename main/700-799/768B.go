package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF768B(in io.Reader, out io.Writer) {
	var n, l, r uint64
	Fscan(in, &n, &l, &r)
	tot := (uint64(1)<<bits.Len64(n) - 1) >> 1
	f := func(m uint64) (res uint64) {
		for n, tot := n, tot; m > 0; n >>= 1 {
			if m >= tot {
				m -= tot
				res += n / 2
				if m > 0 {
					m--
					res += n & 1
				}
			}
			tot >>= 1
		}
		return
	}
	Fprint(out, f(r)-f(l-1))
}

//func main() { CF768B(os.Stdin, os.Stdout) }
