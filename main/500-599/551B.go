package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF551B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	r := func(s []byte) (c [26]int) {
		for _, b := range s {
			c[b-'a']++
		}
		return
	}

	var S, X, Y []byte
	Fscan(in, &S, &X, &Y)
	s, x, y := r(S), r(X), r(Y)
	rx, ry := 0, 0
o:
	for i := 0; ; i++ {
		mi := int(1e9)
		for j, c := range y[:] {
			if s[j] < x[j]*i {
				break o
			}
			if c > 0 {
				mi = min(mi, (s[j]-x[j]*i)/c)
			}
		}
		if i+mi > rx+ry {
			rx, ry = i, mi
		}
	}
	ans := append(bytes.Repeat(X, rx), bytes.Repeat(Y, ry)...)
	for i, c := range x[:] {
		ans = append(ans, bytes.Repeat([]byte{'a' + byte(i)}, s[i]-c*rx-y[i]*ry)...)
	}
	Fprint(out, string(ans))
}

//func main() { CF551B(os.Stdin, os.Stdout) }
