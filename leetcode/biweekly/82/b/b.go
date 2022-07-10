package main

import "sort"

// https://space.bilibili.com/206214/dynamic
func latestTimeCatchTheBus(buses, passengers []int, capacity int) (ans int) {
	sort.Ints(buses)
	sort.Ints(passengers)
	j, c := 0, 0
	for _, t := range buses {
		for c = capacity; c > 0 && j < len(passengers) && passengers[j] <= t; j++ {
			c--
		}
	}
	if c > 0 {
		ans = buses[len(buses)-1] // 最后一班公交还有空位，在它发车时到达
	} else {
		ans = passengers[j-1] // 上一个上车的乘客
	}
	for j--; j >= 0 && passengers[j] == ans; j-- { // 往前找可以插队的位置
		ans--
	}
	return
}
