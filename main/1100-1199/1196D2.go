package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1196D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	mp := []byte{'R': 0, 'G': 1, 'B': 2}

	var T, n, k int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		ans := n
		dp := make([][3]int, n+1)
		for i, b := range s {
			dp[i+1] = [3]int{dp[i][2] + 1, dp[i][0] + 1, dp[i][1] + 1}
			dp[i+1][mp[b]]--
			if i+1 >= k {
				for j, dv := range dp[i+1] {
					ans = min(ans, dv-dp[i+1-k][(j+3-k%3)%3])
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1196D2(os.Stdin, os.Stdout) }
