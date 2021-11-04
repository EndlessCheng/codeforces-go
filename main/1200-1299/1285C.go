package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1285C(in io.Reader, out io.Writer) {
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, ans int64
	Fscan(in, &n)
	for x := int64(1); x*x <= n; x++ {
		if n%x == 0 && gcd(x, n/x) == 1 {
			ans = x
		}
	}
	Fprintln(out, ans, n/ans)
}

//func main() { CF1285C(os.Stdin, os.Stdout) }
