package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF478B(in io.Reader, out io.Writer) {
	var n, m int64
	Fscan(in, &n, &m)
	k := n / m
	Fprint(out, m*k*(k-1)/2+n%m*k, (n-m)*(n-m+1)/2)
}

//func main() { CF478B(os.Stdin, os.Stdout) }
