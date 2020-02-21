package copypasta

import (
	. "fmt"
	"index/suffixarray"
	"reflect"
	"unsafe"
)

func stringCollection() {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	// https://oi-wiki.org/string/hash/
	// 模板题 https://www.luogu.com.cn/problem/P3370
	// 题目推荐 https://cp-algorithms.com/string/string-hashing.html#toc-tgt-7
	// TODO: 二维 hash
	var powP []uint64
	initPowP := func(maxLen int) {
		const prime uint64 = 1e8 + 7
		powP = make([]uint64, maxLen+1)
		powP[0] = 1
		for i := 1; i <= maxLen; i++ {
			powP[i] = powP[i-1] * prime
		}
	}
	calcHash := func(s string) (val uint64) {
		for i := range s {
			val += uint64(s[i]) * powP[i]
		}
		return
	}

	// https://oi-wiki.org/string/kmp/
	// TODO https://oi-wiki.org/string/z-func/
	// https://cp-algorithms.com/string/prefix-function.html
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
		if val := maxMatchLengths[n-1]; val > 0 && n%(n-val) == 0 {
			return n / (n - val)
		}
		return 1 // or -1
	}

	// TODO 扩展 KMP
	// 模板题 https://www.luogu.com.cn/problem/P5410

	// 最小表示法
	// https://oi-wiki.org/string/minimal-string/
	smallestRepresentation := func(s string) string {
		n := len(s)
		s += s
		i := 0
		for j := 1; j < n; {
			k := 0
			for ; k < n && s[i+k] == s[j+k]; k++ {
			}
			if k >= n {
				break
			}
			if s[i+k] < s[j+k] {
				j += k + 1
			} else {
				i, j = j, max(j, i+k)+1
			}
		}
		return s[i : i+n]
	}

	// https://blog.csdn.net/synapse7/article/details/18908413
	// http://manacher-viz.s3-website-us-east-1.amazonaws.com
	// https://oi-wiki.org/string/manacher/#manacher
	// https://cp-algorithms.com/string/manacher.html
	// 模板题 https://www.luogu.com.cn/problem/P3805
	var maxLen []int
	manacher := func(origin string) int {
		n := len(origin)
		s := make([]byte, 2*n+3)
		s[0] = '^'
		for i := range origin {
			s[i<<1|1] = '#'
			s[i<<1+2] = origin[i]
		}
		s[n<<1|1] = '#'
		s[n<<1+2] = '$'
		maxLen = make([]int, 2*n+3)
		var ans, mid, right int
		for i := 1; i < 2*n+2; i++ {
			if i < right {
				maxLen[i] = min(maxLen[mid<<1-i], right-i)
			} else {
				maxLen[i] = 1
			}
			// 取 min 的原因：记点 i 关于 mid 的对称点为 i'，
			// 若以 i' 为中心的回文串范围超过了以 mid 为中心的回文串的范围
			// (此时有 i + len[(mid<<1)-i] >= right，注意 len 是包括中心的半长度)
			// 则 len[i] 应取 right - i (总不能超过边界吧)
			for s[i+maxLen[i]] == s[i-maxLen[i]] {
				maxLen[i]++
			}
			ans = max(ans, maxLen[i])
			if right < i+maxLen[i] {
				mid = i
				right = i + maxLen[i]
			}
		}
		return ans - 1
	}

	// 判断源串中的某一子串 [l...r] 是否为回文串
	manacherQuery := func(l, r int) bool { return maxLen[l+r+2] >= r-l+1 }

	_ = []interface{}{
		initPowP, calcHash,
		kmpSearch, calcMinPeriod,
		smallestRepresentation,
		manacher, manacherQuery,
	}
}

// https://oi-wiki.org/string/sa/#height
// 题目推荐 https://cp-algorithms.com/string/suffix-array.html#toc-tgt-11
// 模板题 https://www.luogu.com.cn/problem/P3809
func suffixArrayCollection() {
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
			if h > 0 {
				h--
			}
			for j := sa[rank[i]-1]; j+h < n && i+h < n && s[j+h] == s[i+h]; h++ {
			}
			lcp[rank[i]-1] = h
		}
		return
	}

	var s []byte
	sa := *(*[]int)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").UnsafeAddr()))
	// TODO: 感觉要再整理一下
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

// 字典树
// https://oi-wiki.org/string/trie/
// 另见 strings_index_trie.go
// 题目推荐 https://codeforces.ml/blog/entry/55274
type trieNode struct {
	sonIDs [26]int
	dupCnt int // 重复节点数
	val    int // 节点附加信息（比如插入的字符串在原数组中的下标）
	// val 也可以是个 []int 或 map，此时 dupCnt == len(val)
}

type trie struct {
	nodes []*trieNode
}

func newTrie() *trie {
	return &trie{
		nodes: []*trieNode{{}}, // init with a root (empty string)
	}
}

func (*trie) ord(c byte) byte { return c - 'a' }
func (*trie) chr(v byte) byte { return v + 'a' }

func (t *trie) add(s []byte, val int) {
	o := t.nodes[0]
	for _, c := range s {
		c = t.ord(c)
		if o.sonIDs[c] == 0 {
			o.sonIDs[c] = len(t.nodes)
			t.nodes = append(t.nodes, &trieNode{})
		}
		o = t.nodes[o.sonIDs[c]]
		//o.dupCnt++ // 写在循环内部表示经过节点 o 的字符串个数
	}
	o.dupCnt++
	//if o.dupCnt == 1 {
	o.val = val
	//}
	//o.val = append(o.val, val)
}

// s 必须在 trie 中存在
func (t *trie) del(s []byte) {
	o := t.nodes[0]
	for _, c := range s {
		o = t.nodes[o.sonIDs[t.ord(c)]]
		//o.dupCnt--
	}
	o.dupCnt--
}

func (t *trie) get(s []byte) *trieNode {
	o := t.nodes[0]
	for _, c := range s {
		id := o.sonIDs[t.ord(c)]
		if id == 0 {
			return nil
		}
		o = t.nodes[id]
	}
	if o.dupCnt == 0 {
		return nil
	} // s 只是某个字符串的前缀
	return o
}

// 在 trie 中寻找字典序最小的以 p 为前缀的字符串
// 若没有，返回 nil, 0
func (t *trie) minPrefix(p []byte) (s []byte, node *trieNode) {
	o := t.nodes[0]
	for _, c := range p {
		id := o.sonIDs[t.ord(c)]
		if id == 0 {
			return
		}
		o = t.nodes[id]
	}
	// trie 中存在字符串 s，使得 p 是 s 的前缀

	for o.dupCnt == 0 {
		for i, id := range o.sonIDs {
			if id > 0 {
				s = append(s, t.chr(byte(i)))
				o = t.nodes[id]
				break
			}
		}
	}
	return s, o
}

// 01-trie
// childIdx 长度为 2，且 trie 上所有字符串长度与 bits 一致 (31)
// 参考《算法竞赛进阶指南》0x16
// 模板题：树上最长异或路径 https://www.luogu.com.cn/problem/P4551
func (t *trie) maxXor(val int) (ans int) {
	bits := [31]byte{}
	for i := range bits {
		bits[i] = byte(val >> uint(30-i) & 1)
	}

	o := t.nodes[0]
	for i, b := range bits {
		if o.sonIDs[b^1] > 0 {
			ans |= 1 << uint(30-i)
			b ^= 1
		}
		o = t.nodes[o.sonIDs[b]]
	}
	return
}

// 可持久化 trie
// TODO https://oi-wiki.org/ds/persistent-trie/
// 模板题（最大异或和） https://www.luogu.com.cn/problem/P4735

// Aho–Corasick algorithm
// https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
// TODO https://oi-wiki.org/string/ac-automaton/
// TODO https://cp-algorithms.com/string/aho_corasick.html

// Suffix automaton
// https://en.wikipedia.org/wiki/Suffix_automaton
// TODO https://oi-wiki.org/string/sam/
// TODO https://cp-algorithms.com/string/suffix-automaton.html
