package main

// https://space.bilibili.com/206214/dynamic
func shortestSequence(rolls []int, k int) int {
	mark := make([]int, k+1) // mark[v] 标记 v 属于哪个子段
	ans, left := 1, k
	for _, v := range rolls {
		if mark[v] < ans {
			mark[v] = ans
			if left--; left == 0 {
				left = k
				ans++
			}
		}
	}
	return ans
}
