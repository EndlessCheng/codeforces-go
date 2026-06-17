package main

// https://space.bilibili.com/206214
func maximumLength(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	ans := cnt[1] - 1 | 1 // 保证 ans 是奇数（奇数不变，偶数减一）
	delete(cnt, 1)

	for x := range cnt {
		res := 0
		for cnt[x] >= 2 {
			res += 2
			x *= x
		}
		res += cnt[x]
		ans = max(ans, res-1|1) // 保证 ans 是奇数（奇数不变，偶数减一）
	}

	return ans
}
