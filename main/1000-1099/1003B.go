package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1003B(in io.Reader, out io.Writer) {
	var a, b, x int
	Fscan(in, &a, &b, &x)
	k := (x + 1) / 2
	a -= k
	b -= k
	R := strings.Repeat
	s := R("01", k)
	if x&1 > 0 {
		s = R("0", a) + s + R("1", b)
	} else if a > 0 {
		s += R("1", b) + R("0", a)
	} else {
		s = R("1", b) + s // WA 了一发：需要判断 a == 0 的情况
	}
	Fprint(out, s)
}

//func main() { CF1003B(os.Stdin, os.Stdout) }
