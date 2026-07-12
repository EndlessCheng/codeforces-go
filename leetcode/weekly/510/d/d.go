package main

// https://space.bilibili.com/206214
func maxConsistentColumns(grid [][]int, limit int) (ans int) {
	n := len(grid[0])
	f := make([]int, n)
	for i := range n {
	next:
		for j := i - 1; j >= 0; j-- { // 枚举上一个保留的列
			if f[j] <= f[i] {
				continue
			}
			for _, row := range grid {
				if abs(row[i]-row[j]) > limit {
					continue next // 列 i 和列 j 不是一致的
				}
			}
			f[i] = f[j]
		}
		f[i]++
		ans = max(ans, f[i])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
