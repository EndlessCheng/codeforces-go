package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF675B(in io.Reader, out io.Writer) {
	var n, a, b, c, d, ans int
	Fscan(in, &n, &a, &b, &c, &d)
	for i := 1; i <= n; i++ {
		s := a + b + i
		if !(s < a+c+1 || s > a+c+n || s < b+d+1 || s > b+d+n || s < c+d+1 || s > c+d+n) {
			ans++
		}
	}
	Fprint(out, int64(ans)*int64(n))
}

//func main() { CF675B(os.Stdin, os.Stdout) }
