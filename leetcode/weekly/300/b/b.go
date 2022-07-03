package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214/dynamic
var dir4 = []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

func spiralMatrix(n int, m int, head *ListNode) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}
	for x, y, di := 0, 0, 0; head != nil; head = head.Next {
		ans[x][y] = head.Val
		d := dir4[di&3]
		if xx, yy := x+d.x, y+d.y; xx < 0 || xx >= n || yy < 0 || yy >= m || ans[xx][yy] != -1 {
			di++
		}
		d = dir4[di&3]
		x += d.x
		y += d.y
	}
	return ans
}
