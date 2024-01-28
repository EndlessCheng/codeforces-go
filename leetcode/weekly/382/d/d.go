package main

// https://space.bilibili.com/206214
func minOrAfterOperations(nums []int, k int) (ans int) {
	mask := 0
	for b := 29; b >= 0; b-- {
		mask |= 1 << b
		cnt := 0
		and := -1
		for _, x := range nums {
			and &= x & mask
			if and != 0 {
				cnt++ // 合并 x 和下一个数
			} else {
				and = -1 // 准备合并下一段
			}
		}
		if cnt > k {
			ans |= 1 << b  // 这个比特位必须填 1
			mask ^= 1 << b // 后面不考虑这个比特位
		}
	}
	return
}
