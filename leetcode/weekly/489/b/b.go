package main

// https://space.bilibili.com/206214
func firstUniqueFreq(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	cc := make([]int, len(nums)+1)
	for _, c := range cnt {
		cc[c]++
	}

	for _, x := range nums {
		if cc[cnt[x]] == 1 {
			return x
		}
	}
	return -1
}
