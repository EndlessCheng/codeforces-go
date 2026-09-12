package main

import "slices"

// https://space.bilibili.com/206214
func carFleet1(target int, position, speed []int) int {
	type pair struct{ p, v int }
	a := make([]pair, len(position))
	for i, p := range position {
		a[i] = pair{p, speed[i]}
	}
	slices.SortFunc(a, func(a, b pair) int { return a.p - b.p })

	st := []float64{}
	for _, p := range a {
		t := float64(target-p.p) / float64(p.v)
		for len(st) > 0 && st[len(st)-1] <= t {
			st = st[:len(st)-1]
		}
		st = append(st, t)
	}
	return len(st)
}

func carFleet2(target int, position, speed []int) int {
	type pair struct{ p, v int }
	a := make([]pair, len(position))
	for i, p := range position {
		a[i] = pair{p, speed[i]}
	}
	slices.SortFunc(a, func(a, b pair) int { return a.p - b.p })

	st := []pair{}
	for _, p := range a {
		for len(st) > 0 {
			q := st[len(st)-1]
			if (target-q.p)*p.v > (target-p.p)*q.v {
				break
			}
			st = st[:len(st)-1]
		}
		st = append(st, p)
	}
	return len(st)
}

func carFleet(target int, position, speed []int) (ans int) {
	n := len(position)
	type pair struct{ p, v int }
	a := make([]pair, n)
	for i, p := range position {
		a[i] = pair{p, speed[i]}
	}
	// 降序排序，这样下面可以正序遍历
	slices.SortFunc(a, func(a, b pair) int { return b.p - a.p })

	maxT := 0.0
	for _, p := range a {
		t := float64(target-p.p) / float64(p.v)
		if t > maxT {
			maxT = t
			ans++
		}
	}
	return
}
