package main

import "cmp"

// https://space.bilibili.com/206214
func countMatchingSubarrays(nums, pattern []int) (ans int) {
	m := len(pattern)
	pattern = append(pattern, 2)
	for i := 1; i < len(nums); i++ {
		pattern = append(pattern, cmp.Compare(nums[i], nums[i-1]))
	}

	n := len(pattern)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && pattern[z[i]] == pattern[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}

	for _, lcp := range z[m+1:] {
		if lcp == m {
			ans++
		}
	}
	return
}

func countMatchingSubarrays2(nums, pattern []int) (ans int) {
	m := len(pattern)
	pi := make([]int, m)
	cnt := 0
	for i := 1; i < m; i++ {
		v := pattern[i]
		for cnt > 0 && pattern[cnt] != v {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == v {
			cnt++
		}
		pi[i] = cnt
	}

	cnt = 0
	for i := 1; i < len(nums); i++ {
		v := cmp.Compare(nums[i], nums[i-1])
		for cnt > 0 && pattern[cnt] != v {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == v {
			cnt++
		}
		if cnt == m {
			ans++
			cnt = pi[cnt-1]
		}
	}
	return
}
