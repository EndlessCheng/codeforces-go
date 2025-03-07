package main

// https://space.bilibili.com/206214
func findMatrix(nums []int) (ans [][]int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for len(cnt) > 0 {
		row := make([]int, 0, len(cnt)) // 预分配空间
		for x := range cnt {
			row = append(row, x)
			if cnt[x]--; cnt[x] == 0 {
				delete(cnt, x)
			}
		}
		ans = append(ans, row)
	}
	return
}
