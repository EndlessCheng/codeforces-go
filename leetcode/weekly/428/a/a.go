package main

// https://space.bilibili.com/206214
func buttonWithLongestTime(events [][]int) int {
	idx, maxDiff := events[0][0], events[0][1]
	for i := 1; i < len(events); i++ {
		p, q := events[i-1], events[i]
		d := q[1] - p[1]
		if d > maxDiff || d == maxDiff && q[0] < idx {
			idx, maxDiff = q[0], d
		}
	}
	return idx
}
