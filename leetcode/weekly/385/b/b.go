package main

import "strconv"

// https://space.bilibili.com/206214
func longestCommonPrefix(arr1, arr2 []int) int {
	has := map[int]bool{}
	for _, v := range arr1 {
		for ; v > 0; v /= 10 {
			has[v] = true
		}
	}

	mx := 0
	for _, v := range arr2 {
		for ; v > 0 && !has[v]; v /= 10 {
		}
		mx = max(mx, v)
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
