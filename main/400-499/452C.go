package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF452C(in io.Reader, out io.Writer) {
	var n, m float64
	Fscan(in, &n, &m)
	if n == 1 {
		Fprint(out, 1)
		return
	}
	Fprintf(out, "%.16f", (1+(n-1)*(m-1)/(n*m-1))/n)
}

//func main() { CF452C(os.Stdin, os.Stdout) }
