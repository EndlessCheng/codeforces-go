package main

import "math"

// https://space.bilibili.com/206214
const mod = 1_000_000_007

func xorAfterQueries(nums []int, queries [][]int) (ans int) {
	n := len(nums)
	B := int(math.Sqrt(float64(len(queries))))
	diff := make([][]int, B)

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		if k < B {
			// 懒初始化
			if diff[k] == nil {
				diff[k] = make([]int, n+k)
				for j := range diff[k] {
					diff[k][j] = 1
				}
			}
			diff[k][l] = diff[k][l] * v % mod
			r = r - (r-l)%k + k
			diff[k][r] = diff[k][r] * pow(v, mod-2) % mod
		} else {
			for i := l; i <= r; i += k {
				nums[i] = nums[i] * v % mod
			}
		}
	}

	for k, d := range diff {
		if d == nil {
			continue
		}
		for start := range k {
			mulD := 1
			for i := start; i < n; i += k {
				mulD = mulD * d[i] % mod
				nums[i] = nums[i] * mulD % mod
			}
		}
	}

	for _, x := range nums {
		ans ^= x
	}
	return
}

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
