package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 思路1：由于求解的是最长的 11...22...11...22...，可以枚举中间位置，分成左右两部分，每部分都为 11...22...，用前缀和搞定
// 这样复杂度是 O(n^2)
// 思路2：题目本质是将数组划分成 4 部分（11... 22... 11... 22...），那么定义 dp[i][j] 表示前 i 个数组成了 j 个部分的最长值
// 遍历一遍即可求出 dp[n][4]，复杂度为 O(n)
// 参考 https://www.luogu.com.cn/problem/solution/CF933A
// 思路2对应题目 https://www.acwing.com/problem/content/3552/

// EXTRA: 数据范围不止 [1,2] http://acm.hdu.edu.cn/showproblem.php?pid=6357

// github.com/EndlessCheng/codeforces-go
func CF933A(_r io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	sum := make([][2]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		sum[i+1][0] = sum[i][0]
		sum[i+1][1] = sum[i][1]
		sum[i+1][a[i]-1]++
	}
	ans := 0
	for i := range a {
		maxL := 0
		for _, s := range sum[:i+1] {
			maxL = max(maxL, s[0]+sum[i][1]-s[1])
		}
		maxR := 0
		for _, s := range sum[i:] {
			maxR = max(maxR, s[0]-sum[i][0]+sum[n][1]-s[1])
		}
		ans = max(ans, maxL+maxR)
	}
	Fprint(out, ans)
}

func CF933A2(_r io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	dp := [5]int{}
	for ; n > 0; n-- {
		if Fscan(in, &v); v == 1 {
			dp[1]++
			dp[2] = max(dp[1], dp[2])
			dp[3] = max(dp[2], dp[3]+1)
			dp[4] = max(dp[3], dp[4])
		} else {
			dp[2] = max(dp[1], dp[2]+1)
			dp[3] = max(dp[2], dp[3])
			dp[4] = max(dp[3], dp[4]+1)
		}
	}
	Fprint(out, dp[4])
}

//func main() { CF933A(os.Stdin, os.Stdout) }
