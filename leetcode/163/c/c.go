package main

import "sort"

func maxSumDivThree(nums []int) int {
	arr1 := []int{}
	arr2 := []int{}
	ans := 0
	for _, v := range nums {
		if v%3 == 0 {
			ans += v
		} else if v%3 == 1 {
			arr1 = append(arr1, v)
		} else {
			arr2 = append(arr2, v)
		}
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	n1, n2 := len(arr1), len(arr2)

	if len(arr1) >= 6 {
		start := n1%3 + 3
		for _, v := range arr1[start:] {
			ans += v
		}
		arr1 = arr1[:start]
	}
	if len(arr2) >= 6 {
		start := n2%3 + 3
		for _, v := range arr2[start:] {
			ans += v
		}
		arr2 = arr2[:start]
	}
	n1, n2 = len(arr1), len(arr2)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	base := ans
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			// choose i arr1 and j arr2
			if (i+2*j)%3 != 0 {
				continue
			}
			tmpSum := base
			for ii := n1 - 1; ii >= n1-i; ii-- {
				tmpSum += arr1[ii]
			}
			for ii := n2 - 1; ii >= n2-j; ii-- {
				tmpSum += arr2[ii]
			}
			ans = max(ans, tmpSum)
		}
	}
	return ans
}
