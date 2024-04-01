package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1299A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	suf := make([]int, n+1)
	suf[n] = -1
	for i := n - 1; i > 0; i-- {
		suf[i] = suf[i+1] &^ a[i]
	}

	pre := -1
	mx, mxI := -1, 0
	for i, v := range a {
		and := pre & v & suf[i+1]
		if and > mx {
			mx, mxI = and, i
		}
		pre &^= v
	}

	Fprint(out, a[mxI])
	for i, v := range a {
		if i != mxI {
			Fprint(out, " ", v)
		}
	}
}

//func main() { cf1299A(os.Stdin, os.Stdout) }
