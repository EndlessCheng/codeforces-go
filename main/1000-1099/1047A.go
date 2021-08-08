package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1047A(in io.Reader, out io.Writer) {
	var n int
	if Fscan(in, &n); n%3 < 2 {
		Fprint(out, 1, 1, n-2)
	} else {
		Fprint(out, 1, 2, n-3)
	}
}

//func main() { CF1047A(os.Stdin, os.Stdout) }
