package main

import "sort"

// https://space.bilibili.com/206214
func minInitialStrength1(monsters []int, boosts [][]int) int64 {
	n := len(monsters)
	bonus := make([]int, n+1)
	for _, b := range boosts {
		bonus[b[0]] += b[2]
		bonus[b[1]+1] -= b[2]
	}

	// 差分数组的前缀和即原数组
	for i := 1; i < n; i++ {
		bonus[i] += bonus[i-1]
	}

	sum := 0
	for _, x := range monsters {
		sum += x
	}

	ans := sort.Search(sum, func(strength int) bool {
		for i, x := range monsters {
			if strength+bonus[i] < x {
				return false
			}
			strength = max(strength-x, 0)
		}
		return true
	})
	return int64(ans)
}

func minInitialStrength2(monsters []int, boosts [][]int) int64 {
	n := len(monsters)
	bonus := make([]int, n+1)
	for _, b := range boosts {
		bonus[b[0]] += b[2]
		bonus[b[1]+1] -= b[2]
	}

	// 差分数组的前缀和即原数组
	for i := 1; i < n; i++ {
		bonus[i] += bonus[i-1]
	}

	f := 0
	for i := n - 1; i >= 0; i-- {
		if f > 0 {
			f += monsters[i]
		} else {
			f = max(monsters[i]-bonus[i], 0)
		}
	}
	return int64(f)
}

func minInitialStrength(monsters []int, boosts [][]int) int64 {
	n := len(monsters)
	bonus := make([]int, n+1)
	for _, b := range boosts {
		bonus[b[0]] += b[2]
		bonus[b[1]+1] -= b[2]
	}

	// 差分数组的前缀和即原数组
	for i := 1; i < n; i++ {
		bonus[i] += bonus[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		if monsters[i] > bonus[i] {
			ans := -bonus[i]
			for _, x := range monsters[:i+1] {
				ans += x
			}
			return int64(ans)
		}
	}
	return 0
}
