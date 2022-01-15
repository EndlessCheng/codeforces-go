package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF954B(in io.Reader, out io.Writer) {
	var n int
	var s string
	Fscan(in, &n, &s)
	i := n / 2
	for s[:i] != s[i:i*2] {
		i--
	}
	if i == 0 {
		Fprint(out, n)
	} else {
		Fprint(out, n-i+1)
	}
}

//func main() { CF954B(os.Stdin, os.Stdout) }
