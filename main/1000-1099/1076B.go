package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1076B(in io.Reader, out io.Writer) {
	var n int64
	Fscan(in, &n)
	if n&1 == 0 {
		Fprint(out, n/2)
		return
	}
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			Fprint(out, 1+(n-i)/2)
			return
		}
	}
	Fprint(out, 1)
}

//func main() { CF1076B(os.Stdin, os.Stdout) }
