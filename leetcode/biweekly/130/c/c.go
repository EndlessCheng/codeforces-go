package main

import "math"

// https://space.bilibili.com/206214
func minimumSubstringsInPartition(s string) int {
	n := len(s)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1 // -1 表示没有计算过
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := math.MaxInt
		cnt := [26]int{}
		k := 0
	next:
		for j := i; j >= 0; j-- {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				k++
			}
			cnt[b]++
			if (i-j+1)%k > 0 {
				continue
			}
			for _, c := range cnt {
				if c > 0 && c != cnt[b] {
					continue next
				}
			}
			res = min(res, dfs(j-1)+1)
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n - 1)
}

func minimumSubstringsInPartition2(s string) int {
	n := len(s)
	f := make([]int, n+1)
	for i := range s {
		f[i+1] = math.MaxInt
		cnt := [26]int{}
		k := 0
	next:
		for j := i; j >= 0; j-- {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				k++
			}
			cnt[b]++
			if (i-j+1)%k > 0 {
				continue
			}
			for _, c := range cnt {
				if c != 0 && c != cnt[b] {
					continue next
				}
			}
			f[i+1] = min(f[i+1], f[j]+1)
		}
	}
	return f[n]
}