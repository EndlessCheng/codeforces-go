package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func comb2(num int) int {
	return num * (num - 1) / 2
}

func subsequencesWithMiddleMode(nums []int) int {
	n := len(nums)
	ans := n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120 // 所有方案数
	suf := map[int]int{}
	for _, num := range nums {
		suf[num]++
	}
	pre := make(map[int]int, len(suf)) // 预分配空间

	var cp, cs, ps, p2s, ps2 int
	for _, c := range suf {
		cs += comb2(c)
	}

	// 枚举 x，作为子序列正中间的数
	for left, x := range nums[:n-2] {
		suf[x]--

		px := pre[x]
		sx := suf[x]

		cs -= sx
		ps -= px
		p2s -= px * px
		ps2 -= (sx*2 + 1) * px

		right := n - 1 - left
		ans -= comb2(left-px) * comb2(right-sx)
		ans -= (cp - comb2(px)) * sx * (right - sx)
		ans -= (cs - comb2(sx)) * px * (left - px)
		ans -= ((ps-px*sx)*(right-sx) - (ps2 - px*sx*sx)) * px
		ans -= ((ps-px*sx)*(left-px)  - (p2s - px*px*sx)) * sx

		cp += px
		ps += sx
		ps2 += sx * sx
		p2s += (px*2 + 1) * sx

		pre[x]++
	}
	return ans % 1_000_000_007
}

func subsequencesWithMiddleMode2(nums []int) int {
	n := len(nums)
	ans := n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120 // 所有方案数
	suf := map[int]int{}
	for _, num := range nums {
		suf[num]++
	}
	pre := make(map[int]int, len(suf)) // 预分配空间
	// 枚举 x，作为子序列正中间的数
	for left, x := range nums[:n-2] {
		suf[x]--
		if left > 1 {
			right := n - 1 - left
			preX, sufX := pre[x], suf[x]
			// 不合法：只有一个 x
			ans -= comb2(left-preX) * comb2(right-sufX)
			// 不合法：只有两个 x，且至少有两个 y（y != x）
			for y, sufY := range suf { // 注意 sufY 可能是 0
				if y == x {
					continue
				}
				preY := pre[y]
				// 左边有两个 y，右边有一个 x，即 yy x xz（z 可以等于 y）
				ans -= comb2(preY) * sufX * (right - sufX)
				// 右边有两个 y，左边有一个 x，即 zx x yy（z 可以等于 y）
				ans -= comb2(sufY) * preX * (left - preX)
				// 左右各有一个 y，另一个 x 在左边，即 xy x yz（z != y）
				ans -= preY * sufY * preX * (right - sufX - sufY)
				// 左右各有一个 y，另一个 x 在右边，即 zy x xy（z != y）
				ans -= preY * sufY * sufX * (left - preX - preY)
			}
		}
		pre[x]++
	}
	return ans % 1_000_000_007
}

func subsequencesWithMiddleMode3(nums []int) int {
	n := len(nums)
	ans := n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120 // 所有方案数

	a := slices.Clone(nums)
	slices.Sort(a)
	a = slices.Compact(a)
	for i, x := range nums {
		nums[i] = sort.SearchInts(a, x)
	}

	suf := make([]int, len(a))
	for _, x := range nums {
		suf[x]++
	}
	pre := make([]int, len(a))
	// 枚举 x，作为子序列正中间的数
	for left, x := range nums[:n-2] {
		suf[x]--
		if left > 1 {
			right := n - 1 - left
			preX, sufX := pre[x], suf[x]
			// 不合法：只有一个 x
			ans -= comb2(left-preX) * comb2(right-sufX)
			// 不合法：只有两个 x，且至少有两个 y（y != x）
			for y, sufY := range suf { // 注意 sufY 可能是 0
				if y == x {
					continue
				}
				preY := pre[y]
				// 左边有两个 y，右边有一个 x，即 yy x xz（z 可以等于 y）
				ans -= comb2(preY) * sufX * (right - sufX)
				// 右边有两个 y，左边有一个 x，即 zx x yy（z 可以等于 y）
				ans -= comb2(sufY) * preX * (left - preX)
				// 左右各有一个 y，另一个 x 在左边，即 xy x yz（z != y）
				ans -= preY * sufY * preX * (right - sufX - sufY)
				// 左右各有一个 y，另一个 x 在右边，即 zy x xy（z != y）
				ans -= preY * sufY * sufX * (left - preX - preY)
			}
		}
		pre[x]++
	}
	return ans % 1_000_000_007
}
