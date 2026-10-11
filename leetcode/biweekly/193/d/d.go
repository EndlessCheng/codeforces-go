package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
// 计算 a 在 b 左侧时，a 和 b 之间的逆序对个数
// a 和 b 已经升序排序
func calcInversions(a, b []int) (inv int) {
	n, m := len(a), len(b)
	i, j := 0, 0
	for i < n && j < m {
		if a[i] > b[j] {
			// a 的子数组 [i, n-1] 中的数都大于 b[j] 
			inv += n - i
			j++
		} else {
			i++
		}
	}
	return
}

func minInversions(words []string, target []int) (ans int) {
	n := len(target)
	pos := make([]int, n)
	for i, x := range target {
		pos[x] = i
	}

	// 创建前缀树（字典树）
	type node struct {
		son [10]*node
		pos int
	}
	root := &node{}
	for i, word := range words {
		cur := root
		for _, ch := range word {
			ch -= 'a'
			if cur.son[ch] == nil {
				cur.son[ch] = &node{}
			}
			cur = cur.son[ch]
		}
		// 对于不是 word 的节点，pos 是 0
		cur.pos = pos[i] + 1 // +1，与 pos = 0 的节点区分开
	}

	// 返回 node 子树的 pos 列表（已从小到大排序）
	var dfs func(*node) []int
	dfs = func(node *node) []int {
		a := [][]int{}
		for _, son := range node.son {
			if son != nil {
				a = append(a, dfs(son))
			}
		}

		// 预处理逆序对个数
		m := len(a)
		inv := make([][]int, m)
		for i, p := range a {
			inv[i] = make([]int, m)
			for j, q := range a {
				if j != i {
					inv[i][j] = calcInversions(p, q)
				}
			}
		}

		u := 1 << m
		sumInv := make([][]int, m)
		for j := range sumInv {
			sum := make([]int, u)
			for i := range m {
				highBit := 1 << i
				for mask, s := range sum[:highBit] {
					sum[highBit|mask] = s + inv[i][j]
				}
			}
			sumInv[j] = sum
		}

		f := make([]int, u)
		for i := 1; i < u; i++ {
			f[i] = math.MaxInt
		}
		for s, fs := range f { // 之前选的下标集合是 s
			// 在 s 的右边添加 j
			for j := range m {
				if s>>j&1 == 0 {
					f[s|1<<j] = min(f[s|1<<j], fs+sumInv[j][s])
				}
			}
		}
		ans += f[u-1]

		res := []int{}
		curPos := node.pos
		if curPos > 0 {
			res = append(res, curPos)
		}
		for _, sub := range a {
			if curPos > 0 {
				// curPos 与子树之间的逆序对个数
				ans += sort.SearchInts(sub, curPos)
			}
			res = append(res, sub...)
		}

		slices.Sort(res)
		return res
	}
	dfs(root)

	return
}
