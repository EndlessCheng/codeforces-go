package main

// https://space.bilibili.com/206214
func bowlSubarrays(nums []int) (ans int64) {
	st := []int{}
	for i, x := range nums {
		for len(st) > 0 && nums[st[len(st)-1]] < x {
			// j=st[len(st)-1] 右侧严格大于 nums[j] 的数的下标是 i
			if i-st[len(st)-1] > 1 { // 子数组的长度至少为 3
				ans++
			}
			st = st[:len(st)-1]
		}
		// i 左侧大于等于 nums[i] 的数的下标是 st[len(st)-1]
		if len(st) > 0 && i-st[len(st)-1] > 1 { // 子数组的长度至少为 3
			ans++
		}
		st = append(st, i)
	}
	return
}
