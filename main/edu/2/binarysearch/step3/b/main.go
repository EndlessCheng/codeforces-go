package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, l int64
	Fscan(in, &n, &k)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] > l {
			l = a[i]
		}
	}
	search := func(l, r int64, f func(int64) bool) int64 {
		for l < r {
			m := (l + r) >> 1
			if f(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}
	Fprint(out, search(l, n*1e9, func(sum int64) bool {
		c, s := int64(1), int64(0)
		for _, v := range a {
			if s+v <= sum {
				s += v
			} else {
				s = v
				c++
			}
		}
		return c <= k
	}))
}

func main() { run(os.Stdin, os.Stdout) }
