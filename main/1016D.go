package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1016D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, s, s0 int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s ^= a[i]
		if i > 0 {
			s0 ^= a[i]
		}
	}
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
		s ^= b[i]
	}

	if s > 0 {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	Fprint(out, s0^b[0])
	for j := 1; j < m; j++ {
		Fprint(out, " ", b[j])
	}
	Fprintln(out)
	for i := 1; i < n; i++ {
		Fprintln(out, a[i], strings.Repeat("0 ", m-1))
	}
}

//func main() { CF1016D(os.Stdin, os.Stdout) }
