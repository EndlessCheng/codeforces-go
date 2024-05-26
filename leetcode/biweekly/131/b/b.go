package main

// https://space.bilibili.com/206214
func occurrencesOfElement(nums, queries []int, x int) []int {
	pos := []int{}
	for i, v := range nums {
		if v == x {
			pos = append(pos, i)
		}
	}
	for i, q := range queries {
		if q > len(pos) {
			queries[i] = -1
		} else {
			queries[i] = pos[q-1]
		}
	}
	return queries
}
