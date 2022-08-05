package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1716C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &m)
		a := make([][2]int, m)
		for j := 0; j < 2; j++ {
			for i := 0; i < m; i++ {
				Fscan(in, &a[i][j])
			}
		}
		// 我的评价是好想不好写的题
		// 为了方便描述，下面的坐标以 (0,0) 为左上角
		// 小技巧：把起点当成 (0,-1)，这样的话需要把 a[0][0] 置为 -1 从而与原问题保持一致
		a[0][0] = -1
		// suf[i][0] 表示能够「一刻不停地」从 (0,i) 走个反 C 型到 (1,i)，在进入 (0,i) 前的最小时间
		// suf[i][1] 表示能够「一刻不停地」从 (1,i) 走个反 C 型到 (0,i)，在进入 (1,i) 前的最小时间
		suf := make([][2]int, m+1)
		for i := m - 1; i >= 0; i-- {
			suf[i][0] = max(suf[i+1][0]-1, max(a[i][0], a[i][1]-(m-i)*2+1))
			suf[i][1] = max(suf[i+1][1]-1, max(a[i][1], a[i][0]-(m-i)*2+1))
		}
		// i 列及其前的格子走蛇形，i+1 列及其后的格子走反 C 型
		// 模拟即可
		ans, t := suf[0][0]+m*2, -1
		for i, col := range a[:m-1] {
			t = max(t, col[i&1]) + 1
			t = max(t, col[i&1^1]) + 1
			ans = min(ans, max(t, suf[i+1][i&1^1])+(m-1-i)*2)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1716C(os.Stdin, os.Stdout) }
