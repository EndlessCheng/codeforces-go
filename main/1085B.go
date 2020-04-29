package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1085B(_r io.Reader, _w io.Writer) {
	var n, k int
	Fscan(_r, &n, &k)
	ans := int(2e9)
	for i := k - 1; i > 0; i-- {
		if n%i == 0 {
			if v := n/i*k + i; v < ans {
				ans = v
			}
		}
	}
	Fprint(_w, ans)
}

//func main() { CF1085B(os.Stdin, os.Stdout) }
