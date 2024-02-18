package main

import "strconv"

// https://space.bilibili.com/206214
func longestCommonPrefix(arr1, arr2 []int) (ans int) {
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

func longestCommonPrefix2(arr1, arr2 []int) (ans int) {
	has := map[int]bool{}
	a := []int{}
	for _, v := range arr1 {
		a = a[:0]
		for x := v; x > 0; x /= 10 {
			a = append(a, x%10)
		}
		v = 0
		for i := len(a) - 1; i >= 0; i-- {
			v = v*10 + a[i]
			has[v] = true
		}
	}

	for _, v := range arr2 {
		a = a[:0]
		for x := v; x > 0; x /= 10 {
			a = append(a, x%10)
		}
		v = 0
		for i := len(a) - 1; i >= 0; i-- {
			v = v*10 + a[i]
			if !has[v] {
				break
			}
			ans = max(ans, len(a)-i)
		}
	}
	return
}
