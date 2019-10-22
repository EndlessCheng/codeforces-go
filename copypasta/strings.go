package copypasta

import (
	. "fmt"
	"index/suffixarray"
	"reflect"
	"unsafe"
)

func hashCollection() {
	const prime uint64 = 1e8 + 7

	var n int
	powP := make([]uint64, n+1)
	powP[0] = 1
	for i := 1; i <= n; i++ {
		powP[i] = powP[i-1] * prime
	}
}

func stringCollection() {
	// code from my answer at https://www.zhihu.com/question/21923021/answer/37475572
	calcMaxMatchLengths := func(pattern string) []int {
		n := len(pattern)
		maxMatchLengths := make([]int, n)
		maxLength := 0
		for i := 1; i < n; i++ {
			c := pattern[i]
			for maxLength > 0 && pattern[maxLength] != c {
				maxLength = maxMatchLengths[maxLength-1]
			}
			if pattern[maxLength] == c {
				maxLength++
			}
			maxMatchLengths[i] = maxLength
		}
		return maxMatchLengths
	}
	// search pattern from text, return all start positions
	kmpSearch := func(text, pattern string) (positions []int) {
		maxMatchLengths := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		count := 0
		for i := range text {
			c := text[i]
			for count > 0 && pattern[count] != c {
				count = maxMatchLengths[count-1]
			}
			if pattern[count] == c {
				count++
			}
			if count == lenP {
				positions = append(positions, i-lenP+1)
				count = maxMatchLengths[count-1]
			}
		}
		return
	}

	calcMinPeriod := func(pattern string) int {
		maxMatchLengths := calcMaxMatchLengths(pattern)
		n := len(pattern)
		if val := maxMatchLengths[n-1]; val > 0 {
			if n%(n-val) == 0 {
				return n / (n - val)
			}
		}
		return 1 // or -1
	}

	_ = []interface{}{kmpSearch, calcMinPeriod}
}

func suffixArray() {
	// lcp[i] = lcp(s[sa[i]:], s[sa[i+1]:])
	calcLCP := func(s []byte, sa []int) (lcp []int) {
		n := len(s)
		rank := make([]int, n+1)
		for i := range rank {
			rank[sa[i]] = i
		}
		lcp = make([]int, n, n+1)
		h := 0
		for i := range lcp {
			j := sa[rank[i]-1]
			if h > 0 {
				h--
			}
			for ; j+h < n && i+h < n; h++ {
				if s[j+h] != s[i+h] {
					break
				}
			}
			lcp[rank[i]-1] = h
		}
		return
	}

	var s []byte
	index := suffixarray.New(s)
	sa := *(*[]int)(unsafe.Pointer(reflect.ValueOf(index).Elem().FieldByName("sa").UnsafeAddr()))
	sa = append([]int{len(s)}, sa...) // 方便定义 lcp
	lcp := calcLCP(s, sa)
	lcp = append(lcp, 0)

	// debug
	for i := range sa {
		if lcp[i] == 0 {
			Println("  " + string(s[sa[i]:]))
		} else {
			Println(lcp[i], string(s[sa[i]:]))
		}
	}

	// TODO: []int 的后缀数组
}

//

type trieNode struct {
	childIdx [26]int
	dupCnt   int // 重复插入计数
	val      int // 节点附加信息
}
type trie struct {
	nodes []*trieNode
}

func newTrie() *trie {
	return &trie{
		nodes: []*trieNode{{}},
	}
}

// insert `s` into trie and add `val` at leaf
func (t *trie) insert(s string, val int) {
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
	o.val = val
}
