package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func survivedRobotsHealths(positions []int, healths []int, directions string) (ans []int) {
	// 创建一个下标数组，对下标数组排序，这样不会打乱输入顺序
	idx := make([]int, len(positions))
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return positions[i] - positions[j] })

	st := []int{}
	for _, i := range idx {
		if directions[i] == 'R' { // 机器人 i 向右
			st = append(st, i)
			continue
		}
		for len(st) > 0 { // 栈顶机器人向右
			j := st[len(st)-1]
			if healths[j] > healths[i] { // 栈顶机器人的健康度大
				healths[i] = 0 // 移除机器人 i
				healths[j]--
				break
			}
			if healths[j] == healths[i] { // 健康度一样大，都移除
				healths[i] = 0
				healths[j] = 0
				st = st[:len(st)-1]
				break
			}
			// 机器人 i 的健康度大
			healths[i]--
			healths[j] = 0 // 移除机器人 j
			st = st[:len(st)-1]
		}
	}

	// 返回幸存机器人的健康度
	for _, h := range healths {
		if h > 0 {
			ans = append(ans, h)
		}
	}
	return
}
