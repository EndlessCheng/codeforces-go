package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1370A(in io.Reader, out io.Writer) {
	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		Fprintln(out, n/2)
	}
}

//func main() { CF1370A(os.Stdin, os.Stdout) }
