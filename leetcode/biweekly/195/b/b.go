package main

// https://space.bilibili.com/206214
func canRemoveAll(types, colors, groupSize []int) bool {
	type tuple struct{ tp, color, cnt int }
	st := []tuple{}
	for i, tp := range types {
		color := colors[i]
		if len(st) > 0 && tp == st[len(st)-1].tp && color == st[len(st)-1].color {
			st[len(st)-1].cnt++
		} else {
			st = append(st, tuple{tp, color, 1})
		}
		// 只要发现满足要求的连续元素，就立刻移除
		if groupSize[st[len(st)-1].tp] == st[len(st)-1].cnt {
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}
