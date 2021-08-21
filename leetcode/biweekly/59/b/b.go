package main

// github.com/EndlessCheng/codeforces-go
func maxMatrixSum(a [][]int) int64 {
	sum, min, neg := 0, int(1e9), false
	for _, r := range a {
		for _, v := range r {
			if v < 0 {
				neg = !neg
				v = -v
			}
			if v < min {
				min = v
			}
			sum += v
		}
	}
	if neg {
		sum -= min * 2
	}
	return int64(sum)
}
