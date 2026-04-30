package main

import "math"

func numSquarefulPerms(nums []int) (ans int) {
	n := len(nums)
	isSquare := make([][]bool, n)
	for i, x := range nums {
		isSquare[i] = make([]bool, n)
		for j, y := range nums {
			rt := int(math.Sqrt(float64(x + y)))
			isSquare[i][j] = rt*rt == x+y
		}
	}

	f := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := range f[0] {
		f[0][i] = 1
	}

	u := 1<<n - 1
	for s := 1; s < u; s++ {
		for i := range n {
			if s>>i&1 > 0 { // i 是填过的数的下标，不能在 s 中
				continue
			}
			for j := range n {
				if s>>j&1 > 0 && isSquare[i][j] {
					f[s][i] += f[s^1<<j][j]
				}
			}
		}
	}

	// 枚举排列的第一个数的下标
	for i := range n {
		ans += f[u^1<<i][i]
	}

	// 去重
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
		ans /= cnt[x] // 比如 nums 有 3 个 x，这里会 /1 再 /2 再 /3，从而实现 /(3!)
	}

	return
}
