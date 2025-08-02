package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxTotalFruits1(fruits [][]int, startPos, k int) int {
	n := len(fruits)
	// 向左最远能到 fruits[left][0]
	left := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos-k })

	right, s := left, 0
	// 从 fruits[left][0] 到 startPos 的水果数
	for ; right < n && fruits[right][0] <= startPos; right++ {
		s += fruits[right][1]
	}

	ans := s
	// 枚举最右走到 fruits[right][0]
	for ; right < n && fruits[right][0] <= startPos+k; right++ {
		s += fruits[right][1] // 枚举最右位置为 fruits[right][0]
		for fruits[right][0]*2-fruits[left][0]-startPos > k &&
			fruits[right][0]-fruits[left][0]*2+startPos > k {
			s -= fruits[left][1] // fruits[left][0] 太远了
			left++
		}
		ans = max(ans, s) // 更新答案最大值
	}
	return ans
}

func maxTotalFruits(fruits [][]int, startPos, k int) (ans int) {
	n := len(fruits)
	// 向左最远能到 fruits[left][0]
	left := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos-k })

	s := 0
	// 枚举最右走到 fruits[right][0]
	for right := left; right < n && fruits[right][0] <= startPos+k; right++ {
		s += fruits[right][1]
		for fruits[right][0]*2-fruits[left][0]-startPos > k &&
			fruits[right][0]-fruits[left][0]*2+startPos > k {
			s -= fruits[left][1] // fruits[left][0] 太远了
			left++
		}
		ans = max(ans, s) // 更新答案最大值
	}
	return
}
