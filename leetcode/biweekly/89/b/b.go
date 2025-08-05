package main

import "math/bits"

// https://space.bilibili.com/206214
func productQueries1(n int, queries [][]int) []int {
	const mod = 1_000_000_007
	// 例如二进制 1100 分解为 100 + 1000
	// 第一轮循环 lowbit(1100) = 100，然后 1100 ^ 100 = 1000
	// 第二轮循环 lowbit(1000) = 1000，然后 1000 ^ 1000 = 0，循环结束
	powers := []int{}
	for n > 0 {
		lowbit := n & -n
		powers = append(powers, lowbit)
		n ^= lowbit
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		mul := 1
		for _, x := range powers[q[0] : q[1]+1] {
			mul = mul * x % mod
		}
		ans[i] = mul
	}
	return ans
}

func productQueries2(n int, queries [][]int) []int {
	const mod = 1_000_000_007
	powers := []int{}
	for n > 0 {
		lowbit := n & -n
		powers = append(powers, lowbit)
		n ^= lowbit
	}

	m := len(powers)
	res := make([][]int, m)
	for i, x := range powers {
		res[i] = make([]int, m)
		res[i][i] = x
		for j := i + 1; j < m; j++ {
			res[i][j] = res[i][j-1] * powers[j] % mod
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = res[q[0]][q[1]]
	}
	return ans
}

const mod = 1_000_000_007
const mx = 436

var pow2 = [mx]int{1}

func init() {
	for i := 1; i < mx; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}
}

func productQueries(n int, queries [][]int) []int {
	s := []int{0}
	for ; n > 0; n &= n - 1 { // n &= n-1 去掉 n 的最低比特 1
		e := bits.TrailingZeros(uint(n))
		// 直接计算 e 的前缀和
		s = append(s, s[len(s)-1]+e)
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		sumE := s[q[1]+1] - s[q[0]]
		ans[i] = pow2[sumE]
	}
	return ans
}
