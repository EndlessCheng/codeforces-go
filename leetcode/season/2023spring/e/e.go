package main

import (
	"strings"
)

// https://space.bilibili.com/206214
// 第二个维度表示棋子颜色：0 表示末尾添加 B，1 表示末尾添加 R
var trans = [7][2]int{
	{1, 2},  // 空
	{3, 6},  // 只有一个 B
	{5, 4},  // 只有一个 R
	{3, -1}, // 连续多个 B
	{-1, 4}, // 连续多个 R
	{-1, 6}, // BR 交替，且以 B 结尾
	{5, -1}, // BR 交替，且以 R 结尾
}

func getSchemeCount(_,_ int, g []string) int64 {
	n, m := len(g), len(g[0])
	a := make([][]byte, n)
	for i, row := range g {
		a[i] = []byte(row)
	}
	if n < m {
		a = rotate(a) // 保证 n >= m
		n, m = m, n
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<(m*3))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var DFS func(int, int) int
	DFS = func(r, mask int) int {
		if r == n { // 找到 1 个合法方案
			return 1
		}
		ptr := &memo[r][mask]
		if *ptr != -1 {
			return *ptr
		}

		row := a[r]
		// 搜索这一行的合法转移
		var dfs func(int, int, int) int
		dfs = func(c, rowMask, colMask int) (res int) {
			if c == m {
				return DFS(r+1, colMask) // 枚举下一行
			}
			// 计算这一行的下一个状态
			next := func(color int) int {
				rm := trans[rowMask][color] // 新的 rowMask
				if rm < 0 { // 非法
					return 0
				}
				c3 := c * 3
				cm := trans[colMask>>c3&7][color] // 新的 colMask 的第 c 列
				if cm < 0 { // 非法
					return 0
				}
				return dfs(c+1, rm, colMask&^(7<<c3)|cm<<c3) // 修改 colMask 的第 c 列
			}
			switch row[c] {
			case 'B': // 填 B
				return next(0)
			case 'R': // 填 R
				return next(1)
			case '?': // 留空 / 填 B / 填 R
				return dfs(c+1, rowMask, colMask) + next(0) + next(1)
			default: // 留空
				return dfs(c+1, rowMask, colMask)
			}
		}
		*ptr = dfs(0, 0, mask)
		return *ptr
	}
	return int64(DFS(0, 0))
}

func rotate(a [][]byte) [][]byte {
	n, m := len(a), len(a[0])
	b := make([][]byte, m)
	for i := range b {
		b[i] = make([]byte, n)
	}
	for i, r := range a {
		for j, v := range r {
			b[j][n-1-i] = v
		}
	}
	return b
}

func main() {
	a := []string{}
	for i := 0; i < 6; i++ {
		a = append(a, strings.Repeat("?", 5))
	}
	getSchemeCount(0,0,a)
}
