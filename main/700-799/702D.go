package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF702D(in io.Reader, out io.Writer) {
	var d, k, a, b, t int64
	Fscan(in, &d, &k, &a, &b, &t)
	ans := b * d
	for _, x := range []int64{k * ((d - 1) / k), k, d} {
		if 0 < x && x <= d {
			if v := a*x + ((x-1)/k)*t + (d-x)*b; v < ans {
				ans = v
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF702D(os.Stdin, os.Stdout) }
