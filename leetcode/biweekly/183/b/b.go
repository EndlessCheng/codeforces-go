package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minOperations1(nums []int, k int) int {
	ans := math.MaxInt
	for x := range k {
		for y := range k {
			if y == x {
				continue
			}
			target := [2]int{x, y}
			sum := 0
			for i, v := range nums {
				d := abs(v%k - target[i%2])
				sum += min(d, k-d) // 直接走到 target[i%2]，或者反向绕一圈到 target[i%2]
			}
			ans = min(ans, sum)
		}
	}
	return ans
}

func calc(a []int, k int) (int, int, int) {
	n := len(a)
	slices.Sort(a)
	for _, x := range a {
		a = append(a, x+k)
	}

	sum := make([]int, n*2+1)
	for i, x := range a {
		sum[i+1] = sum[i] + x
	}

	// 都变成 target 的最小操作次数
	calcOp := func(target int) int {
		i := sort.SearchInts(a[:n], target)
		j := i + sort.SearchInts(a[i:i+n], target+k/2+1)
		return (sum[j] - sum[i]) - (j-i)*target + // [i, j) 中的数都减小到 target
			(n-j+i)*(target+k) - (sum[i+n] - sum[j]) // [j, i+n) 中的数都增大到 target+k
	}

	mn, mn2, bestX := math.MaxInt, math.MaxInt, 0
	for i, x := range a[:n] {
		if i > 0 && a[i] == a[i-1] { // 优化：相同的值无需重复计算
			continue
		}
		op := calcOp(x)
		// 维护最小次小操作次数
		if op < mn {
			mn2 = mn
			mn, bestX = op, x
		} else if op < mn2 {
			mn2 = op
		}
	}

	// 还可以都变成 bestX-1 或者 bestX+1
	mn2 = min(mn2, calcOp((bestX-1+k)%k), calcOp((bestX+1)%k))

	return mn, mn2, bestX
}

func minOperations(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}

	a := [2][]int{}
	for i, x := range nums {
		a[i%2] = append(a[i%2], x%k)
	}

	min1x, min2x, bestX := calc(a[0], k)
	min1y, min2y, bestY := calc(a[1], k)

	if bestX != bestY {
		return min1x + min1y
	}
	return min(min1x+min2y, min2x+min1y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
