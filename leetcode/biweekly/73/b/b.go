package main

import (
	"sort"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func sortJumbled(mapping []int, nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		s := []byte(strconv.Itoa(nums[i]))
		for k, ch := range s {
			s[k] = '0' + byte(mapping[ch&15])
		}
		x, _ := strconv.Atoi(string(s))
		s = []byte(strconv.Itoa(nums[j]))
		for k, ch := range s {
			s[k] = '0' + byte(mapping[ch&15])
		}
		y, _ := strconv.Atoi(string(s))
		return x < y
	})
	return nums
}
