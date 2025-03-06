package main

// https://space.bilibili.com/206214
func beautifulSubsets(nums []int, k int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	ans := 1
	for x, c := range cnt {
		if cnt[x-k] > 0 { // x 不是等差数列的首项
			continue
		}
		// 计算这一组的方案数
		f0, f1 := 1, 1<<c
		for x += k; cnt[x] > 0; x += k {
			f0, f1 = f1, f1+f0*(1<<cnt[x]-1)
		}
		ans *= f1 // 每组方案数相乘
	}
	return ans - 1 // 去掉空集
}
