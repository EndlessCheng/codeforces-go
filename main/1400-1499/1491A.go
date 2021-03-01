package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1491A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, s, op, x int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}
	for ; q > 0; q-- {
		Fscan(in, &op, &x)
		x--
		if op == 1 {
			a[x] ^= 1
			s += 2*a[x] - 1
		} else if x < s {
			Fprintln(out, 1)
		} else {
			Fprintln(out, 0)
		}
	}
}

//func main() { CF1491A(os.Stdin, os.Stdout) }
