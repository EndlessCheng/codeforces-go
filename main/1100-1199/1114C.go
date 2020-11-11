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
		c := int64(0)
		for pp := p; ; pp *= p {
			c += n / pp
			if pp > n/p { // 注意判断方式，防止爆 int64
				break
			}
		}
		if c/e < ans {
			ans = c / e
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
