package main

import "sort"

// https://space.bilibili.com/206214
const mod int = 1e9 + 7

const mx int = 1e5
var omega [mx + 1]int8
func init() { // 预处理 omega
	for i := 2; i <= mx; i++ {
		if omega[i] == 0 {
			for j := i; j <= mx; j += i {
				omega[j]++
			}
		}
	}
}

func maximumScore(nums []int, k int) int {
	n := len(nums)
	left := make([]int, n)  // 质数分数 >= omega[nums[i]] 的最近的数的下标
	right := make([]int, n) // 质数分数 > omega[nums[i]] 的最近的数的下标
	for i := range right {
		right[i] = n
	}
	st := []int{-1}
	for i, v := range nums {
		for len(st) > 1 && omega[nums[st[len(st)-1]]] < omega[v] {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return nums[id[i]] > nums[id[j]] })

	ans := 1
	for _, i := range id {
		tot := (i - left[i]) * (right[i] - i)
		if tot >= k {
			ans = ans * pow(nums[i], k) % mod
			break
		}
		ans = ans * pow(nums[i], tot) % mod
		k -= tot
	}
	return ans
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
