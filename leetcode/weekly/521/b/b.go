package main

// https://space.bilibili.com/206214
func maxEqualAdjacentPairs(nums []int) int {
	base := 0
	type pair struct{ x, y int }
	cnt := map[pair]int{}

	for i := 1; i < len(nums); i++ {
		x, y := nums[i-1], nums[i]
		if x == y {
			base++
		} else {
			// 把 (x,y) 和 (y,x) 都统一为 (x,y)
			if x > y {
				x, y = y, x
			}
			// 统计相邻且不相等的数对个数
			cnt[pair{x, y}]++
		}
	}

	maxCnt := 0
	for _, c := range cnt {
		maxCnt = max(maxCnt, c)
	}

	return base + maxCnt
}
