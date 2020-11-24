package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ v, i int }

	var n, m int
	Fscan(in, &n, &m)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })

	sum := make([]int, n+1)
	for i, p := range a {
		sum[i+1] = sum[i] + p.v
	}

	dp := make([][]int, n+1)
	fa := make([][]pair, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = 1e9
		}
		fa[i] = make([]pair, m+1)
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := i; j <= m; j++ {
			dp[i][j] = dp[i][j-i] // 直接转移
			fa[i][j] = pair{i, j - i}
			for k := 0; k < i; k++ {
				if s := dp[k][j-(i-k)] + (sum[i]-sum[k])*k; s < dp[i][j] { // k~i 都分 1 块，则前面有前 k 个孩子分得的数量比这些孩子多
					dp[i][j] = s
					fa[i][j] = pair{k, j - (i - k)}
				}
			}
		}
	}
	Fprintln(out, dp[n][m])

	ans := make([]int, n)
	var pt func(n, m int)
	pt = func(n, m int) {
		if n == 0 {
			return
		}
		x, y := fa[n][m].v, fa[n][m].i
		pt(x, y)
		if x == n { // 最后一行的直接转移
			for _, p := range a {
				ans[p.i]++
			}
		} else {
			for _, p := range a[x:] {
				ans[p.i] = 1
			}
		}
	}
	pt(n, m)
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
