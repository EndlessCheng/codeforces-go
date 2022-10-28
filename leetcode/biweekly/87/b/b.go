package main

import "sort"

// https://space.bilibili.com/206214
func matchPlayersAndTrainers(players, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	j, m := 0, len(players)
	for _, t := range trainers {
		if j < m && players[j] <= t {
			j++
		}
	}
	return j
}
