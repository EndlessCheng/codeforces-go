package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func splitPainting(segments [][]int) (ans [][]int64) {
	type event struct{ pos, color int }
	events := make([]event, 0, len(segments)*2)
	for _, seg := range segments {
		events = append(events, event{seg[0], seg[2]}, event{seg[1], -seg[2]}) // 记录每个线段左右端点相对应的线段颜色
	}
	sort.Slice(events, func(i, j int) bool { return events[i].pos < events[j].pos }) // 按位置排序

	sum := 0
	for i, e := range events[:len(events)-1] {
		sum += e.color
		if sum > 0 && e.pos < events[i+1].pos {
			ans = append(ans, []int64{int64(e.pos), int64(events[i+1].pos), int64(sum)})
		}
	}
	return
}
