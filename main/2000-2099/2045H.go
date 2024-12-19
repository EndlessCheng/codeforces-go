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
		for j := n - 1; j >= 0; j-- {
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
	type pair struct{ f, k int16 }
	sufMax := make([][]pair, n)
	next := make([][]int16, n)
	for i := range f {
		f[i] = make([]int16, n)
		sufMax[i] = make([]pair, n)
		next[i] = make([]int16, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][n-1] = 1
		next[i][n-1] = int16(n)
		p := pair{1, int16(n - 1)}
		for j := n - 1; j >= i; j-- {
			if less(i, j+1, j+1, n) {
				l := min(int(lcp[i][j+1]), j-i+1)
				f[i][j] = sufMax[j+1][j+1+l].f + 1
				next[i][j] = sufMax[j+1][j+1+l].k
				if f[i][j] > p.f {
					p = pair{f[i][j], int16(j)}
				}
			}
			sufMax[i][j] = p
		}
	}
	Fprintln(out, sufMax[0][0].f)
	for i, j := 0, int(sufMax[0][0].k); i < n; i, j = j+1, int(next[i][j]) {
		Fprintln(out, s[i:j+1])
	}
}

//func main() { cf2045H(bufio.NewReader(os.Stdin), os.Stdout) }
