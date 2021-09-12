package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
	"time"
)

// github.com/EndlessCheng/codeforces-go
func CF1314D(_r io.Reader, out io.Writer) {
	rand.Seed(time.Now().UnixNano())
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, k int
	Fscan(in, &n, &k)
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			Fscan(in, &dis[i][j])
		}
	}

	ans := int(1e9)
	dp := make([]int, n)
	color := make([]int, n)
	for t := 5000; t > 0; t-- {
		dp[0] = 0
		for i := 1; i < n; i++ {
			dp[i] = 1e9
			color[i] = rand.Intn(2)
		}
		for i := 0; i < k; i++ {
			for v, c := range color {
				if c != i&1 {
					dp[v] = 1e9
					for from, c := range color {
						if c == i&1 {
							// dp[i][v] 表示走了 i 条边，当前在节点 v 时的最小花费，那么答案就是 dp[k][0]
							// 其中第一维可以压缩掉
							dp[v] = min(dp[v], dp[from]+dis[from][v])
						}
					}
				}
			}
		}
		ans = min(ans, dp[0])
	}
	Fprint(out, ans)
}

//func main() { CF1314D(os.Stdin, os.Stdout) }
