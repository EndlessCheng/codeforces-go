package copypasta

// 模板题：CF101628K https://codeforces.ml/gym/101628/submission/68323182

type iTrieNode struct {
	sonIDs         [26]*iTrieNode
	fa             *iTrieNode
	curIndexes     *_treap // set
	subTreeIndexes *_treap // set
}

func (o *iTrieNode) pushUpPut(idx int) {
	for ; o.fa != nil; o = o.fa {
		o.subTreeIndexes.put(idx)
		// 其余统计量
	}
}

func (o *iTrieNode) pushUpDel(idx int) {
	for ; o.fa != nil; o = o.fa {
		o.subTreeIndexes.delete(idx)
		// 其余统计量
	}
}

type iTrie struct {
	root *iTrieNode
}

func newIndexTrie() *iTrie {
	return &iTrie{
		&iTrieNode{}, // init with a root (empty string)
	}
}

func (*iTrie) ord(c byte) byte { return c - 'a' }

func (t *iTrie) put(s []byte, idx int) {
	o := t.root
	for _, c := range s {
		c = t.ord(c)
		if o.sonIDs[c] == nil {
			o.sonIDs[c] = &iTrieNode{
				fa:             o,
				curIndexes:     &_treap{rd: 1},
				subTreeIndexes: &_treap{rd: 1},
			}
		}
		o = o.sonIDs[c]
	}
	o.curIndexes.put(idx)
	o.pushUpPut(idx)
}

// s 必须在 iTrie 中存在
func (t *iTrie) del(s []byte, idx int) {
	o := t.root
	for _, c := range s {
		o = o.sonIDs[t.ord(c)]
	}
	o.curIndexes.delete(idx)
	o.pushUpDel(idx)
}

// 在 trie 中找字符串 s 的前缀（这个前缀必须是 trie 中的一个完整字符串）
func (t *iTrie) hasPrefixOfText(s []byte, l, r int) bool {
	o := t.root
	for _, c := range s {
		o = o.sonIDs[t.ord(c)]
		if o == nil {
			return false
		}
		if o.curIndexes.hasValueInRange(l, r) {
			return true
		}
	}
	return false
}

// 在 trie 中找前缀为 p 的字符串
func (t *iTrie) hasTextOfPrefix(p []byte, l, r int) bool {
	o := t.root
	for _, c := range p {
		o = o.sonIDs[t.ord(c)]
		if o == nil {
			return false
		}
	}
	return o.subTreeIndexes.hasValueInRange(l, r)
}

// 占位符
type _treap struct {
	rd              uint
	put             func(int)
	delete          func(int)
	hasValueInRange func(int, int) bool
}
