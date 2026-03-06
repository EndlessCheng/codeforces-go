package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1244G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		a[i] = i + 1
		b[i] = i + 1
	}

	k0 := k
	k -= n * (n + 1) / 2
	if k < 0 {
		Fprintln(out, -1)
		return
	}

	l, r := 0, n-1
	for l < r && k >= r-l {
		k -= r - l
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	if k > 0 && l < r {
		b[l], b[l+k] = b[l+k], b[l]
		k = 0
	}

	Fprintln(out, k0-k)
	for _, v := range a {
		Fprint(out, v, " ")
	}
	Fprintln(out)
	for _, v := range b {
		Fprint(out, v, " ")
	}
}

//func main() { cf1244G(os.Stdin, os.Stdout) }
