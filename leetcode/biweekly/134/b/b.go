package main

import "math"

// https://space.bilibili.com/206214
func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
	mn, s := math.MaxInt, 0
	for _, e := range enemyEnergies {
		mn = min(mn, e)
		s += e
	}
	if currentEnergy < mn {
		return 0
	}
	return int64((currentEnergy + s - mn) / mn)
}
