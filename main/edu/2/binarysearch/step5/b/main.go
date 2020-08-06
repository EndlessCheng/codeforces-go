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
	var n, k int64
	Fscan(in, &n, &k)
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
	Fprint(out, search(1, n*n, func(x int64) bool {
		c := int64(0)
		for i := int64(1); i <= n; i++ {
			if x < i*n {
				c += x / i
			} else {
				c += n
			}
		}
		return c >= k
	}))
}

func main() { run(os.Stdin, os.Stdout) }
