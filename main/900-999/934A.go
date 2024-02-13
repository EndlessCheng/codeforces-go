package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf934A(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}
	sort.Ints(b)

	if max(a[0]*b[0], a[0]*b[m-1]) > max(a[n-1]*b[0], a[n-1]*b[m-1]) {
		a = a[1:]
	} else {
		a = a[:n-1]
	}
	n--
	Fprint(out, max(a[0]*b[0], a[0]*b[m-1], a[n-1]*b[0], a[n-1]*b[m-1]))
}

//func main() { cf934A(os.Stdin, os.Stdout) }
