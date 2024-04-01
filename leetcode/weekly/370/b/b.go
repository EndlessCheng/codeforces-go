package main

// https://space.bilibili.com/206214
func findChampion(n int, edges [][]int) int {
	weak := make([]bool, n)
	for _, e := range edges {
		weak[e[1]] = true
	}

	ans := -1
	for i, w := range weak {
		if w {
			continue
		}
		if ans != -1 {
			return -1 // 冠军只能有一个
		}
		ans = i
	}
	return ans
}
