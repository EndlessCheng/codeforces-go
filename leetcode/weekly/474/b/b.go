package main

// https://space.bilibili.com/206214
func maxProduct(nums []int) int64 {
	mx, mx2 := 0, 0
	for _, x := range nums {
		x = abs(x)
		if x > mx {
			mx2 = mx
			mx = x
		} else if x > mx2 {
			mx2 = x
		}
	}
	return int64(mx * mx2 * 1e5)
}

func abs(x int) int { if x < 0 { return -x }; return x }
