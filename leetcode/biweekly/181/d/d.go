package main

import "sort"

// https://space.bilibili.com/206214
func kthRemainingInteger1(nums []int, queries [][]int) []int {
	// 记录所有偶数的下标
	evenPos := []int{}
	for i, x := range nums {
		if x%2 == 0 {
			evenPos = append(evenPos, i)
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		// 找到询问对应的 evenPos 的子数组
		l := sort.SearchInts(evenPos, q[0])
		r := sort.SearchInts(evenPos, q[1]+1)
		pos := evenPos[l:r]
		k := q[2]

		// 二分答案是第几个正偶数
		// 二分下界：至少是第 k 个正偶数
		// 二分上界：至多是第 k+len(pos) 个正偶数
		left, right := k, k+len(pos)
		res := left + sort.Search(right-left, func(x int) bool {
			x += left
			// 计算子数组中的 <= x*2 的偶数个数
			j := sort.Search(len(pos), func(j int) bool { return nums[pos[j]] > x*2 })
			return x-j >= k
		})
		ans[i] = res * 2 // 答案是第 res 个正偶数
	}
	return ans
}

func kthRemainingInteger(nums []int, queries [][]int) []int {
	// 记录所有偶数的下标
	evenPos := []int{}
	for i, x := range nums {
		if x%2 == 0 {
			evenPos = append(evenPos, i)
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		// 找到询问对应的 evenPos 的子数组
		l := sort.SearchInts(evenPos, q[0])
		r := sort.SearchInts(evenPos, q[1]+1)
		pos := evenPos[l:r]
		k := q[2]

		// 推导过程见 1539 题解
		j := sort.Search(len(pos), func(j int) bool {
			return nums[pos[j]]/2-1-j >= k
		})
		ans[i] = (j + k) * 2
	}
	return ans
}
