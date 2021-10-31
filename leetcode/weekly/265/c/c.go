package main

// github.com/EndlessCheng/codeforces-go
func minimumOperations(nums []int, start, goal int) int {
	vis := [1001]bool{}
	vis[start] = true
	q := []int{start}
	for step := 1; q != nil; step++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, num := range nums {
				for _, x := range []int{v + num, v - num, v ^ num} {
					if x == goal {
						return step
					}
					// 由于超过范围无法继续运算，故不入队（否则入队太多会超时！）
					if 0 <= x && x <= 1000 && !vis[x] {
						vis[x] = true
						q = append(q, x)
					}
				}
			}
		}
	}
	return -1
}
