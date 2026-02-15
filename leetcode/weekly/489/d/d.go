package main

// https://space.bilibili.com/206214
const width = 15 // nums[i] 二进制长度的最大值

type node struct {
	son  [2]*node
	leaf int // 子树叶子个数
}

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{&node{}}
}

func (t *trie) put(val int) {
	cur := t.root
	for i := width - 1; i >= 0; i-- {
		bit := val >> i & 1
		if cur.son[bit] == nil {
			cur.son[bit] = &node{}
		}
		cur = cur.son[bit]
		cur.leaf++
	}
}

func (t *trie) del(val int) {
	cur := t.root
	for i := width - 1; i >= 0; i-- {
		cur = cur.son[val>>i&1]
		cur.leaf-- // 如果减成 0 了，说明子树是空的，可以理解成 cur == nil
	}
}

func (t *trie) maxXor(val int) (ans int) {
	cur := t.root
	for i := width - 1; i >= 0; i-- {
		bit := val >> i & 1
		if cur.son[bit^1] != nil && cur.son[bit^1].leaf > 0 {
			ans |= 1 << i
			bit ^= 1
		}
		cur = cur.son[bit]
	}
	return
}

func maxXor(nums []int, k int) (ans int) {
	sum := make([]int, len(nums)+1)
	for i, x := range nums {
		sum[i+1] = sum[i] ^ x
	}

	t := newTrie()
	var minQ, maxQ []int
	left := 0
	for right, x := range nums {
		// 1. 入
		t.put(sum[right])

		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, right)

		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, right)

		// 2. 出
		for nums[maxQ[0]]-nums[minQ[0]] > k {
			t.del(sum[left])
			left++
			if minQ[0] < left {
				minQ = minQ[1:]
			}
			if maxQ[0] < left {
				maxQ = maxQ[1:]
			}
		}

		// 3. 更新答案
		ans = max(ans, t.maxXor(sum[right+1]))
	}
	return
}
