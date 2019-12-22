package main

import "strconv"

func findNumbers(nums []int) (cnt int) {
	for _, v := range nums {
		if len(strconv.Itoa(v))%2 == 0 {
			cnt++
		}
	}
	return cnt
}
