package main

import (
	"index/suffixarray"
	"unsafe"
)

// https://space.bilibili.com/206214
func smallestUniqueSubarray(nums []int) int {
	n := len(nums)
	// 把 1 个 int 拆成 3 个 byte，从而可以调用库函数计算后缀数组
	tmp := make([]byte, 0, n*3)
	for _, x := range nums {
		tmp = append(tmp, byte(x>>16), byte(x>>8), byte(x))
	}

	type _tp struct {
		_  []byte
		sa []int32
	}
	_sa := (*_tp)(unsafe.Pointer(suffixarray.New(tmp))).sa

	sa := make([]int32, 0, n)
	for _, p := range _sa {
		if p%3 == 0 { // 是 3 的倍数的 _sa[i] 就对应着 nums 的 sa[i]
			sa = append(sa, p/3)
		}
	}

	// 计算后缀名次数组
	// 后缀 nums[i:] 位于后缀字典序中的第 rank[i] 个
	// 特别地，rank[0] 即 nums 在后缀字典序中的排名，rank[n-1] 即 nums[n-1:] 在字典序中的排名
	// 相当于 sa 的反函数，即 rank[sa[i]] = i
	rank := make([]int, n)
	for i, p := range sa {
		rank[p] = i
	}

	// 计算高度数组（也叫 LCP 数组）
	// height[0] = 0（哨兵）
	// height[i] = LCP(nums[sa[i]:], nums[sa[i-1]:])  (i > 0)
	// 获取 nums[i] 所在位置的高度：height[rank[i]]
	height := make([]int, n)
	h := 0
	// 计算 nums 与 nums[sa[rank[0]-1]:] 的 LCP（记作 LCP0）
	// 计算 nums[1:] 与 nums[sa[rank[1]-1]:] 的 LCP（记作 LCP1）
	// 计算 nums[2:] 与 nums[sa[rank[2]-1]:] 的 LCP
	// ...
	// 计算 nums[n-1:] 与 nums[sa[rank[n-1]-1]:] 的 LCP
	// 从 LCP0 到 LCP1，我们只去掉了 nums[0] 和 nums[sa[rank[0]-1]] 这两个数
	// 所以 LCP1 >= LCP0 - 1
	// 这样就能加快 LCP 的计算了（类似滑动窗口）
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && nums[i+h] == nums[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	ans := n
	for i, h := range height {
		// 对于后缀 nums[sa[i]:]，其长为 uniqueLength 的前缀是唯一的
		uniqueLength := h + 1
		if i < n-1 {
			uniqueLength = max(h, height[i+1]) + 1
		}
		// 注意 uniqueLength 不能超过后缀 nums[sa[i]:] 的长度
		if uniqueLength <= n-int(sa[i]) {
			ans = min(ans, uniqueLength)
		}
	}
	return ans
}
