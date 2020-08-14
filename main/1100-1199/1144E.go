package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1144E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s, t []byte
	Fscan(in, &n, &s, &t)
	m := make([]byte, n)
	m[n-1] = (s[n-1] + t[n-1]) / 2
	for i := n - 1; i >= 0; i-- {
		m[i] = (s[i] + t[i]) / 2
		if s[i]&1 != t[i]&1 {
			m[i+1] += 13
			if m[i+1] > 'z' {
				m[i+1] -= 26
				m[i]++
			}
		}
	}
	Fprint(out, string(m))
}

//func main() {
//	CF1144E(os.Stdin, os.Stdout)
//}
