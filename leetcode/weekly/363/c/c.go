package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxNumberOfAlloys(_, _, budget int, composition [][]int, stock, cost []int) (ans int) {
	mx := slices.Min(stock) + budget
	for _, comp := range composition {
		ans += sort.Search(mx-ans, func(num int) bool {
			num += ans + 1
			money := 0
			for i, s := range stock {
				if s < comp[i]*num {
					money += (comp[i]*num - s) * cost[i]
					if money > budget {
						return true
					}
				}
			}
			return false
		})
	}
	return
}
