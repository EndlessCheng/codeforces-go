package main

import (
	"math"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func minMergeCost1(lists [][]int) int64 {
	u := 1 << len(lists)
	sorted := make([][]int, u)
	for i, a := range lists { // 枚举不在 s 中的下标 i
		highBit := 1 << i
		for s, b := range sorted[:highBit] {
			sorted[highBit|s] = merge(a, b)
		}
	}

	f := make([]int, u)
	for i := range f {
		if i&(i-1) == 0 { // i 只包含一个元素，无法分解成两个非空子集
			continue // f[i] = 0
		}
		f[i] = math.MaxInt
		// 枚举 i 的非空真子集 j
		for j := i & (i - 1); j > i^j; j = (j - 1) & i {
			k := i ^ j // j 关于 i 的补集是 k
			lenJ := len(sorted[j])
			lenK := len(sorted[k])
			medJ := sorted[j][(lenJ-1)/2]
			medK := sorted[k][(lenK-1)/2]
			f[i] = min(f[i], f[j]+f[k]+lenJ+lenK+abs(medJ-medK))
		}
	}
	return int64(f[u-1])
}

func minMergeCost2(lists [][]int) int64 {
	u := 1 << len(lists)
	sumLen := make([]int, u)
	for i, a := range lists { // 枚举不在 s 中的下标 i
		highBit := 1 << i
		for s, sl := range sumLen[:highBit] {
			sumLen[highBit|s] = sl + len(a)
		}
	}

	median := make([]int, u)
	for mask, sl := range sumLen {
		k := (sl + 1) / 2
		left, right := int(-1e9), int(1e9)
		median[mask] = left + sort.Search(right-left, func(med int) bool {
			med += left
			cnt := 0
			for s := uint32(mask); s > 0; s &= s - 1 {
				i := bits.TrailingZeros32(s)
				cnt += sort.SearchInts(lists[i], med+1)
				if cnt >= k {
					return true
				}
			}
			return false
		})
	}

	f := make([]int, u)
	for i := range f {
		if i&(i-1) == 0 {
			continue // f[i] = 0
		}
		f[i] = math.MaxInt
		// 枚举 i 的非空真子集 j
		for j := i & (i - 1); j > i^j; j = (j - 1) & i {
			k := i ^ j // j 关于 i 的补集是 k
			f[i] = min(f[i], f[j]+f[k]+sumLen[j]+sumLen[k]+abs(median[j]-median[k]))
		}
	}
	return int64(f[u-1])
}

//

// 88. 合并两个有序数组（创建一个新数组）
func merge(a, b []int) []int {
	i, n := 0, len(a)
	j, m := 0, len(b)
	res := make([]int, 0, n+m)
	for {
		if i == n {
			return append(res, b[j:]...)
		}
		if j == m {
			return append(res, a[i:]...)
		}
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
}

func calcSorted(lists [][]int) [][]int {
	u := 1 << len(lists)
	sorted := make([][]int, u)
	for i, a := range lists {
		highBit := 1 << i
		for s, b := range sorted[:highBit] {
			sorted[highBit|s] = merge(a, b)
		}
	}
	return sorted
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays(a, b []int) int {
	if len(a) > len(b) {
		a, b = b, a
	}

	m, n := len(a), len(b)
	i := sort.Search(m, func(i int) bool {
		j := (m+n+1)/2 - i - 2
		return a[i] > b[j+1]
	}) - 1

	j := (m+n+1)/2 - i - 2
	if i < 0 {
		return b[j]
	}
	if j < 0 {
		return a[i]
	}
	return max(a[i], b[j])
}

func minMergeCost(lists [][]int) int64 {
	n := len(lists)
	m := n / 2
	sorted1 := calcSorted(lists[:m])
	sorted2 := calcSorted(lists[m:])

	u := 1 << n
	half := 1<<m - 1
	sumLen := make([]int, u) // 可以省略，但预处理出来，相比直接在后面 DP 中计算更快
	median := make([]int, u)
	for i := 1; i < u; i++ {
		// 把 i 分成低 m 位和高 n-m 位
		// 低 half 位去 sorted1 中找合并后的数组
		// 高 n-half 位去 sorted2 中找合并后的数组
		sumLen[i] = len(sorted1[i&half]) + len(sorted2[i>>m])
		median[i] = findMedianSortedArrays(sorted1[i&half], sorted2[i>>m])
	}

	f := make([]int, u)
	for i := range f {
		if i&(i-1) == 0 {
			continue
		}
		f[i] = math.MaxInt
		for j := i & (i - 1); j > i^j; j = (j - 1) & i {
			k := i ^ j
			f[i] = min(f[i], f[j]+f[k]+sumLen[i]+abs(median[j]-median[k]))
		}
	}
	return int64(f[u-1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
