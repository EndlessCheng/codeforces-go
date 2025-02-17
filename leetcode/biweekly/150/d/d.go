package main

import (
	"math"
	"strings"
)

// https://space.bilibili.com/206214
// 计算字符串 p 的 pi 数组
func calcPi(p string) []int {
	pi := make([]int, len(p))
	match := 0
	for i := 1; i < len(pi); i++ {
		v := p[i]
		for match > 0 && p[match] != v {
			match = pi[match-1]
		}
		if p[match] == v {
			match++
		}
		pi[i] = match
	}
	return pi
}

// 在文本串 s 中查找模式串 p，返回所有成功匹配的位置（p[0] 在 s 中的下标）
func kmpSearch(s, p string) (pos []int) {
	if p == "" {
		// s 的所有位置都能匹配空串，包括 len(s)
		pos = make([]int, len(s)+1)
		for i := range pos {
			pos[i] = i
		}
		return
	}
	pi := calcPi(p)
	match := 0
	for i := range s {
		v := s[i]
		for match > 0 && p[match] != v {
			match = pi[match-1]
		}
		if p[match] == v {
			match++
		}
		if match == len(p) {
			pos = append(pos, i-len(p)+1)
			match = pi[match-1]
		}
	}
	return
}

func shortestMatchingSubstring(s, p string) int {
	sp := strings.Split(p, "*")
	p1, p2, p3 := sp[0], sp[1], sp[2]

	// 三段各自在 s 中的所有匹配位置
	pos1 := kmpSearch(s, p1)
	pos2 := kmpSearch(s, p2)
	pos3 := kmpSearch(s, p3)

	ans := math.MaxInt
	i, k := 0, 0
	// 枚举中间（第二段），维护最近的左右（第一段和第三段）
	for _, j := range pos2 {
		// 右边找离 j 最近的子串（但不能重叠）
		for k < len(pos3) && pos3[k] < j+len(p2) {
			k++
		}
		if k == len(pos3) { // 右边没有
			break
		}
		// 左边找离 j 最近的子串（但不能重叠）
		for i < len(pos1) && pos1[i] <= j-len(p1) {
			i++
		}
		// 循环结束后，posL[i-1] 是左边离 j 最近的子串下标（首字母在 s 中的下标）
		if i > 0 {
			ans = min(ans, pos3[k]+len(p3)-pos1[i-1])
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
