package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF570B(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	if n == 1 && m == 1 {
		Fprint(out, 1)
	} else if m*2 > n {
		Fprint(out, m-1)
	} else {
		Fprint(out, m+1)
	}
}

//func main() { CF570B(os.Stdin, os.Stdout) }
