package main

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"math"
	"strings"
)

// https://space.bilibili.com/206214
func minOperations(s string, k int) int {
	n := len(s)
	z := strings.Count(s, "0")
	if z == 0 {
		return 0
	}
	if k == n {
		if z == n {
			return 1
		}
		return -1
	}

	ans := math.MaxInt
	// 情况一：操作次数 m 是偶数
	if z%2 == 0 { // z 必须是偶数
		m := max((z+k-1)/k, (z+n-k-1)/(n-k)) // 下界
		ans = m + m%2                        // 把 m 往上调整为偶数
	}

	// 情况二：操作次数 m 是奇数
	if z%2 == k%2 { // z 和 k 的奇偶性必须相同
		m := max((z+k-1)/k, (n-z+n-k-1)/(n-k)) // 下界
		ans = min(ans, m|1)                    // 把 m 往上调整为奇数
	}

	if ans < math.MaxInt {
		return ans
	}
	return -1
}

func minOperations1(s string, k int) (ans int) {
	n := len(s)
	notVis := [2]*redblacktree.Tree[int, struct{}]{}
	for m := range notVis {
		notVis[m] = redblacktree.New[int, struct{}]()
		for i := m; i <= n; i += 2 {
			notVis[m].Put(i, struct{}{})
		}
		notVis[m].Put(n+1, struct{}{}) // 哨兵，下面无需判断 node != nil
	}

	start := strings.Count(s, "0")
	notVis[start%2].Remove(start)
	q := []int{start}
	for q != nil {
		tmp := q
		q = nil
		for _, z := range tmp {
			if z == 0 { // 没有 0，翻转完毕
				return ans
			}
			// notVis[mn % 2] 中的从 mn 到 mx 都可以从 z 翻转到
			mn := z + k - 2*min(k, z)
			mx := z + k - 2*max(0, k-n+z)
			t := notVis[mn%2]
			for node, _ := t.Ceiling(mn); node.Key <= mx; node, _ = t.Ceiling(mn) {
				q = append(q, node.Key)
				t.Remove(node.Key)
			}
		}
		ans++
	}
	return -1
}
