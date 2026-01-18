package main

// https://space.bilibili.com/206214
func bestTower(towers [][]int, center []int, radius int) []int {
	cx, cy := center[0], center[1]
	maxQ, minX, minY := -1, -1, -1
	for _, t := range towers {
		x, y, q := t[0], t[1], t[2]
		if abs(x-cx)+abs(y-cy) <= radius &&
			(q > maxQ || q == maxQ && (x < minX || x == minX && y < minY)) {
			maxQ, minX, minY = q, x, y
		}
	}
	return []int{minX, minY}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
