package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1039A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var t, v int64
	Fscan(in, &n, &t, &v)
	b := make([]int64, n-1, n)
	for i := range b {
		Fscan(in, &b[i])
		b[i] += t
	}
	b = append(b, 3e18)
	x := make([]int, n)
	for i := range x {
		Fscan(in, &x[i])
		if x[i]--; x[i] < i || i > 0 && x[i] < x[i-1] {
			Fprint(out, "No")
			return
		}
		if x[i] == i {
			b[i]--
		}
	}
	for i := 0; i < n-1; i++ {
		if b[i] == b[i+1] || x[i] > i && x[i] < x[i+1] {
			Fprint(out, "No")
			return
		}
	}
	Fprintln(out, "Yes")
	for _, v := range b {
		Fprint(out, v, " ")
	}
}

//func main() { CF1039A(os.Stdin, os.Stdout) }
