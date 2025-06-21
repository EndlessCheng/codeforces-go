package main

import "slices"

// https://space.bilibili.com/206214
func minimumDeletions(word string, k int) int {
	const sigma = 26
	cnt := make([]int, sigma)
	for _, b := range word {
		cnt[b-'a']++
	}
	slices.Sort(cnt)

	var maxSave, s, right int
	for _, base := range cnt {
		for right < sigma && cnt[right] <= base+k {
			s += cnt[right]
			right++
		}
		// 现在 s 表示出现次数不变的字母个数之和
		// 再加上出现次数减少为 base+k 的 len(cnt)-right 种字母，即为保留的字母总数
		maxSave = max(maxSave, s+(base+k)*(sigma-right))
		// 下一轮循环 base 全删
		s -= base
	}
	return len(word) - maxSave
}
