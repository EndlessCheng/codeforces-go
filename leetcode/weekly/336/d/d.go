package main

import "sort"

// https://space.bilibili.com/206214
func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
	type tuple struct{ l, r, s int }
	st := []tuple{{-2, -2, 0}} // 闭区间左右端点，栈底到栈顶的区间长度的和
	for _, p := range tasks {
		start, end, d := p[0], p[1], p[2]
		q := st[sort.Search(len(st), func(i int) bool { return st[i].l >= start })-1]
		d -= st[len(st)-1].s - q.s // 去掉运行中的时间点
		if start <= q.r { // start 在区间 q 内
			d -= q.r - start + 1 // 去掉运行中的时间点
		}
		if d <= 0 {
			continue
		}
		for end-st[len(st)-1].r <= d { // 剩余的 d 填充区间后缀
			top := st[len(st)-1]
			st = st[:len(st)-1]
			d += top.r - top.l + 1 // 合并区间
		}
		st = append(st, tuple{end - d + 1, end, st[len(st)-1].s + d})
	}
	return st[len(st)-1].s
}
