package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 另一种思路是正反各扫一遍，求出以 i 为结尾的最长和以 i 为开头的最长，然后枚举要修改的位置，拼接两部分
// https://www.luogu.com.cn/blog/PlutoXz/solution-cf446a

// github.com/EndlessCheng/codeforces-go
func CF446A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	dp := make([][2]int, n)
	dp[0] = [2]int{1, 1}
	ans := 1
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = dp[i-1][1] + 1
		} else {
			dp[i][0] = 1
			dp[i][1] = 2
		}
		if i > 1 && a[i] > a[i-2]+1 {
			dp[i][1] = max(dp[i][1], dp[i-2][0]+2)
		}
		ans = max(ans, max(max(dp[i][0], dp[i][1]), dp[i-1][0]+1))
	}
	Fprint(out, ans)
}

//func main() { CF446A(os.Stdin, os.Stdout) }
