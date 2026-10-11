package main

// https://space.bilibili.com/206214
func resilientSubarray(nums []int, k int) (ans int) {
	cnt := 0
	for i, x := range nums {
		r := x % k
		if i > 0 && r == nums[i-1]%k {
			cnt++
		} else { // 新的模 k 同余段
			cnt = 1
		}

		// 去掉一个数后，剩余元素之和是 k 的倍数
		if r*(cnt-1)%k == 0 {
			ans = max(ans, cnt)
		}
	}
	return
}
