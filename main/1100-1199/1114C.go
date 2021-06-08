package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1114C(in io.Reader, out io.Writer) {
	var n, x int64
	Fscan(in, &n, &x)
	ans := int64(1e18)
	f := func(p, e int64) {
		k := n
		for n := n; n > 0; n /= p {
			k -= n % p
		}
		k /= p - 1
		if k/e < ans {
			ans = k / e
		}
	}
	for i := int64(2); i*i <= x; i++ {
		e := int64(0)
		for ; x%i == 0; x /= i {
			e++
		}
		if e > 0 {
			f(i, e)
		}
	}
	if x > 1 {
		f(x, 1)
	}
	Fprint(out, ans)
}

//func main() { CF1114C(os.Stdin, os.Stdout) }
