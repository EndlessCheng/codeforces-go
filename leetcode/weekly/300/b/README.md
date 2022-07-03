下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

代码同 [59. 螺旋矩阵 II](https://leetcode.cn/problems/spiral-matrix-ii/)，改为在链表上遍历即可。

```go
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
```
