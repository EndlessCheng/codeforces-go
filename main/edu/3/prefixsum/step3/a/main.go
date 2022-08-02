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

	var n, m, v, q, r1, c1, r2, c2 int
	Fscan(in, &n, &m)
	s := make([][]int64, n+1)
	s[0] = make([]int64, m+1)
	for i := 0; i < n; i++ {
		s[i+1] = make([]int64, m+1)
		for j := 0; j < m; j++ {
			Fscan(in, &v)
			s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j] + int64(v)
		}
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &r1, &c1, &r2, &c2)
		r1--
		c1--
		Fprintln(out, s[r2][c2]-s[r2][c1]-s[r1][c2]+s[r1][c1])
	}
}

func main() { run(os.Stdin, os.Stdout) }
