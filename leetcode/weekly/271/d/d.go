package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxTotalFruits(fruits [][]int, startPos int, k int) (ans int) {
	n := len(fruits)
	sum := make([]int, n+1)
	for i, f := range fruits {
		sum[i+1] = sum[i] + f[1]
	}

	// 先向右再向左
	mid := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos })
	for i := mid; i < n; i++ {
		d := fruits[i][0] - startPos
		if d > k {
			break
		}
		cnt := sum[i+1] - sum[mid]
		d *= 2
		if d < k {
			// 往左最远能到达的水果下标
			left := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos-k+d })
			cnt += sum[mid] - sum[left]
		}
		ans = max(ans, cnt)
	}

	// 先向左再向右
	for i := mid - 1; i >= 0; i-- {
		d := startPos - fruits[i][0]
		if d > k {
			break
		}
		cnt := sum[mid] - sum[i]
		d *= 2
		if d < k {
			// 往右最远能到达的水果下标+1
			right := sort.Search(n, func(i int) bool { return fruits[i][0] > startPos+k-d })
			cnt += sum[right] - sum[mid]
		}
		ans = max(ans, cnt)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
