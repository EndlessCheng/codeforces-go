package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1396A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int64
	Fscan(in, &n)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	if n == 1 {
		Fprint(out, "1 1\n", -a[0], "\n1 1\n0\n1 1\n0")
		return
	}
	Fprintln(out, "1 1\n", -a[0])
	a[0] = 0
	Fprintln(out, 2, n)
	for i, v := range a[1:] {
		Fprint(out, v*(n-1), " ")
		a[i+1] += v * (n - 1)
	}
	Fprintln(out)
	Fprintln(out, 1, n)
	for _, v := range a {
		Fprint(out, -v, " ")
	}
}

//func main() { CF1396A(os.Stdin, os.Stdout) }
