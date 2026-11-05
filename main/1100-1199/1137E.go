package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1137E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, x, y, op, s, k, t int
	Fscan(in, &n, &q)
	a := make([]int, q+2)
	b := make([]int, q+2)
	p := func(t int) int {
		return b[t] + s + a[t]*k
	}
	ceil := func(x, y int) int {
		return (x + y - 1) / y
	}
	for range q {
		Fscan(in, &op, &x)
		if op == 1 {
			n += x
			s, k, t = 0, 0, 0
		} else if op == 2 {
			if p(t) > 0 {
				for t > 0 && ceil(p(t-1)-p(t), a[t]-a[t-1]) < ceil(p(t), n-a[t]) {
					t--
				}
				t++
				a[t] = n
				b[t] = -s - n*k
			}
			n += x
		} else if op == 3 {
			Fscan(in, &y)
			s += x
			k += y
			for t > 0 && p(t-1) <= p(t) {
				t--
			}
		}
		Fprintln(out, a[t]+1, p(t))
	}
}

//func main() { cf1137E(bufio.NewReader(os.Stdin), os.Stdout) }
