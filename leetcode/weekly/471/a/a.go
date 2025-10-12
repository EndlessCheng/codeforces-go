package main

// https://space.bilibili.com/206214
func sumDivisibleByK(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	for x, c := range cnt {
		if c%k == 0 {
			ans += x * c
		}
	}
	return
}
