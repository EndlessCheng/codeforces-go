package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1368A(in io.Reader, out io.Writer) {
	var t, a, b, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &a, &b, &n)
		if a > b {
			a, b = b, a
		}
		i := 0
		for ; a <= n && b <= n; i++ {
			a, b = b, a+b
		}
		Fprintln(out, i)
	}
}

//func main() { CF1368A(os.Stdin, os.Stdout) }
