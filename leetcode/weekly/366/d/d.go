package main

// https://space.bilibili.com/206214
func maxSum(nums []int, k int) (ans int) {
	const mod = 1_000_000_007
	cnt := [30]int{}
	for _, x := range nums {
		for i := range cnt {
			cnt[i] += x >> i & 1
		}
	}
	for ; k > 0; k-- {
		x := 0
		for i := range cnt {
			if cnt[i] > 0 {
				cnt[i]--
				x |= 1 << i
			}
		}
		ans = (ans + x*x) % mod
	}
	return
}
