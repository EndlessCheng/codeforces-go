package main

import (
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF544B(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	if k > (n*n+1)/2 {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for i := 0; i < n; i++ {
		s := bytes.Repeat([]byte{'S'}, n)
		for j := i & 1; j < n && k > 0; j += 2 {
			s[j] = 'L'
			k--
		}
		Fprintf(out, "%s\n", s)
	}
}

//func main() { CF544B(os.Stdin, os.Stdout) }
