package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1528B(in io.Reader, out io.Writer) {
	const mod = 998244353
	const mx int = 1e6
	d := [mx + 1]int{}
	for i := 1; i <= mx; i++ {
		for j := i; j <= mx; j += i {
			d[j]++
		}
	}

	var n int
	Fscan(in, &n)
	s := int64(d[n])
	for i, p2 := n-1, int64(1); i > 0; i-- {
		s = (s + p2*int64(d[i])) % mod
		p2 = p2 << 1 % mod
	}
	Fprint(out, s)
}

//func main() { CF1528B(os.Stdin, os.Stdout) }
