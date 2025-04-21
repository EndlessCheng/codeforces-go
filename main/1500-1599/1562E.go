package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1562E(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		lcp := make([][]int, n+1)
		for i := range lcp {
			lcp[i] = make([]int, n+1)
		}
		for i := n - 1; i >= 0; i-- {
			for j := i - 1; j >= 0; j-- {
				if s[i] == s[j] {
					lcp[i][j] = lcp[i+1][j+1] + 1
				}
			}
		}

		f := make([]int, n)
		for i, row := range lcp[:n] {
			f[i] = n - i
			for j, l := range row[:i] {
				if l < n-i && s[j+l] < s[i+l] {
					f[i] = max(f[i], f[j]+n-i-lcp[i][j])
				}
			}
		}
		Fprintln(out, slices.Max(f))
	}
}

//func main() { cf1562E(bufio.NewReader(os.Stdin), os.Stdout) }
