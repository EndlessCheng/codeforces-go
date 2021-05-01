package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF260D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, c int
	Fscan(in, &n)
	s := make([]int, n+1)
	p := [2][]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &c, &s[i])
		p[c] = append(p[c], i)
	}
	for ; n > 1; n-- {
		v, w := p[0][0], p[1][0]
		d := min(s[v], s[w])
		Fprintln(out, v, w, d)
		s[v] -= d
		s[w] -= d
		if s[v] == 0 && len(p[0]) > 1 {
			p[0] = p[0][1:]
		} else {
			p[1] = p[1][1:]
		}
	}
}

//func main() { CF260D(os.Stdin, os.Stdout) }
