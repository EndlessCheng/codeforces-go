package main

import (
	"index/suffixarray"
	"math"
	"slices"
	"unsafe"
)

// https://space.bilibili.com/206214
func calc(s, t string) int {
	// ts = t + "#" + s
	ts := append([]byte(t), '#')
	tmp := []byte(s)
	slices.Reverse(tmp)
	ts = append(ts, tmp...)
	sa := (*struct {
		_  []byte
		sa []int32
	})(unsafe.Pointer(suffixarray.New(ts))).sa

	// 后缀名次数组 rank
	// 后缀 ts[i:] 位于后缀字典序中的第 rank[i] 个
	// 特别地，rank[0] 即 ts 在后缀字典序中的排名，rank[n-1] 即 ts[n-1:] 在字典序中的排名
	rank := make([]int, len(sa))
	for i, p := range sa {
		rank[p] = i
	}

	// 高度数组 height
	// sa 中相邻后缀的最长公共前缀 LCP
	// height[0] = height[len(sa)] = 0（哨兵）
	// height[i] = LCP(ts[sa[i]:], ts[sa[i-1]:])
	height := make([]int, len(sa)+1)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < len(ts) && j+h < len(ts) && ts[i+h] == ts[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	mx := make([]int, len(s)+1)
	lcp := 0
	// sa[0] 对应 '#' 开头的后缀，不遍历
	for i := 1; i < len(sa); i++ {
		if int(sa[i]) < len(t) {
			lcp = math.MaxInt // 找到了 t 中的后缀，可以开始计算 LCP
		} else {
			lcp = min(lcp, height[i])
			mx[int(sa[i])-len(t)-1] = lcp
		}
	}
	lcp = 0
	for i := len(sa) - 1; i > 0; i-- { // 反着再来一遍
		if int(sa[i]) < len(t) {
			lcp = math.MaxInt
		} else {
			lcp = min(lcp, height[i+1])
			j := int(sa[i]) - len(t) - 1
			mx[j] = max(mx[j], lcp)
		}
	}
	ans := slices.Max(mx) * 2 // |x| = |y| 的情况
	slices.Reverse(mx)

	// 计算 |x| > |y| 的情况
	s2 := append(make([]byte, 0, len(s)*2+3), '^')
	for _, c := range s {
		s2 = append(s2, '#', byte(c))
	}
	s2 = append(s2, '#', '$')
	halfLen := make([]int, len(s2)-2)
	halfLen[1] = 1
	boxM, boxR := 0, 0
	for i := 2; i < len(halfLen); i++ {
		hl := 1
		if i < boxR {
			hl = min(halfLen[boxM*2-i], boxR-i)
		}
		for s2[i-hl] == s2[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}
		halfLen[i] = hl

		if hl > 1 { // 回文子串不为空
			l := (i - hl) / 2 // 回文子串左端点
			ans = max(ans, hl-1+mx[l]*2)
		}
	}
	return ans
}

func longestPalindrome(s, t string) int {
	return max(calc(s, t), calc(reverse(t), reverse(s)))
}

func reverse(s string) string {
	t := []byte(s)
	slices.Reverse(t)
	return string(t)
}
