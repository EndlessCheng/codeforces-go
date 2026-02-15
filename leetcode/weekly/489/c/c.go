package main

import (
	"index/suffixarray"
	"math/bits"
	"slices"
	"unsafe"
)

// https://space.bilibili.com/206214
func almostPalindromic1(s string) (ans int) {
	n := len(s)
	expand := func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		ans = max(ans, r-l-1) // [l+1, r-1] 是回文串
	}

	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		expand(l-1, r) // 删除 s[l]，继续扩展
		expand(l, r+1) // 删除 s[r]，继续扩展
		if ans >= n {  // 优化：提前返回答案
			return n
		}
	}
	return
}

//

type sparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

// 时间复杂度 O(n * log n)
func newSparseTable[T any](a []T, op func(T, T) T) sparseTable[T] {
	n := len(a)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], a)
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = op(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return sparseTable[T]{st, op}
}

// [l,r) 左闭右开，下标从 0 开始
// 返回 op(nums[l:r])
// 时间复杂度 O(1)
func (s sparseTable[T]) query(l, r int) T {
	k := bits.Len(uint(r-l)) - 1
	return s.op(s.st[k][l], s.st[k][r-1<<k])
}

func suffixArrayLCP(s string) func(int, int) int {
	// 后缀数组 sa（后缀序）
	// sa[i] 表示后缀字典序中的第 i 个字符串（的首字母）在 s 中的位置
	// 特别地，后缀 s[sa[0]:] 字典序最小，后缀 s[sa[n-1]:] 字典序最大
	type _tp struct {
		_  []byte
		sa []int32
	}
	sa := (*_tp)(unsafe.Pointer(suffixarray.New([]byte(s)))).sa

	// 计算后缀名次数组
	// 后缀 s[i:] 位于后缀字典序中的第 rank[i] 个
	// 特别地，rank[0] 即 s 在后缀字典序中的排名，rank[n-1] 即 s[n-1:] 在字典序中的排名
	// 相当于 sa 的反函数，即 rank[sa[i]] = i
	rank := make([]int, len(sa))
	for i, p := range sa {
		rank[p] = i
	}

	// 计算高度数组（也叫 LCP 数组）
	// height[0] = 0（哨兵）
	// height[i] = LCP(s[sa[i]:], s[sa[i-1]:])  (i > 0)
	// 获取 s[i] 所在位置的高度：height[rank[i]]
	height := make([]int, len(sa))
	h := 0
	// 计算 s 与 s[sa[rank[0]-1]:] 的 LCP（记作 LCP0）
	// 计算 s[1:] 与 s[sa[rank[1]-1]:] 的 LCP（记作 LCP1）
	// 计算 s[2:] 与 s[sa[rank[2]-1]:] 的 LCP
	// ...
	// 计算 s[n-1:] 与 s[sa[rank[n-1]-1]:] 的 LCP
	// 从 LCP0 到 LCP1，我们只去掉了 s[0] 和 s[sa[rank[0]-1]] 这两个字符
	// 所以 LCP1 >= LCP0 - 1
	// 这样就能加快 LCP 的计算了（类似滑动窗口）
	// 注：实际只计算了 n-1 对 LCP，因为我们跳过了 rank[i] = 0 的情况
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < len(s) && j+h < len(s) && s[i+h] == s[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	st := newSparseTable(height, func(a int, b int) int { return min(a, b) })

	// 返回 LCP(s[i:], s[j:])，即两个后缀的最长公共前缀
	lcp := func(i, j int) int {
		if i == j {
			return len(sa) - i
		}
		// 将 s[i:] 和 s[j:] 通过 rank 数组映射为 height 的下标
		ri, rj := rank[i], rank[j]
		if ri > rj {
			ri, rj = rj, ri
		}
		// ri+1 是因为 height 的定义是 sa[i] 和 sa[i-1]
		// rj+1 是因为 query 是左闭右开
		return st.query(ri+1, rj+1)
	}
	return lcp
}

func almostPalindromic(s string) (ans int) {
	n := len(s)
	revS := []byte(s)
	slices.Reverse(revS)
	lcp := suffixArrayLCP(s + "#" + string(revS))

	// 将 s 改造为 t，这样就不需要分 len(s) 的奇偶来讨论了，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
	// s 和 t 的下标转换关系：
	// (si+1)*2 = ti
	// ti/2-1 = si
	// ti 为偶数（2,4,6,...）对应 s 中的奇回文串
	// ti 为奇数（3,5,7,...）对应 s 中的偶回文串
	t := append(make([]byte, 0, n*2+3), '^')
	for _, c := range s {
		t = append(t, '#', byte(c))
	}
	t = append(t, '#', '$')

	// 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
	// halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
	// 具体地，闭区间 [i-halfLen[i]+1, i+halfLen[i]-1] 是 t 上的一个回文子串
	// 由于 t 中回文子串的首尾字母一定是 #，根据下标转换关系，
	// 可以得到其在 s 中对应的回文子串的区间为 [(i-halfLen[i])/2, (i+halfLen[i])/2-2]
	halfLen := make([]int, len(t)-2)
	halfLen[1] = 1
	// boxR 表示当前右边界下标最大的回文子串的右边界下标+1（初始化成任意 <= 0 的数都可以）
	// boxM 为该最大回文子串的中心位置，二者的关系为 boxR = boxM + halfLen[boxM]
	boxM, boxR := 0, 0
	for i := 2; i < len(halfLen); i++ { // 循环的起止位置对应着原串的首尾字符
		hl := 1 // 注：如果题目比较的是抽象意义的值，单个值可能不满足要求，此时应初始化 hl = 0
		if i < boxR {
			// 记 i 关于 boxM 的对称位置 i'=boxM*2-i
			// 若以 i' 为中心的最长回文子串范围超出了以 boxM 为中心的回文串的范围（即 i+halfLen[i'] >= boxR）
			// 则 halfLen[i] 应先初始化为已知的回文半径 boxR-i，然后再继续暴力匹配
			// 否则 halfLen[i] 与 halfLen[i'] 相等
			hl = min(halfLen[boxM*2-i], boxR-i)
		}
		// 暴力扩展
		// 算法的复杂度取决于这部分执行的次数
		// 由于扩展之后 boxR 必然会更新（右移），且扩展的的次数就是 boxR 右移的次数
		// 因此算法的复杂度 = O(len(t)) = O(len(s))
		for t[i-hl] == t[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}
		halfLen[i] = hl

		// 闭区间 [(i-halfLen[i])/2, (i+halfLen[i])/2-2] 是 s 上的一个回文子串
		l, r := (i-halfLen[i])/2, (i+halfLen[i])/2-2

		// s 本身是回文串，或者删除两端一个字母是回文串
		if r-l+1 >= n-1 {
			return n // 如果 s 本身是回文串，删除回文中心后，仍然是回文串
		}

		// 删除 s[l-1]，继续扩展
		extra := 1 // 删除 [l,r] 外侧的一个字母
		if l-2 >= 0 && r+1 < n {
			extra += lcp(r+1, n*2-l+2) * 2
		}
		ans = max(ans, r-l+1+extra)

		// 删除 s[r+1]，继续扩展
		extra = 1 // 删除 [l,r] 外侧的一个字母
		if l-1 >= 0 && r+2 < n {
			extra += lcp(r+2, n*2-l+1) * 2
		}
		ans = max(ans, r-l+1+extra)
	}

	return
}
