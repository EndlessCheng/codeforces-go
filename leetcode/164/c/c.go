package main

import (
	"sort"
	"strings"
)

type trieNode struct {
	childIdx [26]int
	dupCnt   int
	val      int
}

type trie struct {
	nodes []*trieNode
}

func newTrie() *trie {
	return &trie{
		nodes: []*trieNode{{}},
	}
}

func (t *trie) put(s string, val int) {
	o := t.nodes[0]
	for _, c := range s {
		c -= 'a'
		if o.childIdx[c] == 0 {
			o.childIdx[c] = len(t.nodes)
			t.nodes = append(t.nodes, &trieNode{})
		}
		o = t.nodes[o.childIdx[c]]
	}
	o.dupCnt++
	if o.dupCnt == 1 {
		o.val = val
	}
}

func (t *trie) minPrefix(p string) int {
	o := t.nodes[0]
	for _, c := range p {
		idx := o.childIdx[c-'a']
		if idx == 0 {
			return -1
		}
		o = t.nodes[idx]
	}
	for o.dupCnt == 0 {
		for i := 0; i < 26; i++ {
			if idx := o.childIdx[i]; idx > 0 {
				o = t.nodes[idx]
				break
			}
		}
	}
	return o.val
}

func suggestedProducts(products []string, searchWord string) (ans [][]string) {
	t := newTrie()
	sort.Strings(products)
	for i, p := range products {
		t.put(p, i)
	}
	for i := 1; i <= len(searchWord); i++ {
		sug := []string{}
		prefix := searchWord[:i]
		idx := t.minPrefix(prefix)
		if idx == -1 {
			ans = append(ans, sug)
			continue
		}
		for j := idx; j < len(products); j++ {
			p := products[j]
			if !strings.HasPrefix(p, prefix) {
				break
			}
			sug = append(sug, p)
			if len(sug) == 3 {
				break
			}
		}
		ans = append(ans, sug)
	}
	return
}
