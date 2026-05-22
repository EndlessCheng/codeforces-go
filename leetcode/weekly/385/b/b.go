package main

import "strconv"

// https://space.bilibili.com/206214
func longestCommonPrefix(arr1, arr2 []int) int {
	has := map[int]bool{}
	for _, x := range arr1 {
		for x > 0 && !has[x] { // 如果 x 在 st 中，那么剩余前缀也在 st 中
			has[x] = true
			x /= 10
		}
	}

	mx := 0
	for _, x := range arr2 {
		for x > 0 && !has[x] {
			x /= 10
		}
		mx = max(mx, x)
	}

	if mx == 0 {
		return 0
	}
	return len(strconv.Itoa(mx))
}

func longestCommonPrefix2(arr1, arr2 []int) (ans int) {
	has := map[string]bool{}
	for _, x := range arr1 {
		s := strconv.Itoa(x)
		for i := 1; i <= len(s); i++ {
			has[s[:i]] = true
		}
	}

	for _, x := range arr2 {
		s := strconv.Itoa(x)
		for i := 1; i <= len(s) && has[s[:i]]; i++ {
			ans = max(ans, i)
		}
	}
	return
}
