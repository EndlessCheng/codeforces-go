package main

func shuffle1(nums []int, n int) []int {
	ans := make([]int, n*2)
	for i, x := range nums[:n] {
		ans[i*2] = x
		ans[i*2+1] = nums[n+i]
	}
	return ans
}

func shuffle(nums []int, n int) []int {
	for i, x := range nums {
		if x < 0 { // 已访问
			continue
		}
		for cur := i; ; {
			// 元素 x 要填入 nums[nxt]
			nxt := cur * 2
			if cur >= n {
				nxt = (cur-n)*2 + 1
			}
			if nxt == i { // 回到起点
				nums[i] = -x // 用负数表示访问过
				break
			}
			// 把 x 填入 nums[nxt]（用负数表示访问过）
			// 同时把原来位于 nxt 的数记为 x
			x, nums[nxt] = nums[nxt], -x
			cur = nxt
		}
	}

	// 复原
	for i, x := range nums {
		nums[i] = -x
	}

	return nums
}
