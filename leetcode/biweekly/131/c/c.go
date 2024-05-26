package main

// https://space.bilibili.com/206214
func queryResults(_ int, queries [][]int) []int {
	ans := make([]int, len(queries))
	color := map[int]int{}
	cnt := map[int]int{}
	for i, q := range queries {
		x, y := q[0], q[1]
		if c := color[x]; c > 0 {
			cnt[c]--
			if cnt[c] == 0 {
				delete(cnt, c)
			}
		}
		color[x] = y
		cnt[y]++
		ans[i] = len(cnt)
	}
	return ans
}
