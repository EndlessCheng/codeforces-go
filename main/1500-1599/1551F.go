package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1551F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	var T, n, k, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		if k == 2 {
			Fprintln(out, n*(n-1)/2)
			continue
		}

		var cd []int
		var f func(v, fa, d int)
		f = func(v, fa, d int) {
			cd[d]++
			for _, w := range g[v] {
				if w != fa {
					f(w, v, d+1)
				}
			}
		}

		ans := int64(0)
		for rt, ws := range g {
			if len(ws) < k {
				continue
			}
			dp := make([][]int64, n)
			for i := range dp {
				dp[i] = make([]int64, k+1)
				dp[i][0] = 1
			}
			for _, w := range ws {
				cd = make([]int, n)
				f(w, rt, 0)
				for i, d := range cd {
					if d == 0 {
						break
					}
					for j := k; j > 0; j-- {
						dp[i][j] = (dp[i][j] + dp[i][j-1]*int64(d)) % mod
					}
				}
			}
			for _, dr := range dp {
				ans += dr[k]
			}
		}
		Fprintln(out, ans%mod)
	}
}

//func main() { CF1551F(os.Stdin, os.Stdout) }
