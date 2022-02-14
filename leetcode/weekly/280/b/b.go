package main

import "sort"

// github.com/EndlessCheng/codeforces-go
type pair struct{ num, cnt int }

// 计算出现次数最多的两个元素及其出现次数
func getMaxCnt2(cnt map[int]int) []pair {
	a := make([]pair, 0, max(len(cnt), 2))
	for num, c := range cnt {
		a = append(a, pair{num, c})
	}
	sort.Slice(a, func(i, j int) bool { return a[i].cnt > a[j].cnt })
	return a[:2] // 不足两个时，用 pair{0, 0} 填充
}

func minimumOperations(nums []int) int {
	cnt := [2]map[int]int{{}, {}}
	for i, num := range nums {
		cnt[i&1][num]++
	}
	a0 := getMaxCnt2(cnt[0])
	a1 := getMaxCnt2(cnt[1])
	if a0[0].num != a1[0].num {
		return len(nums) - a0[0].cnt - a1[0].cnt // 不相等时，保留出现次数最多的两个
	}
	return len(nums) - max(a0[0].cnt+a1[1].cnt, a0[1].cnt+a1[0].cnt) // 相等时，保留出现次数最多的和另一个出现次数次多的
}

func max(a, b int) int { if b > a { return b }; return a }
