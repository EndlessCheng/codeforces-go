package main

import "sort"

// https://space.bilibili.com/206214
func rampartDefensiveLine(rampart [][]int) (ans int) {
	n := len(rampart)
	leftSpace := rampart[n-1][0] - rampart[0][1]
	for _, p := range rampart[1 : n-1] {
		leftSpace -= p[1] - p[0]
	}
	return sort.Search(leftSpace/(n-2), func(mx int) bool {
		mx++
		preR := rampart[0][1]
		for i := 1; i < n-1; i++ {
			r := rampart[i][1]
			space := mx - (rampart[i][0] - preR)
			if space > 0 {
				r += space // 向右膨胀
				if r > rampart[i+1][0] { // 无法膨胀
					return true
				}
			}
			preR = r
		}
		return false
	})
}
