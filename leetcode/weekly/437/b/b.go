package main

import "slices"

// https://space.bilibili.com/206214
func maxWeight(pizzas []int) (ans int64) {
	slices.SortFunc(pizzas, func(a, b int) int { return b - a })
	days := len(pizzas) / 4
	odd := (days + 1) / 2
	for _, x := range pizzas[:odd] {
		ans += int64(x)
	}
	for i := range days / 2 {
		ans += int64(pizzas[odd+i*2+1])
	}
	return
}
