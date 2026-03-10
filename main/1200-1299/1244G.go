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
	p := make([]int, n)
	q := make([]int, n)
	for i := range p {
		p[i] = i + 1
		q[i] = i + 1
	}

	d := k - n*(n+1)/2
	if d < 0 {
		Fprintln(out, -1)
		return
	}

	l, r := 0, n-1
	for l < r && d >= r-l {
		d -= r - l
		q[l], q[r] = q[r], q[l]
		l++
		r--
	}
	if l < r && d > 0 {
		q[l], q[l+d] = q[l+d], q[l]
		d = 0
	}

	Fprintln(out, k-d)
	for _, v := range p {
		Fprint(out, v, " ")
	}
	Fprintln(out)
	for _, v := range q {
		Fprint(out, v, " ")
	}
}

//func main() { cf1244G(os.Stdin, os.Stdout) }
