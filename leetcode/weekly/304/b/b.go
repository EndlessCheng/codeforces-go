package main

import (
	"math"
)

/*

下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---


*/

// https://space.bilibili.com/206214
func maximumGroups(grades []int) int {
	n := len(grades)
	x := int((math.Sqrt(float64(1+8*n)) - 1) / 2)
	if (x+1)*(x+2)/2 <= n {
		x++
	}
	return x
}
