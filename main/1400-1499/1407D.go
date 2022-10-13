package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1407D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var n, v int
	Fscan(in, &n, &v)
	type pair struct{ v, f int }
	s := []pair{{v, 0}}
	t := []pair{{v, 0}}
	for ; n > 1; n-- {
		Fscan(in, &v)
		f := int(2e9)
		for len(s) > 0 && v > s[len(s)-1].v {
			f = min(f, s[len(s)-1].f)
			s = s[:len(s)-1]
		}
		if len(s) > 0 {
			f = min(f, s[len(s)-1].f)
			if s[len(s)-1].v == v {
				s = s[:len(s)-1]
			}
		}
		for len(t) > 0 && v < t[len(t)-1].v {
			f = min(f, t[len(t)-1].f)
			t = t[:len(t)-1]
		}
		if len(t) > 0 {
			f = min(f, t[len(t)-1].f)
			if t[len(t)-1].v == v {
				t = t[:len(t)-1]
			}
		}
		s = append(s, pair{v, f + 1})
		t = append(t, pair{v, f + 1})
	}
	Fprint(out, t[len(t)-1].f)
}

//func main() { CF1407D(os.Stdin, os.Stdout) }
