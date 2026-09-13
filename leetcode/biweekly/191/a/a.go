package main

// https://space.bilibili.com/206214
func countSpecialIntegers(nums []int) (ans int) {
	pos := map[int][]int{}
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}

	for _, p := range pos {
		if len(p) == 3 && p[1]-p[0] == p[2]-p[1] {
			ans++
		}
	}
	return
}
