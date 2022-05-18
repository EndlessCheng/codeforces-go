package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF197A(in io.Reader, out io.Writer) {
	var a, b, r int
	Fscan(in, &a, &b, &r)
	if r*2 <= a && r*2 <= b {
		Fprint(out, "First")
	} else {
		Fprint(out, "Second")
	}
}

//func main() { CF197A(os.Stdin, os.Stdout) }
