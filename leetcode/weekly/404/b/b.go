package main

// https://space.bilibili.com/206214
func maximumLength(nums []int) (ans int) {
	f := [2][2]int{}
	for _, x := range nums {
		x %= 2
		for y, fxy := range f[x] {
			f[y][x] = fxy + 1
			ans = max(ans, f[y][x])
		}
	}
	return
}
