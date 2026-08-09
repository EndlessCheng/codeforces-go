package main

import "slices"

// https://space.bilibili.com/206214
func minPrice(prices, discounts []int) float64 {
	slices.SortFunc(prices, func(a, b int) int { return b - a })
	slices.SortFunc(discounts, func(a, b int) int { return b - a })

	ans := 0
	for i, p := range prices {
		d := 0
		if i < len(discounts) {
			d = discounts[i]
		}
		ans += p * (100 - d)
	}
	return float64(ans) / 100
}
