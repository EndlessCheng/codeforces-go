package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1548A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, m, v, w, q, tp, ans int
	Fscan(in, &n, &m)
	c := make([]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		c[min(v, w)]++
	}
	for _, v := range c[1:] {
		if v == 0 {
			ans++
		}
	}
	for Fscan(in, &q); q > 0; q-- {
		if Fscan(in, &tp); tp < 3 {
			Fscan(in, &v, &w)
			v = min(v, w)
			if tp == 1 {
				if c[v] == 0 {
					ans--
				}
				c[v]++
			} else {
				c[v]--
				if c[v] == 0 {
					ans++
				}
			}
		} else {
			Fprintln(out, ans)
		}
	}
}

//func main() { CF1548A(os.Stdin, os.Stdout) }
