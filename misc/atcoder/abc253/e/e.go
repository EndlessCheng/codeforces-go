package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, m, k int
	Fscan(in, &n, &m, &k)
	s := make([]int, m+1)
	for i := range s {
		s[i] = i
	}
	for ; n > 1; n-- {
		sf := make([]int, m+1)
		for j := 0; j < m; j++ {
			fj := s[m]
			if k > 0 {
				fj -= s[min(j+k, m)] - s[max(j-k+1, 0)]
			}
			sf[j+1] = (sf[j] + fj) % mod
		}
		s = sf
	}
	Fprint(out, (s[m]+mod)%mod)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
