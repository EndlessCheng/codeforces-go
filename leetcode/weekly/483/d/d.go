package main

import "math"

// https://space.bilibili.com/206214
func minMergeCost(lists [][]int) int64 {
	u := 1 << len(lists)
	sumLen := make([]int, u)
	sorted := make([][]int, u)
	median := make([]int, u)
	for i, a := range lists { // 枚举不在 s 中的下标 i
		highBit := 1 << i
		for s, sl := range sumLen[:highBit] {
			t := highBit | s
			sumLen[t] = sl + len(a)
			b := merge(sorted[s], a)
			sorted[t] = b
			median[t] = b[(len(b)-1)/2]
		}
	}

	f := make([]int, u)
	for i := range f {
		if i&(i-1) == 0 {
			continue // f[i] = 0
		}
		f[i] = math.MaxInt
		// 枚举 i 的非空真子集 j
		for j := i & (i - 1); j > 0; j = (j - 1) & i {
			k := i ^ j // j 关于 i 的补集是 k
			f[i] = min(f[i], f[j]+f[k]+sumLen[j]+sumLen[k]+abs(median[j]-median[k]))
		}
	}
	return int64(f[u-1])
}

// 88. 合并两个有序数组（创建一个新数组）
func merge(a, b []int) []int {
	i, n := 0, len(a)
	j, m := 0, len(b)
	res := make([]int, 0, n+m)
	for i < n || j < m {
		if j == m || i < n && a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
