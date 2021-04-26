package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF468C(in io.Reader, out io.Writer) {
	var m int64
	Fscan(in, &m)
	l := m - 1e18%m*9%m*9%m
	Fprint(out, l, l+1e18-1)
}

//func main() { CF468C(os.Stdin, os.Stdout) }
