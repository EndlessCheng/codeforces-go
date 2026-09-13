package main

import "sort"

// https://space.bilibili.com/206214
const mx = 2_000_000_002

var palindromes = [2][]int{{0}, {-1}}

// 预处理 [1, mx] 中的回文数
func init() {
	for base := 1; ; base *= 10 {
		// 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			if x > mx {
				return
			}
			// 按照 x 的奇偶性分组
			palindromes[x%2] = append(palindromes[x%2], x)
		}

		// 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
		for i := base; i < base*10; i++ {
			x := i
			for t := i; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			if x > mx {
				return
			}
			palindromes[x%2] = append(palindromes[x%2], x)
		}
	}
}

func minOperations(nums []int) (ans int64) {
	for _, x := range nums {
		p := palindromes[x%2]
		i := sort.SearchInts(p, x)
		ans += int64(min(p[i]-x, x-p[i-1]))
	}
	return ans / 2
}
