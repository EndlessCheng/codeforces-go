package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1147C(in io.Reader, out io.Writer) {
	var n, v, c int
	min := 99
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		if Fscan(in, &v); v < min {
			min, c = v, 1
		} else if v == min {
			c++
		}
	}
	if c > n/2 {
		Fprint(out, "Bob")
	} else {
		Fprint(out, "Alice")
	}
}

//func main() { CF1147C(os.Stdin, os.Stdout) }
