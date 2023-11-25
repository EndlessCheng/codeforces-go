package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1179A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, m int
	Fscan(in, &n, &q)
	a := make([]int, n, n*2-1)
	for i := range a {
		Fscan(in, &a[i])
	}

	b := make([]int, 1, n)
	b[0] = a[0]
	for i := 1; i < n; i++ {
		x, y := b[i-1], a[i]
		if x < y {
			x, y = y, x
		}
		b = append(b, x)
		a = append(a, y)
	}

	for ; q > 0; q-- {
		Fscan(in, &m)
		if m <= n {
			Fprintln(out, b[m-1], a[m])
		} else {
			Fprintln(out, b[n-1], a[n+(m-n)%(n-1)])
		}
	}
}

//func main() { CF1179A(os.Stdin, os.Stdout) }
