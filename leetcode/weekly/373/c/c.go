package main

import "slices"

// https://space.bilibili.com/206214
func lexicographicallySmallestArray1(nums []int, limit int) []int {
	n := len(nums)
	pos := make([]int, n)
	for i := range pos {
		pos[i] = i
	}
	slices.SortFunc(pos, func(i, j int) int { return nums[i] - nums[j] })
	// 排序后，nums[pos[i]] 是递增的

	ans := make([]int, n)
	start := 0
	for i, p := range pos {
		if i == n-1 || nums[pos[i+1]]-nums[p] > limit { // 这一段的末尾
			// subPos 是 ans 中的一组空位（不一定有序）
			// 我们需要把 subPos 对应的 nums 中的数从小到大地填入空位（从左到右填）
			// 为了能从左到右填，需要把 subPos 排序
			subPos := slices.Clone(pos[start : i+1])
			slices.Sort(subPos)
			for j, q := range subPos {
				ans[q] = nums[pos[start+j]]
			}
			start = i + 1
		}
	}
	return ans
}

func lexicographicallySmallestArray2(nums []int, limit int) []int {
	n := len(nums)
	pos := make([]int, n)
	for i := range pos {
		pos[i] = i
	}
	slices.SortFunc(pos, func(i, j int) int { return nums[i] - nums[j] })
	// 排序后，nums[pos[i]] 是递增的

	groups := [][]int{}
	belong := make([]int, n)
	for i, p := range pos {
		if i == 0 || nums[p]-nums[pos[i-1]] > limit {
			groups = append(groups, []int{})
		}
		// 保存同一组内的数据，同时记录 pos[i] 属于哪一组
		gid := len(groups) - 1
		groups[gid] = append(groups[gid], nums[p])
		belong[p] = gid
	}

	ans := make([]int, n)
	for i, gid := range belong {
		ans[i] = groups[gid][0]
		groups[gid] = groups[gid][1:]
	}
	return ans
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	pos := make([]int, n)
	for i := range pos {
		pos[i] = i
	}
	slices.SortFunc(pos, func(i, j int) int { return nums[i] - nums[j] })
	// 排序后，nums[pos[i]] 是递增的

	groups := []int{}
	belong := make([]int, n)
	for i, p := range pos {
		if i == 0 || nums[p]-nums[pos[i-1]] > limit {
			groups = append(groups, i) // 新的段，只需记录开始下标
		}
		// 记录 pos[i] 属于哪一段
		belong[p] = len(groups) - 1
	}

	ans := make([]int, n)
	for i, gid := range belong {
		ans[i] = nums[pos[groups[gid]]]
		groups[gid]++
	}
	return ans
}
