package main

// https://space.bilibili.com/206214
func lastVisitedIntegers(nums []int) (ans []int) {
	seen := []int{}
	k := 0
	for _, x := range nums {
		if x > 0 {
			seen = append(seen, x)
			k = 0
		} else {
			k++
			if k > len(seen) {
				ans = append(ans, -1)
			} else {
				ans = append(ans, seen[len(seen)-k]) // 倒数第 k 个
			}
		}
	}
	return
}
