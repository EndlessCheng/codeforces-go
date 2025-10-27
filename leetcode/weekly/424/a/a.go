package main

// https://space.bilibili.com/206214
func countValidSelections(nums []int) (ans int) {
	total := 0
	for _, x := range nums {
		total += x
	}

	pre := 0
	for _, x := range nums {
		if x > 0 {
			pre += x
		} else {
			ans += max(2-abs(pre*2-total), 0)
		}
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
