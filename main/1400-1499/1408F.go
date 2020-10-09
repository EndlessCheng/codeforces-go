package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1408F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var f func(l, r int)
	f = func(l, r int) {
		if r-l == 1 {
			Fprintln(out, l, r)
			return
		}
		m := (l + r) / 2
		f(l, m)
		f(m+1, r)
		for i := l; i <= m; i++ {
			Fprintln(out, i, m+i-l+1)
		}
	}
	var n int
	Fscan(in, &n)
	if n == 1 {
		Fprint(out, 0)
		return
	}
	d := bits.Len(uint(n)) - 1
	m := 1 << d
	Fprintln(out, d*m)
	f(1, m)
	f(n-m+1, n)
}

//func main() { CF1408F(os.Stdin, os.Stdout) }
