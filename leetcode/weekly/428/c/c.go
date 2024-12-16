package main

import (
	"index/suffixarray"
	"math/bits"
	"unsafe"
)

// https://space.bilibili.com/206214
func calcZ(s []int) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0 // z-box 左右边界
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
	}
	return z
}

func beautifulSplits(nums []int) (ans int) {
	n := len(nums)
	z0 := calcZ(nums)
	for i := 1; i < n-1; i++ {
		z := calcZ(nums[i:])
		for j := i + 1; j < n; j++ {
			if i <= j-i && z0[i] >= i || z[j-i] >= j-i {
				ans++
			}
		}
	}
	return
}

func beautifulSplits2(nums []int) (ans int) {
	n := len(nums)
	// lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
	lcp := make([][]int, n+1)
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			if nums[i] == nums[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}

	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if i <= j-i && lcp[0][i] >= i || lcp[i][j] >= j-i {
				ans++
			}
		}
	}
	return
}

func beautifulSplits3(a []int) (ans int) {
	s := []byte{}
	for _, v := range a {
		s = append(s, byte(v))
	}
	type _tp struct {
		_  []byte
		sa []int32
	}
	sa := (*_tp)(unsafe.Pointer(suffixarray.New(s))).sa

	rank := make([]int, len(sa))
	for i, p := range sa {
		rank[p] = i
	}
	height := make([]int, len(sa))
	h := 0
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

	logN := bits.Len(uint(len(sa)))
	st := make([][]int, len(sa))
	for i, v := range height {
		st[i] = make([]int, logN)
		st[i][0] = v
	}
	for j := 1; 1<<j <= len(sa); j++ {
		for i := 0; i+1<<j <= len(sa); i++ {
			st[i][j] = min(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	_q := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return min(st[l][k], st[r-1<<k][k]) }
	lcp := func(i, j int) int {
		if i == j {
			return len(sa) - i
		}
		ri, rj := rank[i], rank[j]
		if ri > rj {
			ri, rj = rj, ri
		}
		return _q(ri+1, rj+1)
	}
	isPrefix := func(l1, r1, l2, r2 int) bool {
		len1, len2 := r1-l1, r2-l2
		return len1 <= len2 && lcp(l1, l2) >= len1
	}

	n := len(s)
	for l := 1; l < n-1; l++ {
		for r := l + 1; r < n; r++ {
			if isPrefix(0, l, l, r) || isPrefix(l, r, r, n) {
				ans++
			}
		}
	}
	return
}
