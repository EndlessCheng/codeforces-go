package main

// https://space.bilibili.com/206214
func maxProductPair(nums []int, target int) []int {
	pos := map[int]int{}
	bestI, bestJ := -1, -1
	for i, x := range nums {
		if x*2 == target { // x = target-x
			continue
		}
		j, ok := pos[target-x]
		if ok && (bestI < 0 || x*(target-x) > nums[bestI]*(target-nums[bestI])) { // 也可以比谁离 target/2 更近
			bestI, bestJ = i, j
		}
		pos[x] = i
	}

	if bestI >= 0 && nums[bestI] < nums[bestJ] {
		bestI, bestJ = bestJ, bestI // 保证 nums[bestI] > nums[bestJ]
	}
	return []int{bestI, bestJ}
}
