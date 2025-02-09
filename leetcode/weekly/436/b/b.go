package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func assignElements(groups []int, elements []int) []int {
	mx := slices.Max(groups)
	target := make([]int, mx+1)
	for i := range target {
		target[i] = -1
	}

	for i, x := range elements {
		if x > mx || target[x] >= 0 { // x 及其倍数已被标记
			continue
		}
		for y := x; y <= mx; y += x { // 枚举 x 的倍数 y
			if target[y] < 0 {
				target[y] = i // 标记 y 可以被 x 整除
			}
		}
	}

	// 回答询问
	for i, x := range groups {
		groups[i] = target[x] // 原地修改
	}
	return groups
}
