package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1213E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s, t string
	Fscan(bufio.NewReader(in), &n, &s, &t)
	Fprintln(out, "YES")
	if s[0] == s[1] && t[0] == t[1] {
		Fprint(out, strings.Repeat("abc", n))
	} else if s[0] != s[1] && t[0] != t[1] {
		for _, p := range []string{"abc", "acb", "bac", "bca", "cab", "cba"} {
			if p[:2] != s && p[:2] != t && p[1:] != s && p[1:] != t {
				for i := range p {
					Fprint(out, strings.Repeat(p[i:i+1], n))
				}
				break
			}
		}
	} else {
		if t[0] == t[1] {
			s, t = t, s
		}
		if t[0] != s[0] && t[1] != s[0] {
			Fprint(out, strings.Repeat(s[:1]+t[1:]+t[:1], n))
		} else {
			miss := byte('a')
			for ; miss == s[0] || miss == t[0] || miss == t[1]; miss++ {
			}
			b := t[0]
			if b == s[0] {
				b = t[1]
			}
			Fprint(out, strings.Repeat(s[:1]+string(miss), n), strings.Repeat(string(b), n))
		}
	}
}

//func main() { CF1213E(os.Stdin, os.Stdout) }
