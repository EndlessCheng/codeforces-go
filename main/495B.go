package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF495B(in io.Reader, out io.Writer) {
	var a, b, c int
	if Fscan(in, &a, &b); a == b {
		Fprint(out, "infinity")
		return
	}
	a -= b
	for d := 1; d*d <= a; d++ {
		if a%d == 0 {
			if d > b {
				c++
			}
			if d*d < a && a/d > b {
				c++
			}
		}
	}
	Fprint(out, c)
}

//func main() { CF495B(os.Stdin, os.Stdout) }
