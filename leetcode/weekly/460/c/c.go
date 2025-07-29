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
		if len(primeFactors[x]) == 1 { // x 是质数
			groups[x] = append(groups[x], i)
		}
	}

	vis := make([]bool, n)
	vis[n-1] = true
	q := []int{n - 1}
	for {
		tmp := q
		q = nil
		for _, i := range tmp {
			if i == 0 {
				return
			}
			if !vis[i-1] {
				vis[i-1] = true
				q = append(q, i-1)
			}
			if i < n-1 && !vis[i+1] {
				vis[i+1] = true
				q = append(q, i+1)
			}
			// 逆向思维：从 i 倒着跳到 nums[i] 的质因子 p 的下标 j
			for _, p := range primeFactors[nums[i]] {
				for _, j := range groups[p] {
					if !vis[j] {
						vis[j] = true
						q = append(q, j)
					}
				}
				groups[p] = nil // 避免重复访问下标列表
			}
		}
		ans++
	}
}

func minJumps1(nums []int) (ans int) {
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
