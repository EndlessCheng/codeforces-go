package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func triangularSum1(nums []int) int {
	// 每循环一轮，数组长度就减一
	for n := len(nums) - 1; n > 0; n-- {
		for i := range n {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
	}
	return nums[0]
}

const mod = 10

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

const mx = 1000

// 计算组合数，需要计算阶乘及其逆元
var (
	f    [mx + 1]int // f[n] = n!
	invF [mx + 1]int // invF[n] = n!^-1
	p2   [mx + 1]int // n! 中的 2 的幂次
	p5   [mx + 1]int // n! 中的 5 的幂次
)

func init() {
	f[0] = 1
	invF[0] = 1
	for i := 1; i <= mx; i++ {
		x := i

		// 分离质因子 2，计算 2 的幂次
		e2 := bits.TrailingZeros(uint(x))
		x >>= e2

		// 分离质因子 5，计算 5 的幂次
		e5 := 0
		for x%5 == 0 {
			e5++
			x /= 5
		}

		f[i] = f[i-1] * x % mod
		invF[i] = pow(f[i], 3) // 欧拉定理求逆元
		p2[i] = p2[i-1] + e2
		p5[i] = p5[i-1] + e5
	}
}

var pow2 = [4]int{2, 4, 8, 6}

func comb(n, k int) int {
	res := f[n] * invF[k] * invF[n-k]
	e2 := p2[n] - p2[k] - p2[n-k]
	if e2 > 0 {
		res *= pow2[(e2-1)%4]
	}
	if p5[n]-p5[k]-p5[n-k] > 0 {
		res *= 5
	}
	return res
}

func triangularSum(nums []int) (ans int) {
	for i, x := range nums {
		ans += comb(len(nums)-1, i) * x
	}
	return ans % mod
}
