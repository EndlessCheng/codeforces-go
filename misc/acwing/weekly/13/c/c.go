package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, cnt, ans int
	var s []byte
	Fscan(in, &n, &m, &s)
	g := make([][]int, n)
	deg := make([]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		deg[w]++
	}

	dp := make([][26]int, n)
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		cnt++
		dp[v][s[v]-'a']++
		ans = max(ans, dp[v][s[v]-'a'])
		for _, w := range g[v] {
			for i, dv := range dp[v] {
				dp[w][i] = max(dp[w][i], dv)
			}
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}
	if cnt < n {
		ans = -1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
