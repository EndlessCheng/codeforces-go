package main

// https://space.bilibili.com/206214
func matrixSumQueries(n int, queries [][]int) (ans int64) {
	vis := [2]map[int]bool{{}, {}}
	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		tp, index, val := q[0], q[1], q[2]
		if !vis[tp][index] { // 后面（>i）没有对这一行/列的操作
			// 这一行/列剩余 n-len(vis[tp^1]) 个没有被修改的
			ans += int64(n-len(vis[tp^1])) * int64(val)
			vis[tp][index] = true
		}
	}
	return
}
