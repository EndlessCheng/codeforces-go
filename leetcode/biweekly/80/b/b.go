package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214/dynamic
func successfulPairs1(spells, potions []int, success int64) []int {
	slices.Sort(potions)
	for i, x := range spells {
		spells[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/x+1)
	}
	return spells
}

func successfulPairs(spells, potions []int, success int64) []int {
	mx := slices.Max(potions)
	cnt := make([]int, mx+1)
	for _, y := range potions {
		cnt[y]++ // 统计每种药水的出现次数
	}
	// 计算 cnt 的后缀和
	for i := mx - 1; i >= 0; i-- {
		cnt[i] += cnt[i+1]
	}
	// 计算完毕后，cnt[i] 就是 potions 值 >= i 的药水个数

	for i, x := range spells {
		low := (int(success)-1)/x + 1
		if low <= mx {
			spells[i] = cnt[low]
		} else {
			spells[i] = 0
		}
	}
	return spells
}
