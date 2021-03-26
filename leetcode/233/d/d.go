package main

// github.com/EndlessCheng/codeforces-go
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

func countPairs(a []int, low, high int) (ans int) {
	t := newTrie()
	t.put(a[0])
	for _, v := range a[1:] {
		ans += t.countLimitXOR(v, high+1) - t.countLimitXOR(v, low)
		t.put(v)
	}
	return
}
