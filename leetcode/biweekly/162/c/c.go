package main

import "math"

// https://space.bilibili.com/206214
func solve(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
	minFinish := math.MaxInt
	for i, start := range landStartTime {
		minFinish = min(minFinish, start+landDuration[i])
	}

	res := math.MaxInt
	for i, start := range waterStartTime {
		res = min(res, max(start, minFinish)+waterDuration[i])
	}
	return res
}

func earliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
	landWater := solve(landStartTime, landDuration, waterStartTime, waterDuration)
	waterLand := solve(waterStartTime, waterDuration, landStartTime, landDuration)
	return min(landWater, waterLand)
}
