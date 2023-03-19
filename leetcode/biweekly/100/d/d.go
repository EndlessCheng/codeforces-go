package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func repairCars(ranks []int, cars int) int64 {
	minR := ranks[0]
	cnt := [101]int{}
	for _, r := range ranks {
		if r < minR {
			minR = r
		}
		cnt[r]++
	}
	return int64(sort.Search(minR*cars*cars, func(t int) bool {
		s := 0
		for r := minR; r <= 100; r++ {
			s += int(math.Sqrt(float64(t/r))) * cnt[r]
		}
		return s >= cars
	}))
}
