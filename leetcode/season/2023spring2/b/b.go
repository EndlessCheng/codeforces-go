package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func rampartDefensiveLine(rampart [][]int) int {
	n := len(rampart)
	s := rampart[n-1][0] - rampart[0][1]
	for _, p := range rampart[1 : n-1] {
		s -= p[1] - p[0]
	}
	return sort.Search(s/(n-2), func(m int) bool {
		m++
		preR := math.MinInt / 2
		for i, p := range rampart[:n-1] {
			r := p[1]
			space := m - (p[0] - preR) // 向左膨胀后的剩余长度
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
