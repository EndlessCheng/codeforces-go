package main

// github.com/EndlessCheng/codeforces-go
func countPoints(ps [][]int, qs [][]int) []int {
	ans := make([]int, len(qs))
	for i, q := range qs {
		x, y, r := q[0], q[1], q[2]
		for _, p := range ps {
			if (p[0]-x)*(p[0]-x)+(p[1]-y)*(p[1]-y) <= r*r {
				ans[i]++
			}
		}
	}
	return ans
}
