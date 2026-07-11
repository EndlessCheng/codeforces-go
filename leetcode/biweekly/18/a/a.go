package main

import (
	"slices"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func arrayRankTransform1(arr []int) []int {
	// 排序去重
	sortedArr := slices.Clone(arr)
	slices.Sort(sortedArr)
	sortedArr = slices.Compact(sortedArr)

	for i, x := range arr {
		// 二分得到编号
		arr[i] = sort.SearchInts(sortedArr, x) + 1
	}
	return arr
}

func arrayRankTransform(arr []int) []int {
	// 排序
	sortedArr := slices.Clone(arr)
	slices.Sort(sortedArr)

	// 去重的同时构建哈希表
	rank := make(map[int]int, len(sortedArr))
	for i, x := range sortedArr {
		if i == 0 || x != sortedArr[i-1] {
			rank[x] = len(rank) + 1
		}
	}

	for i, x := range arr {
		arr[i] = rank[x]
	}
	return arr
}
