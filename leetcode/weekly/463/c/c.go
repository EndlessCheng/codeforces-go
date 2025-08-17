package main

import "math"

// https://space.bilibili.com/206214
func minArraySum(nums []int, k int) int64 {
	minF := make([]int, k)
	// sum[0] = 0，对应的 f[0] = 0
	for i := 1; i < k; i++ {
		minF[i] = math.MaxInt
	}
	f, sum := 0, 0
	for _, x := range nums {
		sum = (sum + x) % k
		// 不删除 x，那么转移来源为 f + x
		// 删除 x，问题变成剩余前缀的最小和
		// 其中剩余前缀的元素和模 k 等于 sum，对应的 f 值的最小值记录在 minF[sum] 中
		f = min(f+x, minF[sum])
		// 维护前缀和 sum 对应的最小和，由于上面计算了 min，这里无需再计算 min
		minF[sum] = f
	}
	return int64(f)
}

func minArraySum2(nums []int, k int) int64 {
	minF := map[int]int{0: 0} // sum[0] = 0，对应的 f[0] = 0
	f, sum := 0, 0
	for _, x := range nums {
		sum = (sum + x) % k
		// 不删除 x
		f += x
		// 删除 x，问题变成剩余前缀的最小和
		// 其中剩余前缀的元素和模 k 等于 sum，对应的 f 值的最小值记录在 minF[sum] 中
		if mn, ok := minF[sum]; ok {
			f = min(f, mn)
		}
		// 维护前缀和 sum 对应的最小和，由于上面计算了 min，这里无需再计算 min
		minF[sum] = f
	}
	return int64(f)
}
