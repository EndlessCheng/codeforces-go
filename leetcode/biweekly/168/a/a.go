package main

import (
	"index/suffixarray"
	"math/bits"
	"slices"
	"unsafe"
)

// https://space.bilibili.com/206214
func lexSmallest1(s string) string {
	n := len(s)
	ans := s // k = 1 时，操作不改变 s
	for k := 2; k <= n; k++ {
		t := []byte(s[:k])
		slices.Reverse(t)
		ans = min(ans, string(t)+s[k:])

		t = []byte(s[n-k:])
		slices.Reverse(t)
		ans = min(ans, s[:n-k]+string(t))
	}
	return ans
}

//

type sparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

func newSparseTable[T any](nums []T, op func(T, T) T) sparseTable[T] {
	n := len(nums)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], nums)
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = op(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return sparseTable[T]{st, op}
}

// [l, r) 下标从 0 开始
func (s sparseTable[T]) query(l, r int) T {
	k := bits.Len(uint(r-l)) - 1
	return s.op(s.st[k][l], s.st[k][r-1<<k])
}

func lexSmallest(s string) string {
	n := len(s)
	t := []byte(s)
	slices.Reverse(t)
	t = append(t, '#')
	t = append(t, s...)

	// === 后缀数组模板开始 ===

	// 后缀数组 sa（后缀序）
	// sa[i] 表示后缀字典序中的第 i 个字符串（的首字母）在 s 中的位置
	// 特别地，后缀 s[sa[0]:] 字典序最小，后缀 s[sa[n-1]:] 字典序最大
	sa := (*struct {
		_  []byte
		sa []int32
	})(unsafe.Pointer(suffixarray.New(t))).sa

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
	// height[i] = LCP(s[sa[i]:], s[sa[i-1]:])   (i > 0)
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
			for j := int(sa[rk-1]); i+h < len(t) && j+h < len(t) && t[i+h] == t[j+h]; h++ {
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

	// 比较两个子串，返回 strings.Compare(s[l1:r1], s[l2:r2])
	compareSubstring := func(l1, r1, l2, r2 int) int {
		len1, len2 := r1-l1, r2-l2
		l := lcp(l1, l2)
		if l >= min(len1, len2) {
			// 一个是子串另一个子串的前缀，或者完全相等
			return len1 - len2
		}
		// 此时两个子串一定不相等
		return rank[l1] - rank[l2] // 也可以写 int(s[l1+l]) - int(s[l2+l])
	}

	// === 后缀数组模板结束 ===

	// 反转前缀
	ansK := 1
	for k := 2; k <= n; k++ {
		c := compareSubstring(n-k, n-k+ansK, n-ansK, n)
		if c < 0 || c == 0 && compareSubstring(n-k+ansK, n, n+1+ansK, n+1+k) < 0 {
			ansK = k
		}
	}
	pre := []byte(s[:ansK])
	slices.Reverse(pre)
	ans := string(pre) + s[ansK:]

	// 反转真后缀
	// 剪枝：如果 s[0] > ans[0]，那么反转真后缀一定不优
	if s[0] == ans[0] {
		ansK = 1
		for k := 2; k < n; k++ {
			c := compareSubstring(0, k-ansK, n*2+1-k, n*2+1-ansK)
			if c < 0 || c == 0 && compareSubstring(k-ansK, k, 0, ansK) < 0 {
				ansK = k
			}
		}
		suf := []byte(s[n-ansK:])
		slices.Reverse(suf)
		ans = min(ans, s[:n-ansK]+string(suf))
	}
	return ans
}
