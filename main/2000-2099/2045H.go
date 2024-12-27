package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2045H(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	n := len(s)
	lcp := make([][]int16, n+1)
	for i := range lcp {
		lcp[i] = make([]int16, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if s[i] == s[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}
	less := func(l1, r1, l2, r2 int) bool {
		len1, len2 := r1-l1, r2-l2
		l := int(lcp[l1][l2])
		if l >= min(len1, len2) {
			return len1 < len2
		}
		return s[l1+l] < s[l2+l]
	}

	f := make([][]int16, n)
	sufMaxK := make([][]int16, n)
	next := make([][]int16, n)
	for i := range f {
		f[i] = make([]int16, n)
		sufMaxK[i] = make([]int16, n)
		next[i] = make([]int16, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][n-1] = 1
		next[i][n-1] = int16(n)
		k := int16(n - 1)
		for j := n - 1; j >= i; j-- {
			if less(i, j+1, j+1, n) {
				l := min(int(lcp[i][j+1]), j-i+1)
				k2 := sufMaxK[j+1][j+1+l]
				next[i][j] = k2
				f[i][j] = f[j+1][k2] + 1
				if f[i][j] > f[i][k] {
					k = int16(j)
				}
			}
			sufMaxK[i][j] = k
		}
	}
	Fprintln(out, f[0][sufMaxK[0][0]])
	for i, j := 0, int(sufMaxK[0][0]); i < n; i, j = j+1, int(next[i][j]) {
		Fprintln(out, s[i:j+1])
	}
}

//func main() { cf2045H(bufio.NewReader(os.Stdin), os.Stdout) }
