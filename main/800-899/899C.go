package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF899C(in io.Reader, out io.Writer) {
	var n int64
	Fscan(in, &n)
	s := n * (n + 1) / 2
	Fprintln(out, s&1)
	s /= 2
	a := []interface{}{}
	for i := n; s > i; i-- {
		s -= i
		a = append(a, i)
	}
	a = append(a, s)
	Fprint(out, len(a), " ")
	Fprintln(out, a...)
}

//func main() { CF899C(os.Stdin, os.Stdout) }
