package main

import (
	"math"
	"math/bits"
	"slices"
	"sort"
	"strconv"
)

// https://space.bilibili.com/206214
var specialNumbers []int

func init() {
	const oddMask = 0x155
	for mask := 1; mask < 512; mask++ {
		t := mask & oddMask
		if t&(t-1) > 0 { // 至少有两个奇数
			continue
		}

		perm := []int{}
		size := 0
		odd := 0
		for s := uint(mask); s > 0; s &= s - 1 {
			x := bits.TrailingZeros(s) + 1
			size += x
			for range x / 2 {
				perm = append(perm, x)
			}
			if x%2 > 0 {
				odd = x
			}
		}
		if size > 16 {
			continue
		}

		permutations(len(perm), len(perm), func(idx []int) bool {
			pal := 0
			for _, i := range idx {
				pal = pal*10 + perm[i]
			}
			v := pal
			if odd > 0 {
				pal = pal*10 + odd
			}
			// 反转 pal 的左半，拼在 pal 后面
			for ; v > 0; v /= 10 {
				pal = pal*10 + v%10
			}
			specialNumbers = append(specialNumbers, pal)
			return false
		})
	}
	slices.Sort(specialNumbers)
}

func specialPalindrome(n int64) int64 {
	i := sort.SearchInts(specialNumbers, int(n+1))
	return int64(specialNumbers[i])
}

func permutations(n, r int, do func(ids []int) (Break bool)) {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	if do(ids[:r]) {
		return
	}
	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}
	for {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				tmp := ids[i]
				copy(ids[i:], ids[i+1:])
				ids[n-1] = tmp
				cycles[i] = n - i
			} else {
				j := cycles[i]
				ids[i], ids[n-j] = ids[n-j], ids[i]
				if do(ids[:r]) {
					return
				}
				break
			}
		}
		if i == -1 {
			return
		}
	}
}

var size [512]int

func init() {
	const oddMask = 0x155
	for mask := 1; mask < 512; mask++ {
		t := mask & oddMask
		if t&(t-1) > 0 { // 至少有两个奇数
			continue
		}
		for s := uint(mask); s > 0; s &= s - 1 {
			size[mask] += bits.TrailingZeros(s) + 1
		}
	}
}

func specialPalindrome1(Num int64) int64 {
	num := int(Num)
	targetSize := len(strconv.Itoa(num))
	ans := math.MaxInt
	for mask := 1; mask < 512; mask++ {
		sz := size[mask]
		if sz != targetSize && sz != targetSize+1 {
			continue
		}

		perm := make([]int, 0, sz/2)
		odd := 0
		for s := uint(mask); s > 0; s &= s - 1 {
			x := bits.TrailingZeros(s) + 1
			for range x / 2 {
				perm = append(perm, x)
			}
			if x%2 > 0 {
				odd = x
			}
		}

		permutations(len(perm), len(perm), func(idx []int) (Break bool) {
			x := 0
			for _, i := range idx {
				x = x*10 + perm[i]
			}
			tmp := x
			if odd > 0 {
				x = x*10 + odd
			}
			// 反转 x 的左半，拼在 x 后面
			for v := tmp; v > 0; v /= 10 {
				x = x*10 + v%10
			}
			if x >= ans { // 最优性剪枝，不再继续枚举
				return true
			}
			if x > num { // 满足要求
				ans = x
				return true
			}
			return false
		})
	}
	return int64(ans)
}
