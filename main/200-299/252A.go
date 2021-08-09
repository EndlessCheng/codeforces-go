package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF252A(in io.Reader, out io.Writer) {
	var n, v, ans int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		a[i+1] = a[i] ^ v
	}
	for i, v := range a {
		for _, w := range a[:i] {
			if v^w > ans {
				ans = v ^ w
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF252A(os.Stdin, os.Stdout) }
