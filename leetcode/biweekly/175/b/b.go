package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minimumK1(nums []int) int {
	n := len(nums)
	left := int(math.Ceil(math.Sqrt(float64(n)))) // 答案的下界
	right := slices.Max(nums)
	ans := left + sort.Search(right-left, func(k int) bool {
		k += left
		sum := n
		for _, x := range nums {
			sum += (x - 1) / k
		}
		return sum <= k*k
	})
	return ans
}

func minimumK(nums []int) int {
	n := len(nums)
	nonPositive := func(k int) int {
		sum := n
		for _, x := range nums {
			sum += (x - 1) / k
		}
		return sum
	}

	sum := 0
	for _, x := range nums {
		sum += x
	}

	left := max(int(math.Ceil(math.Sqrt(float64(n)))), int(math.Ceil(math.Cbrt(float64(sum))))) // 答案的下界
	right := int(math.Ceil(math.Sqrt(float64(nonPositive(left)))))                              // 答案的上界
	
	fmt.Println(left, right, right - left)
	
	ans := left + sort.Search(right-left, func(k int) bool {
		k += left
		return nonPositive(k) <= k*k
	})
	return ans
}

func main() {
	a := make([]int, 1e5)
	for i := range a {
		a[i] = 1e5
	}
	minimumK(a)
}
