package main

import (
	"math"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func minMergeCost(lists [][]int) int64 {
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
		left, right := int(-1e9), int(1e9)
		median[mask] = left + sort.Search(right-left, func(med int) bool {
			med += left
			cnt := 0
			for s := uint32(mask); s > 0; s &= s - 1 {
				i := bits.TrailingZeros32(s)
				cnt += sort.SearchInts(lists[i], med+1)
				if cnt >= (sl+1)/2 {
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
		for j := i & (i - 1); j > 0; j = (j - 1) & i {
			k := i ^ j // j 关于 i 的补集是 k
			f[i] = min(f[i], f[j]+f[k]+sumLen[j]+sumLen[k]+abs(median[j]-median[k]))
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
