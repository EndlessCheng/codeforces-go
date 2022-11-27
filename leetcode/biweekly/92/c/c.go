package main

import "strings"

// https://space.bilibili.com/206214
func bestClosingTime(customers string) (ans int) {
	cost := strings.Count(customers, "Y")
	maxCost := cost
	for i, c := range customers {
		if c == 'N' {
			cost++
		} else {
			cost--
			if cost < maxCost {
				cost = maxCost
				ans = i + 1
			}
		}
	}
	return
}
