package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func slidingWindow(s []byte) [][26]int {
	res := make([][26]int, len(s))

	// 计算至少有 k 种字母的子串个数
	for k := 1; k <= 26; k++ {
		cnt := [26]int{}
		kinds := 0
		left := 0
		for i, ch := range s {
			if cnt[ch-'a'] == 0 {
				kinds++
			}
			cnt[ch-'a']++
			for kinds >= k {
				cnt[s[left]-'a']--
				if cnt[s[left]-'a'] == 0 {
					kinds--
				}
				left++
			}
			res[i][k-1] = left // 右端点为 i 的合法子串个数
		}
		if left == 0 { // 没有 k 种字母，也没有 > k 种字母，不用再算
			break
		}
	}

	for i := range res {
		// (至少 j 种) - (至少 j+1 种) = (恰好 j 种)
		for j := range 25 {
			res[i][j] -= res[i][j+1]
		}
	}

	return res
}

func countMagicalPairs1(S string) (ans int64) {
	s := []byte(S)
	pre := slidingWindow(s)
	slices.Reverse(s)
	suf := slidingWindow(s) // 相当于倒着滑窗
	slices.Reverse(suf)

	for i := range len(pre) - 1 {
		for j := 0; j < 26 && pre[i][j] > 0 && suf[i+1][j] > 0; j++ {
			ans += int64(pre[i][j]) * int64(suf[i+1][j])
		}
	}
	return
}

// 模板来源：https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/
func logTrick(s []byte) [][26]int {
	// 子串包含的字母集合，最小左端点
	type pair struct{ mask, left int }
	data := []pair{}
	res := make([][26]int, len(s))

	for i, ch := range s {
		bit := 1 << (ch - 'a')
		// 计算右端点为 i 的子串的 mask
		for j := range data {
			data[j].mask |= bit
		}
		// ch 单独作为子串
		data = append(data, pair{bit, i})

		// 原地去重（相同 mask 只保留最左边的）
		idx := 1
		for j := 1; j < len(data); j++ {
			if data[j].mask != data[j-1].mask {
				data[idx] = data[j]
				idx++
			}
		}
		data = data[:idx]

		for k, p := range data {
			right := i
			if k < len(data)-1 {
				right = data[k+1].left - 1
			}
			// 对于左端点在 [p.left, right]，右端点为 i 的子串，mask 都相同
			res[i][len(data)-1-k] = right - p.left + 1
		}
	}

	return res
}

func countMagicalPairs2(S string) (ans int64) {
	s := []byte(S)
	pre := logTrick(s)
	slices.Reverse(s)
	suf := logTrick(s) // 相当于倒着遍历 s
	slices.Reverse(suf)

	for i := range len(pre) - 1 {
		for j := 0; j < 26 && pre[i][j] > 0 && suf[i+1][j] > 0; j++ {
			ans += int64(pre[i][j]) * int64(suf[i+1][j])
		}
	}
	return
}

//

func countMagicalPairs(s string) (ans int64) {
	n := len(s)
	// nextIdx[i] 表示下一个 s[i] 的位置
	nextIdx := make([]int, n)
	last := [26]int{}
	for i := range last {
		last[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		ch := s[i] - 'a'
		nextIdx[i] = last[ch]
		last[ch] = i
	}

	// 右侧最近不同字母的位置
	right := []int{}
	for ch, p := range last {
		if p < n {
			right = append(right, p)
			last[ch] = -1
		}
	}
	slices.Sort(right)
	right = append(right, n) // 哨兵

	// 左侧最近不同字母的位置
	left := []int{-1} // 哨兵

	// 枚举分割线 (i,i+1)
	for i := range n - 1 {
		ch := s[i] - 'a'
		// 更新左侧最近不同字母的位置
		idx := last[ch]
		if idx >= 0 {
			j := sort.SearchInts(left, idx)
			left = append(left[:j], left[j+1:]...) // 删除旧的
		}
		left = append(left, i) // 添加新的
		last[ch] = i

		// 更新右侧最近不同字母的位置
		right = right[1:] // 删除旧的
		idx = nextIdx[i]
		if idx < n {
			j := sort.SearchInts(right, idx)
			right = slices.Insert(right, j, idx) // 添加新的
		}

		nl := len(left)
		for k := range min(nl, len(right)) - 1 {
			ans += int64(left[nl-1-k]-left[nl-2-k]) * int64(right[k+1]-right[k])
		}
	}

	return
}
