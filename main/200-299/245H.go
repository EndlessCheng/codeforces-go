package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF245H(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	var q, l, r int
	Fscan(in, &s, &q)
	n := len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		f[i][i] = 1
		if i+1 < n && s[i] == s[i+1] {
			f[i][i+1] = 1
		}
	}
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			if s[i] == s[j] {
				f[i][j] = f[i+1][j-1]
			}
		}
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] += f[i][j-1] + f[i+1][j] - f[i+1][j-1]
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		Fprintln(out, f[l-1][r-1])
	}
}

//func main() { CF245H(os.Stdin, os.Stdout) }
