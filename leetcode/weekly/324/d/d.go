package main

// https://space.bilibili.com/206214
func cycleLengthQueries(_ int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i, q := range queries {
		res := 1
		for a, b := q[0], q[1]; a != b; res++ {
			if a > b {
				a /= 2
			} else {
				b /= 2
			}
		}
		ans[i] = res
	}
	return ans
}
