package main

// github.com/EndlessCheng/codeforces-go
func minDifference(a []int, qs [][]int) []int {
	sum := make([][101]int, len(a)+1)
	for i, v := range a {
		sum[i+1] = sum[i]
		sum[i+1][v]++
	}
	ans := make([]int, len(qs))
outer:
	for i, q := range qs {
		l, r, d, pre := q[0], q[1]+1, int(1e9), int(-1e9)
		for v := 1; v <= 100; v++ {
			cnt := sum[r][v] - sum[l][v] // v 的个数
			if cnt == r-l { // 所有元素都相同
				ans[i] = -1
				continue outer
			}
			if cnt > 0 { // 子数组包含元素 v
				if v-pre < d {
					d = v - pre
				}
				pre = v
			}
		}
		ans[i] = d
	}
	return ans
}
