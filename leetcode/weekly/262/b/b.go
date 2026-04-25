package main

import "slices"

/* 中位数

要使任意两元素最终相等，这两个元素之差必须是 $x$ 的倍数，否则无法通过加减 $x$ 来相等。

假设要让所有元素均为 $y$，可以发现：
- $y$ 每增加 $x$，小于或等于 $y$ 的元素要多操作一次，大于 $y$ 的元素要少操作一次；
- $y$ 每减小 $x$，大于或等于 $y$ 的元素要多操作一次，小于 $y$ 的元素要少操作一次。

因此 $y$ 选在所有元素的中位数上是最「均衡」的。

*/

// github.com/EndlessCheng/codeforces-go
func minOperations(grid [][]int, x int) int {
	k := len(grid) * len(grid[0])
	a := make([]int, 0, k) // 预分配空间
	target := grid[0][0] % x

	// 1. 判断是否无解
	for _, row := range grid {
		for _, v := range row {
			if v%x != target { // 每个数模 x 都必须相等
				return -1
			}
		}
		a = append(a, row...)
	}

	// 2. 计算 grid 的中位数 median
	slices.Sort(a)
	median := a[k/2]

	// 3. 计算操作次数
	ans := 0
	for _, v := range a {
		ans += abs(v - median)
	}
	return ans / x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
