package main

import "slices"

// https://space.bilibili.com/206214
func matchPlayersAndTrainers1(players, trainers []int) int {
	slices.Sort(players)
	slices.Sort(trainers)
	j, m := 0, len(trainers)
	for i, p := range players {
		for j < m && trainers[j] < p {
			j++
		}
		if j == m { // 无法找到匹配的训练师
			return i
		}
		j++ // 匹配一位训练师
	}
	return len(players) // 所有运动员都有匹配的训练师
}

func matchPlayersAndTrainers(players, trainers []int) int {
	slices.Sort(players)
	slices.Sort(trainers)
	j, m := 0, len(players)
	for _, t := range trainers {
		if j < m && players[j] <= t {
			j++
		}
	}
	return j
}
