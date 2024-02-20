package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1198B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, p, x int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &q)
	sufMax := make([]int, q)
	changeTime := make([]int, n)
	for t := 0; t < q; t++ {
		Fscan(in, &op)
		if op == 1 {
			Fscan(in, &p, &x)
			p--
			a[p] = x
			changeTime[p] = t
		} else {
			Fscan(in, &x)
			sufMax[t] = x
		}
	}
	for i := q - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i], sufMax[i+1])
	}
	for i, v := range a {
		Fprint(out, max(v, sufMax[changeTime[i]]), " ")
	}
}

//func main() { cf1198B(os.Stdin, os.Stdout) }
