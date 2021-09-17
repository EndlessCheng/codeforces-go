package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF996A(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	Fprint(out, n/100+n/20%5+n/10%2+n/5%2+n%5)
}

//func main() { CF996A(os.Stdin, os.Stdout) }
