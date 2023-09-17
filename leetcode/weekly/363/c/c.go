package main

import "sort"

// https://space.bilibili.com/206214
func maxNumberOfAlloys(_, _, budget int, composition [][]int, stock, cost []int) (ans int) {
	mx := stock[0]
	for _, s := range stock {
		mx = min(mx, s)
	}
	mx += budget
	for _, com := range composition {
		res := sort.Search(mx, func(num int) bool {
			num++
			money := 0
			for i, s := range stock {
				if s < com[i]*num {
					money += (com[i]*num - s) * cost[i]
					if money > budget {
						return true
					}
				}
			}
			return false
		})
		ans = max(ans, res)
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
