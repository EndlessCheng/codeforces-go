package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1562E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// 假了，这方法会 T

	var T, n int16
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
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
		lessEq := func(l1, r1, l2, r2 int16) bool {
			len1, len2 := r1-l1, r2-l2
			l := lcp[l1][l2]
			if len1 == len2 && l >= len1 {
				return true
			}
			if l >= len1 || l >= len2 {
				return len1 < len2
			}
			return s[l1+l] < s[l2+l]
		}

		type pair struct{ l, r int16 }
		dp := []pair{}
		for i := int16(0); i < n; i++ {
			for j := i + 1; j <= n; j++ {
				p := sort.Search(len(dp), func(p int) bool { return lessEq(i, j, dp[p].l, dp[p].r) })
				if p < len(dp) {
					dp[p] = pair{i, j}
				} else {
					dp = append(dp, pair{i, j})
				}
			}
		}
		Fprintln(out, len(dp))
	}
}

//func main() { CF1562E(os.Stdin, os.Stdout) }
