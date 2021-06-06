package main

import (
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func minWastedSpace(a []int, boxes [][]int) int {
	sort.Ints(a)
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	ans := math.MaxInt64
	for _, box := range boxes {
		sort.Ints(box)
		if box[len(box)-1] < a[len(a)-1] { // 最大的箱子不够装最大的包裹
			continue
		}
		res, l := 0, 0
		for _, v := range box {
			r := sort.SearchInts(a, v+1)
			res += (r-l)*v - (sum[r] - sum[l])
			l = r
		}
		ans = min(ans, res)
	}
	if ans < math.MaxInt64 {
		return ans % (1e9 + 7)
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
