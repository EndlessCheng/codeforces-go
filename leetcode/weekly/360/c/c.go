package main

import "math/bits"

// https://space.bilibili.com/206214
func minOperations(nums []int, target int) (ans int) {
	s := 0
	cnt := [31]int{}
	for _, v := range nums {
		s += v
		cnt[bits.TrailingZeros(uint(v))]++
	}
	if s < target {
		return -1
	}
	s = 0
	for i := 0; 1<<i <= target; {
		s += cnt[i] << i
		mask := 1<<(i+1) - 1
		if s >= target&mask {
			i++
			continue
		}
		ans++
		for i++; cnt[i] == 0; i++ {
			ans++
		}
	}
	return
}
