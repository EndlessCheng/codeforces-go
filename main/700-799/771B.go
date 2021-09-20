package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF771B(in io.Reader, out io.Writer) {
	var n, k int
	var s string
	Fscan(in, &n, &k)
	a := make([][]byte, n)
	for i := 0; i < k-1; i++ {
		a[i] = []byte{'A' + byte(i/26), 'a' + byte(i%26)}
	}
	for i := k - 1; i < n; i++ {
		if Fscan(in, &s); s[0] == 'Y' {
			a[i] = []byte{'A' + byte(i/26), 'a' + byte(i%26)}
		} else {
			a[i] = a[i-k+1]
		}
	}
	for _, v := range a {
		Fprintf(out, "%s ", v)
	}
}

//func main() { CF771B(os.Stdin, os.Stdout) }
