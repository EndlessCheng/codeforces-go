package main

import "math/big"

// https://space.bilibili.com/206214
func minZeroArray(nums []int, queries [][]int) int {
	for _, x := range nums {
		if x > 0 {
			goto normal
		}
	}
	return 0
normal:
	f := make([]*big.Int, len(nums))
	for i := range f {
		f[i] = big.NewInt(1)
	}
	p := new(big.Int)
next:
	for k, q := range queries {
		val := uint(q[2])
		for i := q[0]; i <= q[1]; i++ {
			f[i].Or(f[i], p.Lsh(f[i], val))
		}
		for i, x := range nums {
			if f[i].Bit(x) == 0 {
				continue next
			}
		}
		return k + 1
	}
	return -1
}

func minZeroArray1(nums []int, queries [][]int) int {
	for _, x := range nums {
		if x > 0 {
			goto normal
		}
	}
	return 0
normal:
	n := len(nums)
	f := make([][]bool, n)
	for i, x := range nums {
		f[i] = make([]bool, x+1)
		f[i][0] = true
	}
next:
	for k, q := range queries {
		val := q[2]
		for i := q[0]; i <= q[1]; i++ {
			for j := nums[i]; j >= val; j-- {
				f[i][j] = f[i][j] || f[i][j-val]
			}
		}
		for i, x := range nums {
			if !f[i][x] {
				continue next
			}
		}
		return k + 1
	}
	return -1
}
