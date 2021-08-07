package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1065B(in io.Reader, out io.Writer) {
	var n, m, mx int64
	Fscan(in, &n, &m)
	mi := n - m*2
	if mi < 0 {
		mi = 0
	}
	for ; mx*(mx-1)/2 < m; mx++ {
	}
	Fprint(out, mi, n-mx)
}

//func main() { CF1065B(os.Stdin, os.Stdout) }
