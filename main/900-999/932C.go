package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF932C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, a, b int
	Fscan(in, &n, &a, &b)
	for ka := 0; ka <= n; ka += a {
		if (n-ka)%b > 0 {
			continue
		}
		for r := a; r <= ka; r += a {
			l := r - a + 1
			for i := l + 1; i <= r; i++ {
				Fprint(out, i, " ")
			}
			Fprint(out, l, " ")
		}
		for r := ka + b; r <= n; r += b {
			l := r - b + 1
			for i := l + 1; i <= r; i++ {
				Fprint(out, i, " ")
			}
			Fprint(out, l, " ")
		}
		return
	}
	Fprint(out, -1)
}

//func main() { CF932C(os.Stdin, os.Stdout) }
