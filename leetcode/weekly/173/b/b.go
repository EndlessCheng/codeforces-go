package main

import "sort"

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) (ans []int) {
	type pair struct{ l, r int }
	ps := []pair{}
	for _, r := range restaurants {
		id, rate, veganFriend, price, distance := r[0], r[1], r[2], r[3], r[4]
		if veganFriendly == 1 && veganFriend == 0 || price > maxPrice || distance > maxDistance {
			continue
		}
		ps = append(ps, pair{rate, id})
	}
	sort.Slice(ps, func(i, j int) bool {
		pi, pj := ps[i], ps[j]
		return pi.l > pj.l || pi.l == pj.l && pi.r > pj.r
	})
	for _, p := range ps {
		ans = append(ans, p.r)
	}
	return
}
