package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func fullBloomFlowers(flowers [][]int, persons []int) []int {
	diff := map[int]int{}
	for _, f := range flowers {
		diff[f[0]]++
		diff[f[1]+1]--
	}

	n := len(diff)
	times := make([]int, 0, n)
	for t := range diff {
		times = append(times, t)
	}
	sort.Ints(times)

	for i, p := range persons {
		persons[i] = p<<32 | i
	}
	sort.Ints(persons)

	ans := make([]int, len(persons))
	i, sum := 0, 0
	for _, p := range persons {
		for ; i < n && times[i] <= p>>32; i++ {
			sum += diff[times[i]]
		}
		ans[uint32(p)] = sum
	}
	return ans
}

func fullBloomFlowers2(flowers [][]int, persons []int) []int {
	n := len(flowers) * 2
	events := make([]int, 0, n)
	for _, f := range flowers {
		events = append(events, f[0]<<1|1, (f[1]+1)<<1)
	}
	sort.Ints(events)

	for i, p := range persons {
		persons[i] = p<<32 | i
	}
	sort.Ints(persons)

	ans := make([]int, len(persons))
	i, cnt := 0, 0
	for _, p := range persons {
		for ; i < n && events[i]>>1 <= p>>32; i++ {
			cnt += events[i]&1<<1 - 1
		}
		ans[uint32(p)] = cnt
	}
	return ans
}
