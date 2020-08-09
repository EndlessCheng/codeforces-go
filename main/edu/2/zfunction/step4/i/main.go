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
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	reverse := func(a []byte) []byte {
		n := len(a)
		r := make([]byte, n)
		for i, v := range a {
			r[n-1-i] = v
		}
		return r
	}

	var t, p []byte
	var k int
	Fscan(in, &t, &p, &k)
	n, m := len(t), len(p)
	if n < m {
		Fprint(out, 0)
		return
	}
	f := func(s []byte) []int {
		n := len(s)
		z := make([]int, n)
		for i, l, r := 1, 0, 0; i < n; i++ {
			z[i] = max(0, min(z[i-l], r-i+1))
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				l, r = i, i+z[i]
				z[i]++
			}
		}
		return z[m+1:]
	}
	z := f(append(append(p, '#'), t...))
	zr := f(append(append(reverse(p), '#'), reverse(t)...))
	ans := []interface{}{}
	for i, l := range z[:n-m+1] {
		if l+k+zr[n-m-i] >= m {
			ans = append(ans, i+1)
		}
	}
	Fprintln(out, len(ans))
	Fprint(out, ans...)
}

func main() { run(os.Stdin, os.Stdout) }
