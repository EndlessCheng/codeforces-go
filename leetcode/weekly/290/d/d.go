package main

import "sort"

// https://space.bilibili.com/206214
func fullBloomFlowers(flowers [][]int, people []int) []int {
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

	id := make([]int, len(people))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return people[id[i]] < people[id[j]] })

	j, sum := 0, 0
	for _, i := range id {
		for ; j < n && times[j] <= people[i]; j++ {
			sum += diff[times[j]]
		}
		people[i] = sum
	}
	return people
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
