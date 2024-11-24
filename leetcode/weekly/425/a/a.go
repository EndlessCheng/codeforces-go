package main

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"math"
)

// https://space.bilibili.com/206214
func minimumSumSubarray(nums []int, l, r int) int {
	ans := math.MaxInt
	n := len(nums)
	s := make([]int, n+1)
	cnt := redblacktree.New[int, int]()
	for j := 1; j <= n; j++ {
		s[j] = s[j-1] + nums[j-1]
		if j < l {
			continue
		}
		c, _ := cnt.Get(s[j-l])
		cnt.Put(s[j-l], c+1)
		if lower, ok := cnt.Floor(s[j] - 1); ok {
			ans = min(ans, s[j]-lower.Key)
		}
		if j >= r {
			v := s[j-r]
			c, _ := cnt.Get(v)
			if c == 1 {
				cnt.Remove(v)
			} else {
				cnt.Put(v, c-1)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func minimumSumSubarray2(nums []int, l, r int) int {
	ans := math.MaxInt
	n := len(nums)
	for i := range n - l + 1 {
		s := 0
		for j := i; j < n && j-i+1 <= r; j++ {
			s += nums[j]
			if s > 0 && j-i+1 >= l {
				ans = min(ans, s)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
