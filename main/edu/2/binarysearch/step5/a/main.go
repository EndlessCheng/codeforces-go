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
	type pair struct{ l, r int64 }

	var n, k int64
	Fscan(in, &n, &k)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
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
	Fprint(out, search(-2e9, 2e9, func(x int64) bool {
		var c int64
		for _, p := range a {
			if x >= p.l {
				if x < p.r {
					c += x - p.l + 1
				} else {
					c += p.r - p.l + 1
				}
			}
		}
		return c >= k+1
	}))
}

func main() { run(os.Stdin, os.Stdout) }
