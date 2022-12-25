package main

import "sort"

// https://space.bilibili.com/206214
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	return sort.Search(price[len(price)-1], func(d int) bool {
		d++
		cnt, x0 := 1, price[0]
		for _, x := range price[1:] {
			if x >= x0+d {
				cnt++
				x0 = x
			}
		}
		return cnt < k
	})
}
