package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1151C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	f := func(n int64) (res int64) {
		var s, o, e, st int64
		odd := true
		for i := int64(1); s < n; i <<= 1 {
			if odd {
				st = 2*o + 1
				o = (o + i) % mod
			} else {
				st = 2*e + 2
				e = (e + i) % mod
			}
			if s+i > n {
				i = n - s
			}
			s += i
			res += i % mod * (st + i%mod - 1) % mod
			odd = !odd
		}
		return
	}

	var l, r int64
	Fscan(in, &l, &r)
	Fprint(out, ((f(r)-f(l-1))%mod+mod)%mod)
}

//func main() { CF1151C(os.Stdin, os.Stdout) }
