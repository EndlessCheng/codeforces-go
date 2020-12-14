package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func findBestValue(a []int, target int) int {
	a = append(a, 0)
	sort.Ints(a)
	sum := 0
	for _, v := range a {
		sum += v
	}
	n := len(a)
	mi, ans := abs(sum-target), a[n-1]
	for i := n - 1; i > 0; i-- {
		for j := a[i] - 1; j >= a[i-1]; j-- {
			sum -= n - i
			if d := abs(sum - target); d <= mi {
				mi, ans = d, j
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
