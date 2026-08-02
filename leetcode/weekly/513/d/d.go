package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick []int

func (t fenwick) add(i int) {
	for ; i < len(t); i += i & -i {
		t[i]++
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func countRatioSubarrays1(nums []int, a, b int) int64 {
	sum := make([]int, len(nums)+1)
	value := [2]int{-b, a}
	for i, x := range nums {
		sum[i+1] = sum[i] + value[x&1] // 偶数视作 -b，奇数视作 a
	}

	// sum 排序去重
	sorted := slices.Clone(sum)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	t := make(fenwick, len(sorted)+1)
	ans := 0
	for _, s := range sum {
		s = sort.SearchInts(sorted, s) + 1 // 离散化（从 1 开始）
		ans += t.pre(s)                    // 计算在 s 左边有多少个 <= s 的数
		t.add(s)
	}
	return int64(ans)
}

//

func mergeCount(sum []int) int {
	n := len(sum)
	if n <= 1 {
		return 0
	}

	left := slices.Clone(sum[:n/2])
	right := slices.Clone(sum[n/2:])
	cnt := mergeCount(left) + mergeCount(right) // left 和 right 各自的合法数对个数

	l, r := 0, 0
	for i := range sum {
		// 计算一个在 left 中，另一个在 right 中的合法数对个数
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			sum[i] = left[l]
			l++
		} else {
			cnt += l // left[:l] 中的数都 <= right[r]，这有 l 个
			sum[i] = right[r]
			r++
		}
	}

	return cnt
}

func countRatioSubarrays(nums []int, a, b int) int64 {
	sum := make([]int, len(nums)+1)
	value := [2]int{-b, a}
	for i, x := range nums {
		sum[i+1] = sum[i] + value[x&1] // 偶数视作 -b，奇数视作 a
	}

	return int64(mergeCount(sum))
}
