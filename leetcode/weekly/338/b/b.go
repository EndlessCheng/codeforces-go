package main

import "sort"

// https://space.bilibili.com/206214
var p = []int{0} // 哨兵，避免二分越界

func init() {
	const mx = 1000
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			p = append(p, i) // 预处理质数列表
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func primeSubOperation(nums []int) bool {
	pre := 0
	for _, x := range nums {
		if x <= pre {
			return false
		}
		pre = x - p[sort.SearchInts(p, x-pre)-1] // < x-pre 的最大质数
	}
	return true
}
