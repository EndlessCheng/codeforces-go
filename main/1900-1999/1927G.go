package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1927G(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n, n+1)
		for i := range a {
			Fscan(in, &a[i])
		}
		a = append(a, 1)

		dp := make([][][]int, n)
		for i := range dp {
			dp[i] = make([][]int, n+1)
			for j := range dp[i] {
				dp[i][j] = make([]int, n+1)
				for k := range dp[i][j] {
					dp[i][j][k] = -1
				}
			}
		}
		var f func(int, int, int) int
		f = func(i, j, doneL int) int {
			if doneL == 0 {
				return 0
			}
			if i < 0 {
				return 1e9
			}
			p := &dp[i][j][doneL]
			if *p != -1 {
				return *p
			}

			// [doneL, n-1] 已完成
			// 如果 j != n，则表示 [j+1, doneL-1] 未完成

			// 不选
			res := f(i-1, j, doneL)

			// 向左
			if j == n {
				if i >= doneL-1 {
					res = min(res, f(i-1, n, min(doneL, max(i-a[i]+1, 0)))+1)
				} else {
					res = min(res, f(i-1, i, doneL)+1) // 不连续
				}
			} else if i >= j-a[j] && i-a[i] < j-a[j] {
				res = min(res, f(i-1, i, doneL)+1) // 仍然不连续，但是最左可以被更新
			}

			// 向右
			if i < doneL && i+a[i] >= doneL {
				res = min(res, f(i-1, n, min(i, max(j-a[j]+1, 0)))+1)
			}

			*p = res
			return res
		}
		Fprintln(out, f(n-1, n, n))
	}
}

//func main() { cf1927G(os.Stdin, os.Stdout) }
