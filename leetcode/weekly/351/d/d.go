package main

import "sort"

// https://space.bilibili.com/206214
func survivedRobotsHealths(positions, healths []int, directions string) []int {
	type data struct {
		i, p, h int
		d       byte
	}
	a := make([]data, len(positions))
	for i, p := range positions {
		a[i] = data{i, p, healths[i], directions[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].p < a[j].p })

	var toLeft, st []data
next:
	for _, p := range a {
		if p.d == 'R' { // 向右，存入栈中
			st = append(st, p)
			continue
		}
		// p 向左，与栈中向右的机器人碰撞
		for len(st) > 0 {
			top := &st[len(st)-1]
			if top.h > p.h { // 栈顶的健康度大
				top.h--
				continue next
			}
			if top.h == p.h { // 健康度一样大
				st = st[:len(st)-1]
				continue next
			}
			p.h-- // p 的健康度大
			st = st[:len(st)-1] // 移除栈顶
		}
		toLeft = append(toLeft, p)
	}

	// 合并剩余的机器人
	toLeft = append(toLeft, st...)
	sort.Slice(toLeft, func(i, j int) bool { return toLeft[i].i < toLeft[j].i })
	ans := make([]int, len(toLeft))
	for i, p := range toLeft {
		ans[i] = p.h
	}
	return ans
}
