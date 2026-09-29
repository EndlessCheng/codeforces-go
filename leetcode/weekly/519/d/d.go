package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func solve(a []int, low, high int) (res int) {
	if low == high || len(a) <= 1 {
		return
	}

	var lowSt, highSt, b, c []int
	mid := (low + high) / 2

	for i, x := range a {
		if x <= mid { // x 在下部，作为 nums[i]
			for len(lowSt) > 0 && a[lowSt[len(lowSt)-1]] < x {
				lowSt = lowSt[:len(lowSt)-1] // 因为 x 的出现，栈顶不能作为 nums[i]
			}
			lowSt = append(lowSt, i)
			b = append(b, x)
		} else { // x 在上部，作为 nums[j]
			// 找到 x 左侧第一个小于 x 的最近元素，作为 nums[k]
			for len(highSt) > 0 && a[highSt[len(highSt)-1]] >= x {
				highSt = highSt[:len(highSt)-1]
			}
			res += len(lowSt)
			if len(highSt) > 0 {
				// lowSt 中 < highSt[len(highSt)-1] 的下标不能作为 nums[i]
				res -= sort.SearchInts(lowSt, highSt[len(highSt)-1])
			}
			highSt = append(highSt, i)
			c = append(c, x)
		}
	}

	return res + solve(b, low, mid) + solve(c, mid+1, high)
}

func shadowPairs1(nums []int) int {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	for i, x := range nums {
		nums[i] = sort.SearchInts(sorted, x)
	}

	return solve(nums, 0, len(sorted)-1)
}

func shadowPairs2(nums []int) (ans int) {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	var solve func([]int, int, int)
	solve = func(a []int, low, high int) {
		if low == high || len(a) <= 1 {
			return
		}

		var lowSt, highSt, b, c []int
		mid := (low + high) / 2
		midNum := sorted[mid]

		// 把 lowSt 按照 highSt 中的数划分成若干段，维护每一段的第一个数在 lowSt 中的位置
		segStarts := []int{}

		for i, x := range a {
			if x <= midNum { // x 在下部，作为 nums[i]
				for len(lowSt) > 0 && a[lowSt[len(lowSt)-1]] < x {
					lowSt = lowSt[:len(lowSt)-1] // 因为 x 的出现，栈顶不能作为 nums[i]
				}

				for len(segStarts) > 0 && segStarts[len(segStarts)-1] >= len(lowSt) {
					segStarts = segStarts[:len(segStarts)-1] // 最后一段是空的
				}

				if len(segStarts) == 0 || len(highSt) > 0 && highSt[len(highSt)-1] > lowSt[len(lowSt)-1] {
					// 即将插入 lowSt 的 i 是新段的第一个数
					segStarts = append(segStarts, len(lowSt))
				}

				lowSt = append(lowSt, i)
				b = append(b, x)
			} else { // x 在上部，作为 nums[j]
				// 找到 x 左侧第一个小于 x 的最近元素，作为 nums[k]
				for len(highSt) > 0 && a[highSt[len(highSt)-1]] >= x {
					highSt = highSt[:len(highSt)-1]
				}

				p := -1
				if len(highSt) > 0 {
					p = highSt[len(highSt)-1]
				}
				for len(segStarts) > 1 && lowSt[segStarts[len(segStarts)-2]] > p {
					segStarts = segStarts[:len(segStarts)-1] // 合并最后两段
				}

				// 统计 lowSt 中 > highSt[len(highSt)-1] 的元素个数
				if len(segStarts) > 0 && (len(highSt) == 0 || highSt[len(highSt)-1] < lowSt[len(lowSt)-1]) {
					ans += len(lowSt) - segStarts[len(segStarts)-1]
				}

				highSt = append(highSt, i)
				c = append(c, x)
			}
		}

		solve(b, low, mid)
		solve(c, mid+1, high)
	}

	solve(nums, 0, len(sorted)-1)
	return
}

func shadowPairs(nums []int) (ans int) {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	n := len(nums)
	memB := make([]int, 0, n)
	memC := make([]int, 0, n)

	var solve func([]int, int, int)
	solve = func(a []int, low, high int) {
		if low == high || len(a) <= 1 {
			return
		}

		var lowSt, highSt []int
		b := memB[:0]
		c := memC[:0]
		mid := (low + high) / 2
		midNum := sorted[mid]

		// 把 lowSt 按照 highSt 中的数划分成若干段，维护每一段的第一个数在 lowSt 中的位置
		segStarts := []int{}

		for i, x := range a {
			if x <= midNum { // x 在下部，作为 nums[i]
				for len(lowSt) > 0 && a[lowSt[len(lowSt)-1]] < x {
					lowSt = lowSt[:len(lowSt)-1] // 因为 x 的出现，栈顶不能作为 nums[i]
				}

				for len(segStarts) > 0 && segStarts[len(segStarts)-1] >= len(lowSt) {
					segStarts = segStarts[:len(segStarts)-1] // 最后一段是空的
				}

				if len(segStarts) == 0 || len(highSt) > 0 && highSt[len(highSt)-1] > lowSt[len(lowSt)-1] {
					// 即将插入 lowSt 的 i 是新段的第一个数
					segStarts = append(segStarts, len(lowSt))
				}

				lowSt = append(lowSt, i)
				b = append(b, x)
			} else { // x 在上部，作为 nums[j]
				// 找到 x 左侧第一个小于 x 的最近元素，作为 nums[k]
				for len(highSt) > 0 && a[highSt[len(highSt)-1]] >= x {
					highSt = highSt[:len(highSt)-1]
				}

				p := -1
				if len(highSt) > 0 {
					p = highSt[len(highSt)-1]
				}
				for len(segStarts) > 1 && lowSt[segStarts[len(segStarts)-2]] > p {
					segStarts = segStarts[:len(segStarts)-1] // 合并最后两段
				}

				// 统计 lowSt 中 > highSt[len(highSt)-1] 的元素个数
				if len(segStarts) > 0 && (len(highSt) == 0 || highSt[len(highSt)-1] < lowSt[len(lowSt)-1]) {
					ans += len(lowSt) - segStarts[len(segStarts)-1]
				}

				highSt = append(highSt, i)
				c = append(c, x)
			}
		}

		nb := len(b)
		copy(a, b)
		copy(a[nb:], c)
		solve(a[:nb], low, mid)
		solve(a[nb:], mid+1, high)
	}

	solve(nums, 0, len(sorted)-1)
	return
}
