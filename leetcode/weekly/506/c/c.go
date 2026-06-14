package main

import "math"

// https://space.bilibili.com/206214
func maxRatings(units [][]int) int64 {
	ans := 0
	if len(units[0]) == 1 {
		// 每个设备都只有一个单元
		for _, unit := range units {
			ans += unit[0]
		}
		return int64(ans)
	}

	mn, mn2 := math.MaxInt, math.MaxInt
	for _, unit := range units {
		// 计算最小次小
		unitMin, unitMin2 := math.MaxInt, math.MaxInt
		for _, x := range unit {
			if x < unitMin {
				unitMin2 = unitMin
				unitMin = x
			} else if x < unitMin2 {
				unitMin2 = x
			}
		}

		ans += unitMin2 // 先加上次小
		mn2 = min(mn2, unitMin2)
		mn = min(mn, unitMin)
	}

	// 把包含 mn2 的那个设备作为集中站
	ans += mn - mn2 // mn2 改成全局最小
	return int64(ans)
}
