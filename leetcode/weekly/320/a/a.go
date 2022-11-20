package main

// https://space.bilibili.com/206214
func unequalTriplets(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	a, c := 0, len(nums)
	for _, b := range cnt {
		c -= b
		ans += a * b * c
		a += b
	}
	return
}
