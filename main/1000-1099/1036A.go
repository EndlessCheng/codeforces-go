package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1036A(in io.Reader, out io.Writer) {
	var n, k int64
	Fscan(in, &n, &k)
	Fprint(out, (k-1)/n+1)
}

//func main() { CF1036A(os.Stdin, os.Stdout) }
