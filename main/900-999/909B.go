package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF909B(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	Fprint(out, (n+1)/2*(n/2+1))
}

//func main() { CF909B(os.Stdin, os.Stdout) }
