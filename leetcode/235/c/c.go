package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minAbsoluteSumDiff(x []int, y []int) (ans int) {
	sum := 0
	for i, v := range x {
		sum += abs(v - y[i])
	}
	a := append([]int(nil), x...)
	sort.Ints(a)
	ans = sum
	for j, v := range x {
		w := y[j]
		old := abs(v - w)
		i := sort.SearchInts(a, w)
		if i < len(a) {
			ans = min(ans, sum-old+a[i]-w)
		}
		if i > 0 {
			ans = min(ans, sum-old+w-a[i-1])
		}
	}
	return ans % (1e9 + 7)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
