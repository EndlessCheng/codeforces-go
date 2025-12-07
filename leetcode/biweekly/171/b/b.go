package main

import (
	"math"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
var palindromes []int

// 预处理二进制回文数
func init() {
	const mx = 5000
	const base = 2

	// 哨兵，0 也是回文数
	palindromes = append(palindromes, 0)

outer:
	for pw := 1; ; pw *= base {
		// 生成奇数长度回文数
		for i := pw; i < pw*base; i++ {
			x := i
			for t := i / base; t > 0; t /= base {
				x = x*base + t%base
			}
			if x > mx {
				break outer
			}
			palindromes = append(palindromes, x)
		}

		// 生成偶数长度回文数
		for i := pw; i < pw*base; i++ {
			x := i
			for t := i; t > 0; t /= base {
				x = x*base + t%base
			}
			if x > mx {
				break outer
			}
			palindromes = append(palindromes, x)
		}
	}

	// 哨兵，5049 是大于 5000 的第一个二进制回文数
	palindromes = append(palindromes, 5049)
}

func minOperations1(nums []int) []int {
	for i, x := range nums {
		j := sort.SearchInts(palindromes, x)
		// 变成 >= x 的二进制回文数，或者变成 < x 的二进制回文数
		nums[i] = min(palindromes[j]-x, x-palindromes[j-1])
	}
	return nums
}

//

func minOperations(nums []int) []int {
	for i, x := range nums {
		res := math.MaxInt
		n := bits.Len(uint(x))
		m := n / 2
		left := x >> m
		for l := left - 1; l <= left+1; l++ {
			// 左半反转到右半
			// 如果 n 是奇数，那么去掉回文中心再反转
			right := bits.Reverse(uint(l>>(n%2))) >> (bits.UintSize - m)
			pal := l<<m | int(right)
			res = min(res, abs(x-pal))
		}
		nums[i] = res
	}
	return nums
}

func abs(x int) int { if x < 0 { return -x }; return x }
