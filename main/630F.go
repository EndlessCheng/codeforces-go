package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF630F(_r io.Reader, _w io.Writer) {
	c := func(n, k int) int64 {
		res := int64(1)
		for i := 1; i <= k; i++ {
			res = res * int64(n-k+i) / int64(i)
		}
		return res
	}
	var n int
	Fscan(_r, &n)
	Fprintln(_w, c(n, 5)+c(n, 6)+c(n, 7))
}

//func main() { CF630F(os.Stdin, os.Stdout) }
