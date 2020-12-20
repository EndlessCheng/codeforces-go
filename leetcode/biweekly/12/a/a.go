package main

import "sort"

// github.com/EndlessCheng/codeforces-go
type Leaderboard struct{}

var s map[int]int

func Constructor() (_ Leaderboard) {
	s = map[int]int{}
	return
}

func (Leaderboard) AddScore(playerId, score int) {
	s[playerId] += score
}

func (Leaderboard) Top(k int) (ans int) {
	a := make([]int, 0, len(s))
	for _, v := range s {
		a = append(a, v)
	}
	sort.Ints(a)
	for _, v := range a[len(a)-k:] {
		ans += v
	}
	return
}

func (Leaderboard) Reset(playerId int) {
	s[playerId] = 0
}

/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
