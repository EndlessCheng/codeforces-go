package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	kmp := func(text, s []int) (pos []int) {
		f := make([]int, n)
		c := 0
		for i := 1; i < n; i++ {
			b := s[i]
			for c > 0 && s[c] != b {
				c = f[c-1]
			}
			if s[c] == b {
				c++
			}
			f[i] = c
		}
		c = 0
		for i, b := range text {
			for c > 0 && s[c] != b {
				c = f[c-1]
			}
			if s[c] == b {
				c++
			}
			if c == n && i-n+1 < n {
				pos = append(pos, i-n+1)
				c = f[c-1]
			}
		}
		return
	}

	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	c := make([]int, n)
	for i, v := range a {
		c[i] = v ^ a[(i+1)%n]
	}
	c = append(c, c...)
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	d := make([]int, n)
	for i, v := range b {
		d[i] = v ^ b[(i+1)%n]
	}
	for _, p := range kmp(c, d) {
		Fprintln(out, p, a[p]^b[0])
	}
}

func main() { run(os.Stdin, os.Stdout) }
