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
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	// 把连续相同颜色的砖块合并成一个砖块，这样在下面计算转移时，保证转移来源区间的颜色与当前区间端点的颜色是不同的
	k := 0
	for _, w := range a[1:] {
		if a[k] != w {
			k++
			a[k] = w
		}
	}

	// dp[l][r] 表示使区间 [l,r] 颜色相同的最小操作次数
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for sz := 1; sz <= k; sz++ {
		for l, r := 0, sz; r <= k; r++ {
			if a[l] == a[r] { // 如果 l r 颜色相同就可以从区间 [l+1,r-1] 转移，涂成和 l 相同的颜色
				dp[l][r] = dp[l+1][r-1] + 1
			} else { // 如果 l r 颜色不同就做决策，是涂成 l 的颜色还是涂成 r 的颜色，取最小值
				dp[l][r] = min(dp[l+1][r], dp[l][r-1]) + 1
			}
			l++
		}
	}
	Fprint(out, dp[0][k])
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
