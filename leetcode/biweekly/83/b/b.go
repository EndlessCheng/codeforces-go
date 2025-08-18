package main

// https://space.bilibili.com/206214/dynamic
func zeroFilledSubarray1(nums []int) (ans int64) {
	last := -1
	for i, x := range nums {
		if x != 0 {
			last = i // 记录上一个非 0 元素的位置
		} else {
			ans += int64(i - last)
		}
	}
	return
}

func zeroFilledSubarray(nums []int) (ans int64) {
	cnt0 := 0
	for _, x := range nums {
		if x != 0 {
			cnt0 = 0
		} else {
			cnt0++ // 右端点为 i 的全 0 子数组比右端点为 i-1 的全 0 子数组多一个
			ans += int64(cnt0)
		}
	}
	return
}
