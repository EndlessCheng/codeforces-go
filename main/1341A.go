package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1341A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, a, b, c, d int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &a, &b, &c, &d)
		if c-d > n*(a+b) || c+d < n*(a-b) {
			Fprintln(out, "No")
		} else {
			Fprintln(out, "Yes")
		}
	}
}

//func main() { CF1341A(os.Stdin, os.Stdout) }
