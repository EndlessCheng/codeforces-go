package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF277B(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	if m == 3 && n > 4 {
		Fprint(out, -1)
		return
	}
	for i := 0; i < m; i++ {
		Fprintln(out, i, 1e7+i*i)
	}
	for i := 0; i < n-m; i++ {
		Fprintln(out, i, -1e7-i*i)
	}
}

//func main() { CF277B(os.Stdin, os.Stdout) }
