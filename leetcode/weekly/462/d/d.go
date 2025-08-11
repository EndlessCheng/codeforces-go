package main

import (
	"math"
	"math/bits"
	"slices"
	"sort"
	"strconv"
)

// https://space.bilibili.com/206214
func specialPalindrome(Num int64) int64 {
	subSum := make([]int, 1<<9)
	a := []int{1, 9, 8, 7, 6, 5, 4, 3, 2}
	mp := [10]int{}
	oddMask := 0
	for i, v := range a {
		mp[v] = i
		if v%2 > 0 {
			oddMask |= 1 << i
		}
		highBit := 1 << i
		for mask, s := range subSum[:highBit] {
			subSum[highBit|mask] = s + v
		}
	}
	for mask := 1; mask < 1<<9; mask++ {
		t := mask & oddMask
		if t&(t-1) > 0 {
			subSum[mask] = -1
		}
	}

	full := func(pal, odd int) int {
		v := pal
		if odd > 0 {
			pal = pal*10 + odd
		}
		for ; v > 0; v /= 10 {
			pal = pal*10 + v%10
		}
		return pal
	}

	cnt := [10]int{}
	extend := func(pal, mask, odd int) int {
		for j := 1; j < 10; j++ {
			need := 0
			if mask>>mp[j]&1 > 0 {
				need = j
			} else if cnt[j] > 0 {
				need = j - cnt[j]
			}
			for range need / 2 {
				pal = pal*10 + j
			}
			if need > 0 && j%2 > 0 {
				odd = j
			}
		}
		return full(pal, odd)
	}

	num := int(Num)
	s := strconv.Itoa(num)
	n := len(s)

	ans := 0
	var dfs func(int, int, int, bool) bool
	dfs = func(i, pal, odd int, limit bool) bool {
		if odd > 0 && n%2 == 0 {
			return false
		}

		if i == n/2 {
			if n%2 == 0 {
				for j, c := range cnt {
					if c > 0 && c != j {
						return false
					}
				}
				// 左半反转到右半
				pal = full(pal, 0)
				if pal > num {
					ans = pal
				}
				return pal > num
			} else { // else 可以省略，为了代码格式对齐保留
				if odd == 0 {
					odd = 1
				}
				cnt[odd]++
				defer func() { cnt[odd]-- }()
				for j, c := range cnt {
					if c > 0 && c != j {
						return false
					}
				}
				// 左半反转到右半
				pal = full(pal, odd)
				if pal > num {
					ans = pal
				}
				return pal > num
			}
		}

		if !limit {
			// 中间随便填
			left := n - i*2
			mask := 1<<9 - 1
			for j, c := range cnt {
				if c > 0 {
					left -= j - c
					mask ^= 1 << mp[j]
				}
			}
			if left < 0 {
				return false
			}

			// 枚举 mask 的子集，从大到小
			for sub, ok := mask, true; ok; ok = sub != mask {
				if (odd == 0 || sub&oddMask == 0) && subSum[sub] == left {
					ans = extend(pal, sub, odd)
					return true
				}
				sub = (sub - 1) & mask
			}
			return false
		}

		low := int(s[i] - '0')
		for v := low; v <= 9; v++ {
			if cnt[v]+2 > v || odd > 0 && v%2 > 0 && v != odd {
				continue
			}
			newOdd := odd
			if v%2 > 0 {
				newOdd = v
			}
			cnt[v] += 2
			if dfs(i+1, pal*10+v, newOdd, limit && v == low) {
				return true
			}
			cnt[v] -= 2
		}
		return false
	}

	if dfs(0, 0, 0, true) {
		return int64(ans)
	}

	// 没找到就取长为 n+1 的最小回文数
	for mask := 1<<9 - 1; ; mask-- {
		if subSum[mask] == n+1 {
			return int64(extend(0, mask, 0))
		}
	}
}

//

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
	specialNumbers = slices.Compact(specialNumbers)
}

func specialPalindrome2(n int64) int64 {
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
