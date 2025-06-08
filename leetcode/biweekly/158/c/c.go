package main

import (
	"math"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func maxGCDScore1(nums []int, k int) int64 {
	ans := 0
	for i := range nums {
		lowbitMin, lowbitCnt := math.MaxInt, 0
		g := 0
		for j := i; j >= 0; j-- {
			x := nums[j]
			lb := x & -x
			if lb < lowbitMin {
				lowbitMin, lowbitCnt = lb, 1
			} else if lb == lowbitMin {
				lowbitCnt++
			}

			g = gcd(g, x)
			newG := g
			if lowbitCnt <= k {
				newG *= 2
			}
			ans = max(ans, newG*(i-j+1))

			if g*2*(i+1) <= ans {
				break
			}
		}
	}
	return int64(ans)
}

func maxGCDScore(nums []int, k int) int64 {
	mx := bits.Len(uint(slices.Max(nums)))
	lowbitPos := make([][]int, mx)

	ans := 0
	type interval struct{ g, l, r int } // 左开右闭 (l,r]
	intervals := []interval{}
	for i, x := range nums {
		tz := bits.TrailingZeros(uint(x))
		lowbitPos[tz] = append(lowbitPos[tz], i) // 用 tz 代替 x 的 lowbit

		for j, p := range intervals {
			intervals[j].g = gcd(p.g, x)
		}
		intervals = append(intervals, interval{x, i - 1, i})

		// 去重（合并 g 相同的区间）
		idx := 1
		for j := 1; j < len(intervals); j++ {
			if intervals[j].g != intervals[j-1].g {
				intervals[idx] = intervals[j]
				idx++
			} else {
				intervals[idx-1].r = intervals[j].r
			}
		}
		intervals = intervals[:idx]

		// 此时我们将区间 [0,i] 划分成了 len(intervals) 个左闭右开区间
		// 对于任意 p∈intervals，任意 j∈(p.l,p.r]，gcd(区间[j,i]) 的计算结果均为 p.g
		for _, p := range intervals {
			// 不做任何操作
			ans = max(ans, p.g*(i-p.l))
			// 看看能否乘 2
			tz := bits.TrailingZeros(uint(p.g))
			pos := lowbitPos[tz]
			minL := p.l
			if len(pos) > k {
				minL = max(minL, pos[len(pos)-k-1])
			}
			if minL < p.r { // 可以乘 2
				ans = max(ans, p.g*2*(i-minL))
			}
		}
	}
	return int64(ans)
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
