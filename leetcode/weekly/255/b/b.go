package main

import (
	"fmt"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func findDifferentBinaryString(nums []string) string {
	has := map[int]bool{}
	for _, s := range nums {
		v, _ := strconv.ParseInt(s, 2, 64)
		has[int(v)] = true
	}
	v := 0
	for ; has[v]; v++ {
	}
	return fmt.Sprintf("%0*b", len(nums[0]), v)
}

// 康托对角线
func findDifferentBinaryString2(nums []string) (ans string) {
	for i, s := range nums {
		ans += string(s[i] ^ 1)
	}
	return
}
