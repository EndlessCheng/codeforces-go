package main

import (
	"index/suffixarray"
	"math/rand"
	"sort"
	"unsafe"
)

// https://space.bilibili.com/206214
func smallestUniqueSubarray2(nums []int) int {
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

func smallestUniqueSubarray(nums []int) int {
	n := len(nums)
	// 多项式哈希（方便计算子数组哈希值）
	// 哈希函数 hash(s) = nums[0] * base^(n-1) + nums[1] * base^(n-2) + ... + nums[n-2] * base + nums[n-1]
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8)
	powBase := make([]int, n+1) // powBase[i] = base^i
	preHash := make([]int, n+1) // 前缀哈希值 preHash[i] = hash(s[:i])
	powBase[0] = 1
	for i, x := range nums {
		powBase[i+1] = powBase[i] * base % mod
		preHash[i+1] = (preHash[i]*base + x) % mod // 秦九韶算法计算多项式哈希
	}
	// 计算子数组 nums[l:r] 的哈希值，注意这是左闭右开区间 [l,r)
	// 计算方法类似前缀和
	subHash := func(l, r int) int {
		return ((preHash[r]-preHash[l]*powBase[r-l])%mod + mod) % mod
	}

	ans := sort.Search(n, func(length int) bool {
		cnt := map[int]int{}
		for i := length; i <= n; i++ {
			h := subHash(i-length, i)
			cnt[h]++
		}
		for _, c := range cnt {
			if c == 1 {
				return true
			}
		}
		return false
	})
	return ans
}
