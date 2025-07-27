package main

// https://space.bilibili.com/206214
const mx = 1_000_001

var primeFactors = [mx][]int{}

func init() {
	// 预处理每个数的质因子列表，思路同埃氏筛
	for i := 2; i < mx; i++ {
		if primeFactors[i] == nil { // i 是质数
			for j := i; j < mx; j += i { // i 的倍数有质因子 i
				primeFactors[j] = append(primeFactors[j], i)
			}
		}
	}
}

func minJumps(nums []int) (ans int) {
	n := len(nums)
	groups := map[int][]int{}
	for i, x := range nums {
		for _, p := range primeFactors[x] {
			groups[p] = append(groups[p], i) // 对于质数 p，可以跳到下标 i
		}
	}

	vis := make([]bool, n)
	vis[0] = true
	q := []int{0}
	for {
		tmp := q
		q = nil
		for _, i := range tmp {
			if i == n-1 {
				return
			}
			idx := groups[nums[i]]
			idx = append(idx, i+1)
			if i > 0 {
				idx = append(idx, i-1)
			}
			for _, j := range idx { // 可以从 i 跳到 j
				if !vis[j] {
					vis[j] = true
					q = append(q, j)
				}
			}
			delete(groups, nums[i]) // 避免重复访问下标列表
		}
		ans++
	}
}
