package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214/dynamic
func successfulPairs(spells, potions []int, success int64) []int {
	slices.Sort(potions)
	for i, x := range spells {
		spells[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/x+1)
	}
	return spells
}
