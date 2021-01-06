package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1471B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		s := int64(0)
		Fscan(in, &n, &x)
		a := make([]int, n)
		c := make([]int8, n)
		mi := 0
		for i := range a {
			Fscan(in, &a[i])
			v := a[i]
			for s += int64(v); v%x == 0; v /= x {
				c[i]++
			}
			if c[i] < c[mi] {
				mi = i
			}
		}
		s *= int64(c[mi] + 1)
		for _, v := range a[:mi] {
			s += int64(v)
		}
		Fprintln(out, s)
	}
}

//func main() { CF1471B(os.Stdin, os.Stdout) }
