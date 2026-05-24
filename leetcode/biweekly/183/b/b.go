package main

import (
	"math"
	"slices"
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//

func calc(a []int, k int) {
	slices.Sort(a)
	a = append(a, a[0]+k) // 哨兵

	
	for i := 1; i < len(a); i++ {
		
	}
}

func minOperations(nums []int, k int) int {
	a := [2][]int{}
	for i, x := range nums {
		a[i%2] = append(a[i%2], x)
	}

	ans := 0
	return ans
}
