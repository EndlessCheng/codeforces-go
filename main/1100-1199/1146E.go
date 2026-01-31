package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1146E(in io.Reader, _w io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, x int
	var op string
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	const mx int = 1e5
	var time, flip [mx + 2]int
	f := 1
	for i := 1; i <= q; i++ {
		Fscan(in, &op, &x)
		if op == "<" {
			f = -f
			x = -x
		}
		if x >= 0 {
			time[x+1] = i
			flip[x+1] = -f
		} else {
			f = -f
			time[-x] = i
			flip[-x] = -f
		}
		if op == "<" {
			f = -f
		}
	}

	for i := 1; i <= mx; i++ {
		if time[i-1] > time[i] {
			time[i] = time[i-1]
			flip[i] = flip[i-1]
		}
	}

	for _, v := range a {
		w := abs(v)
		if flip[w] != 0 {
			v = w * flip[w]
		}
		Fprint(out, v*f, " ")
	}
}

//func main() { cf1146E(bufio.NewReader(os.Stdin), os.Stdout) }
