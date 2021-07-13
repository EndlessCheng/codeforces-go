package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF78C(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m, &k)
	if n&1 > 0 {
		if k == 1 && m > 1 {
			Fprint(out, "Timur")
			return
		}
		for d := 2; d*d <= m; d++ {
			if m%d == 0 {
				if m/d >= k {
					Fprint(out, "Timur")
					return
				}
				break
			}
		}
	}
	Fprint(out, "Marsel")
}

//func main() { CF78C(os.Stdin, os.Stdout) }
