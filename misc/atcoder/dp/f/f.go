package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var s, t []byte
	Fscan(in, &s, &t)
	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i, v := range s {
		for j, w := range t {
			if v == w {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}

	lcs := []byte{}
	i, j := n-1, m-1
	for i >= 0 && j >= 0 {
		if s[i] == t[j] {
			lcs = append(lcs, s[i])
			i--
			j--
		} else if f[i+1][j+1] == f[i][j+1] {
			i--
		} else {
			j--
		}
	}
	for i, j := 0, len(lcs)-1; i < j; i, j = i+1, j-1 {
		lcs[i], lcs[j] = lcs[j], lcs[i]
	}
	Fprintf(out, "%s", lcs)
}

func main() { run(os.Stdin, os.Stdout) }
