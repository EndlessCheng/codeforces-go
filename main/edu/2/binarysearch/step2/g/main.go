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
	var fill, n int64
	Fscan(in, &fill, &n)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
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
	Fprint(out, search(1, n*1e9+1, func(m int64) bool {
		c := int64(0)
		for _, v := range a {
			if v > m {
				v = m
			}
			c += v
		}
		return c < m*fill
	})-1)
}

func main() { run(os.Stdin, os.Stdout) }
