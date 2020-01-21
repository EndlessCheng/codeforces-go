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

// github.com/EndlessCheng/codeforces-go
func CF933A(_r io.Reader, _w io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

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

//func main() {
//	CF933A(os.Stdin, os.Stdout)
//}
