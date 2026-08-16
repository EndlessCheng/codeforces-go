package main

import "math"

// https://space.bilibili.com/206214
func nearestDrone(drones [][]int, target []int) int {
	tx, ty := target[0], target[1]
	minDis := math.MaxInt
	ans := -1
	for i, d := range drones {
		dis := abs(tx-d[0]) + abs(ty-d[1])
		if dis < minDis && dis <= d[2] {
			minDis = dis
			ans = i
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
