package main

import "sort"

// github.com/EndlessCheng/codeforces-go
const bitLen = 30

type node struct{ son [2]*node }
type trie struct{ root *node }

func bin(v int) []byte {
	s := make([]byte, bitLen+1)
	for i := range s {
		s[i] = byte(v >> (bitLen - i) & 1)
	}
	return s
}

func (t *trie) put(v int) *node {
	o := t.root
	for _, b := range bin(v) {
		if o.son[b] == nil {
			o.son[b] = &node{}
		}
		o = o.son[b]
	}
	return o
}

func (t *trie) maxXor(v int) (ans int) { // TEMPLATE
	o := t.root
	for i, b := range bin(v) {
		if o.son[b^1] != nil {
			ans |= 1 << (bitLen - i)
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

func maximizeXor(a []int, qs [][]int) (ans []int) {
	ans = make([]int, len(qs))
	sort.Ints(a)
	for i := range qs {
		qs[i] = append(qs[i], i)
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i][1] < qs[j][1] })

	t, i := &trie{&node{}}, 0
	for _, q := range qs {
		for ; i < len(a) && a[i] <= q[1]; i++ {
			t.put(a[i])
		}
		if i == 0 {
			ans[q[2]] = -1
		} else {
			ans[q[2]] = t.maxXor(q[0])
		}
	}
	return
}
