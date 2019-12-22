package main

type trieNode struct {
	childIdx [26]int
	cnt      int
}

type trie struct {
	nodes []*trieNode
}

func newTrie() *trie {
	return &trie{
		nodes: []*trieNode{{}},
	}
}

func (t *trie) put(s []byte) {
	o := t.nodes[0]
	for _, c := range s {
		if o.childIdx[c] == 0 {
			o.childIdx[c] = len(t.nodes)
			t.nodes = append(t.nodes, &trieNode{})
		}
		o = t.nodes[o.childIdx[c]]
		o.cnt++
	}
}

func (t *trie) get(s []byte) int {
	o := t.nodes[0]
	for _, c := range s {
		idx := o.childIdx[c]
		if idx == 0 {
			return 0
		}
		o = t.nodes[idx]
	}
	return o.cnt
}

func maxFreq(ss string, maxLetters int, minSize int, maxSize int) (ans int) {
	s := []byte(ss)
	for i := range s {
		s[i] -= 'a'
	}

	n := len(s)
	t := newTrie()
	i := 0
	for ; i+26 <= n; i++ {
		t.put(s[i : i+26])
	}
	for ; i < n; i++ {
		t.put(s[i:n])
	}

	valid := func(substr []byte) bool {
		mp := map[byte]bool{}
		for _, c := range substr {
			mp[c] = true
		}
		return len(mp) <= maxLetters
	}
	for subLen := minSize; subLen <= maxSize; subLen++ {
		for i := 0; i+subLen <= n; i++ {
			if substr := s[i : i+subLen]; valid(substr) {
				if cnt := t.get(substr); cnt > ans {
					ans = cnt
				}
			}
		}
	}
	return
}
