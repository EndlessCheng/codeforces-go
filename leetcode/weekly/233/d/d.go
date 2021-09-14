package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func fwt(a []int) {
	n := len(a)
	for l, k := 2, 1; l <= n; l, k = l<<1, k<<1 {
		for i := 0; i < n; i += l {
			for j := 0; j < k; j++ {
				a[i+j], a[i+j+k] = a[i+j]+a[i+j+k], a[i+j]-a[i+j+k]
			}
		}
	}
}

func ifwt(a []int) {
	n := len(a)
	for l, k := 2, 1; l <= n; l, k = l<<1, k<<1 {
		for i := 0; i < n; i += l {
			for j := 0; j < k; j++ {
				a[i+j], a[i+j+k] = (a[i+j]+a[i+j+k])/2, (a[i+j]-a[i+j+k])/2
			}
		}
	}
}

func countPairs(nums []int, low, high int) int {
	mx := 0
	for _, v := range nums {
		if v > mx {
			mx = v
		}
	}
	a := make([]int, 1<<bits.Len(uint(mx)))
	for _, v := range nums {
		a[v]++
	}
	fwt(a)
	for i, v := range a {
		a[i] *= v
	}
	ifwt(a)
	ans := 0
	for _, v := range a[low:min(high+1, len(a))] {
		ans += v
	}
	return ans / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//

// 普通解法
type trieNode struct {
	son [2]*trieNode
	cnt int
}

type trie struct{ root *trieNode }

func newTrie() *trie { return &trie{&trieNode{}} }

const trieBitLen = 14

func (t *trie) put(v int) *trieNode {
	o := t.root
	for i := trieBitLen; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
		o.cnt++
	}
	return o
}

func (t *trie) countLimitXOR(v, limit int) (cnt int) {
	o := t.root
	for i := trieBitLen; i >= 0; i-- {
		b := v >> i & 1
		if limit>>i&1 > 0 {
			if o.son[b] != nil {
				cnt += o.son[b].cnt
			}
			b ^= 1
		}
		if o.son[b] == nil {
			return
		}
		o = o.son[b]
	}
	return
}

func countPairs2(a []int, low, high int) (ans int) {
	t := newTrie()
	t.put(a[0])
	for _, v := range a[1:] {
		ans += t.countLimitXOR(v, high+1) - t.countLimitXOR(v, low)
		t.put(v)
	}
	return
}
