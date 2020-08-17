package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1368E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, m, v, w int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		g := make([][]int, n+1)
		for ; m > 0; m-- {
			Fscan(in, &v, &w)
			g[w] = append(g[w], v)
		}

		ans := []interface{}{}
		dp := make([]int, n+1)
		for w, vs := range g {
			for _, v := range vs {
				if dp[v]+1 > dp[w] {
					dp[w] = dp[v] + 1
				}
			}
			if dp[w] > 1 {
				dp[w] = -1
				ans = append(ans, w)
			}
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { CF1368E(os.Stdin, os.Stdout) }
