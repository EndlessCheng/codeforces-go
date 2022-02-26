package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF817D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := int64(0)
	s := make([]int, 0, n)
	r := make([]int, n)
	f := func() {
		s = s[:0]
		for i := n - 1; i >= 0; i-- {
			v := a[i]
			for len(s) > 0 && a[s[len(s)-1]] < v {
				s = s[:len(s)-1]
			}
			if len(s) > 0 {
				r[i] = s[len(s)-1]
			} else {
				r[i] = n
			}
			s = append(s, i)
		}

		s = s[:0]
		for i, v := range a {
			for len(s) > 0 && a[s[len(s)-1]] <= v {
				s = s[:len(s)-1]
			}
			l := -1
			if len(s) > 0 {
				l = s[len(s)-1]
			}
			ans += int64(i-l) * int64(r[i]-i) * int64(v)
			s = append(s, i)
		}
	}
	f()
	for i, v := range a {
		a[i] = -v
	}
	f()
	Fprint(out, ans)
}

//func main() { CF817D(os.Stdin, os.Stdout) }
