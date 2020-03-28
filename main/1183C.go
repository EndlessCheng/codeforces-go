package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1183C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, k, n, a, b int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &k, &n, &a, &b)
		k -= b*n + 1
		if k < 0 {
			Fprintln(out, -1)
		} else {
			if k/(a-b) < n {
				n = k / (a - b)
			}
			Fprintln(out, n)
		}
	}
}

//func main() { CF1183C(os.Stdin, os.Stdout) }
