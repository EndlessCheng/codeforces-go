package main

// https://space.bilibili.com/206214
func minGroupsForValidAssignment(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	k := len(nums)
	for _, c := range cnt {
		k = min(k, c)
	}
	for ; ; k-- {
		ans := 0
		for _, c := range cnt {
			if c/k < c%k {
				ans = 0
				break
			}
			ans += (c + k) / (k + 1)
		}
		if ans > 0 {
			return ans
		}
	}
}

func min(a, b int) int { if b < a { return b }; return a }
