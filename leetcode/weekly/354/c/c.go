package main

// https://space.bilibili.com/206214
func minimumIndex(nums []int) int {
	// 也可以用摩尔投票实现
	freq := map[int]int{}
	mode := nums[0]
	for _, x := range nums {
		freq[x]++
		if freq[x] > freq[mode] {
			mode = x
		}
	}

	total := freq[mode]
	freq1 := 0
	for i, x := range nums {
		if x == mode {
			freq1++
		}
		if freq1*2 > i+1 && (total-freq1)*2 > len(nums)-1-i {
			return i
		}
	}
	return -1
}
