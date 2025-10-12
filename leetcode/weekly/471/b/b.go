package main

import (
	"bytes"
	"slices"
)

// https://space.bilibili.com/206214
func longestBalanced1(s string) (ans int) {
	for i := range s {
		cnt := make([]int, 26)
	next:
		for j := i; j < len(s); j++ {
			cnt[s[j]-'a']++
			base := cnt[s[j]-'a']
			for _, c := range cnt {
				if c > 0 && c != base {
					continue next
				}
			}
			ans = max(ans, j-i+1)
		}
	}
	return
}

func longestBalanced2(s string) (ans int) {
	for i := range s {
		cnt := [26]int{}
		mx, kinds := 0, 0
		for j := i; j < len(s); j++ {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				kinds++
			}
			cnt[b]++
			mx = max(mx, cnt[b])
			if mx*kinds == j-i+1 {
				ans = max(ans, j-i+1)
			}
		}
	}
	return
}

func longestBalanced(s string) (ans int) {
	n := len(s)
	sufOrders := make([][]byte, n)
	order := []byte{}
	move := func(b byte) {
		// 把最近出现的字母 b 移到 order 末尾
		j := bytes.IndexByte(order, b)
		if j >= 0 {
			order = append(order[:j], order[j+1:]...)
		}
		order = append(order, b)
	}
	for i := n - 1; i >= 0; i-- {
		move(s[i] - 'a')
		sufOrders[i] = slices.Clone(order)
	}

	order = []byte{}
	cnt := [27]int{} // cnt[26] 作为 mask，用来区分 tmp[ch] 是差值还是原始值
	pos := map[[27]int]int{}
	for i, b := range s {
		sufOrder := sufOrders[i]
		minCh := byte(25)
		cnt[26] = 0
		for j := len(sufOrder) - 1; j >= 0; j-- {
			cnt[26] |= 1 << sufOrder[j]
			minCh = min(minCh, sufOrder[j])
			// 注意此时 cnt 并不包含 s[i]，我们计算的是前缀 s[:i] 的信息
			// 在子串中的字母，计算差值
			// 不在子串中的字母，维持原样
			d := cnt
			for _, ch := range sufOrder[j:] {
				d[ch] -= cnt[minCh]
			}
			// 记录 d 首次出现的位置
			if _, ok := pos[d]; !ok {
				pos[d] = i - 1
			}
		}

		// 把最近出现的字母移到 order 末尾
		move(byte(b - 'a'))

		cnt[b-'a']++
		minCh = byte(25)
		cnt[26] = 0
		for j := len(order) - 1; j >= 0; j-- {
			cnt[26] |= 1 << order[j]
			minCh = min(minCh, order[j])
			d := cnt
			for _, ch := range order[j:] {
				d[ch] -= cnt[minCh]
			}
			// 再次遇到完全一样的状态，说明找到了一个平衡子串，左端点为 l+1，右端点为 i
			if l, ok := pos[d]; ok {
				ans = max(ans, i-l)
			}
		}
	}
	return
}
