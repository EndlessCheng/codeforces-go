package main

// https://space.bilibili.com/206214
func findChampion(n int, edges [][]int) int {
	weak := make([]bool, n)
	for _, e := range edges {
		weak[e[1]] = true
	}
	ans := -1
	for i, w := range weak {
		if !w {
			if ans != -1 {
				return -1
			}
			ans = i
		}
	}
	return ans
}
