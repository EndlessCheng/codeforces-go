package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF1804D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m, mn, mx int
	var s string
	for Fscan(in, &n, &m); n > 0; n-- {
		Fscan(in, &s)
		c1 := strings.Count(s, "1")
		c := 0
		for _, t := range strings.FieldsFunc(s, func(r rune) bool { return r == '0' }) {
			c += len(t) / 2
		}
		mn += c1 - min(c, m/4)

		c = m / 4
		for i := 1; i < m && c > 0; i++ {
			if s[i] == '0' || s[i-1] == '0' {
				c--
				i++
			}
		}
		mx += c1 - c
	}
	Fprint(out, mn, mx)
}

//func main() { CF1804D(os.Stdin, os.Stdout) }
