package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1633D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	d := [1001]int{}
	for i := 2; i < 1001; i++ {
		d[i] = 1e9
	}
	for i := 1; i < 1000; i++ {
		for x := 1; x <= i; x++ {
			if i+i/x < 1001 {
				d[i+i/x] = min(d[i+i/x], d[i]+1)
			}
		}
	}

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		m := 0
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
			b[i] = d[b[i]]
			m += b[i]
		}
		m = min(m, k)
		c := make([]int, n)
		for i := range c {
			Fscan(in, &c[i])
		}
		f := make([]int, m+1)
		for i, v := range c {
			w := b[i]
			for j := m; j >= w; j-- {
				f[j] = max(f[j], f[j-w]+v)
			}
		}
		Fprintln(out, f[m])
	}
}

//func main() { CF1633D(os.Stdin, os.Stdout) }
