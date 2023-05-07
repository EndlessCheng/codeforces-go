package main

// https://space.bilibili.com/206214
func colorTheArray(n int, queries [][]int) []int {
	ans := make([]int, len(queries))
	a := make([]int, n+2) // 避免讨论下标出界的情况
	cnt := 0
	for qi, q := range queries {
		i, c := q[0]+1, q[1] // 下标改成从 1 开始
		if a[i] > 0 {
			if a[i] == a[i-1] {
				cnt--
			}
			if a[i] == a[i+1] {
				cnt--
			}
		}
		a[i] = c
		if a[i] == a[i-1] {
			cnt++
		}
		if a[i] == a[i+1] {
			cnt++
		}
		ans[qi] = cnt
	}
	return ans
}
