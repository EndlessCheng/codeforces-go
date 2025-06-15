package main

// https://space.bilibili.com/206214
func specialTriplets(nums []int) (ans int) {
	const mod = 1_000_000_007
	suf := map[int]int{}
	for _, x := range nums {
		suf[x]++
	}

	pre := map[int]int{}
	for _, x := range nums { // x = nums[j]
		suf[x]-- // 撤销
		// 现在 pre 中的是 [0,j-1]，suf 中的是 [j+1,n-1]
		ans += pre[x*2] * suf[x*2]
		pre[x]++
	}
	return ans % mod
}
