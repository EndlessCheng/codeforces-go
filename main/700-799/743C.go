package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF743C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	if n == 1 {
		Fprint(out, -1)
	} else {
		Fprint(out, n, n+1, n*(n+1))
	}
}

//func main() { CF743C(os.Stdin, os.Stdout) }
