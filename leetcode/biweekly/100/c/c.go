package main

// https://space.bilibili.com/206214
func findScore(nums []int) (ans int64) {
	for i, n := 0, len(nums); i < n; i += 2 { // i 选了 i+1 不能选
		i0 := i
		for i+1 < n && nums[i] > nums[i+1] { // 找到下坡的坡底
			i++
		}
		for j := i; j >= i0; j -= 2 { // 从坡底 i 到坡顶 i0，每隔一个累加
			ans += int64(nums[j])
		}
	}
	return
}
