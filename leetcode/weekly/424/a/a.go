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
		} else if pre*2 == total {
			ans += 2
		} else if abs(pre*2-total) == 1 {
			ans += 1
		}
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
