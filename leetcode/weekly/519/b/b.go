package main

import (
	"math"
	"sort"
	"strconv"
)

// https://space.bilibili.com/206214
const mx = 2_000_000_002

var palindromes = [2][]int{{0}, {0}} // 哨兵

// 预处理 [1, mx] 中的回文数
func init() {
	for base := 1; ; base *= 10 {
		// 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10 // 去掉 i 的最低位，反转，拼在 i 的右边
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
				x = x*10 + t%10 // 反转 i，拼在 i 的右边
			}
			if x > mx {
				return
			}
			palindromes[x%2] = append(palindromes[x%2], x)
		}
	}
}

func minOperations1(nums []int) (ans int64) {
	for _, x := range nums {
		p := palindromes[x%2]
		i := sort.SearchInts(p, x)
		ans += int64(min(p[i]-x, x-p[i-1]))
	}
	return ans / 2
}

// 返回离 num 最近的与 num 同奇偶的正回文数与 num 的绝对差
func nearestPalindromicDiff(num int) int {
	if num <= 9 {
		return 0 // num 已经是回文数
	}

	minD := math.MaxInt
	update := func(pal int) {
		minD = min(minD, abs(pal-num))
	}

	s := strconv.Itoa(num)
	m := len(s) // num 的十进制长度

	if num%2 > 0 {
		update(int(math.Pow10(m-1)) - 1) // 十进制长为 m-1 的最大奇回文数 999..999
		update(int(math.Pow10(m)) + 1)   // 十进制长为 m+1 的最小奇回文数 100..001
	} else {
		if m == 2 {
			update(8)
		} else {
			update(int(math.Pow10(m-2))*9 - 2) // 十进制长为 m-1 的最大偶回文数 899..998
		}
		update(int(math.Pow10(m))*2 + 2) // 十进制长为 m+1 的最小偶回文数 200..002
	}

	highDigit := int(s[0] - '0')
	if highDigit%2 == num%2 {
		left, _ := strconv.Atoi(s[:(m+1)/2])
		// 枚举十进制长为 m 的邻近回文数
		for l := left - 1; l <= left+1; l++ {
			// l 最高位的奇偶性必须与 num 的相同
			if int(strconv.Itoa(l)[0]%2) != num%2 {
				continue
			}

			pal := l
			x := l
			if m%2 > 0 {
				x /= 10
			}
			for ; x > 0; x /= 10 {
				pal = pal*10 + x%10
			}
			update(pal)
		}
	} else {
		// 最高位为 highDigit - 1
		// 例如 num = 354..676，生成回文数 299..992
		if highDigit > 1 {
			update((int(math.Pow10(m-1))+1)*highDigit - 11)
		}

		// 最高位为 highDigit + 1
		// 例如 num = 354..676，生成回文数 400..004
		if highDigit < 9 {
			update((int(math.Pow10(m-1)) + 1) * (highDigit + 1))
		}
	}

	return minD
}

func minOperations(nums []int) (ans int64) {
	for _, x := range nums {
		ans += int64(nearestPalindromicDiff(x))
	}
	return ans / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
