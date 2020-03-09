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
	// TODO 二维 hash
	// TODO 建议随机 hash 防 hack
	var powP []uint64
	initPowP := func(maxLen int) {
		const prime uint64 = 1e8 + 7
		powP = make([]uint64, maxLen)
		powP[0] = 1
		for i := 1; i < maxLen; i++ {
			powP[i] = powP[i-1] * prime
		}
	}
	calcHash := func(s []byte) (val uint64) {
		for i, c := range s {
			val += uint64(c) * powP[i]
		}
		return
	}

	// https://oi-wiki.org/string/kmp/
	// TODO https://oi-wiki.org/string/z-func/
	// https://cp-algorithms.com/string/prefix-function.html
	// code from my answer at https://www.zhihu.com/question/21923021/answer/37475572
	calcMaxMatchLengths := func(s string) []int {
		n := len(s)
		maxMatchLengths := make([]int, n)
		maxLength := 0
		for i := 1; i < n; i++ {
			c := s[i]
			for maxLength > 0 && s[maxLength] != c {
				maxLength = maxMatchLengths[maxLength-1]
			}
			if s[maxLength] == c {
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
	// EXTRA: 最小循环节
	calcMinPeriod := func(s string) int {
		n := len(s)
		maxMatchLengths := calcMaxMatchLengths(s)
		if val := maxMatchLengths[n-1]; val > 0 && n%(n-val) == 0 {
			return n / (n - val)
		}
		return 1 // 无小于 n 的循环节
	}

	// TODO 扩展 KMP
	// 模板题 https://www.luogu.com.cn/problem/P5410

	// 最小表示法
	// TODO：待整理
	// https://oi-wiki.org/string/minimal-string/
	smallestRepresentation := func(s []byte) []byte {
		n := len(s)
		s = append(s, s...) // 或者 copy
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

	// TODO：待整理
	// https://blog.csdn.net/synapse7/article/details/18908413
	// http://manacher-viz.s3-website-us-east-1.amazonaws.com
	// https://oi-wiki.org/string/manacher/#manacher
	// https://cp-algorithms.com/string/manacher.html
	// 模板题 https://www.luogu.com.cn/problem/P3805
	var maxLen []int
	manacher := func(origin []byte) int {
		n := len(origin)
		s := make([]byte, 2*n+3)
		s[0] = '^'
		for i, c := range origin {
			s[2*i+1] = '#'
			s[2*i+2] = c
		}
		s[2*n+1] = '#'
		s[2*n+2] = '$'
		maxLen = make([]int, 2*n+3)
		var ans, mid, right int
		for i := 1; i < 2*n+2; i++ {
			if i < right {
				maxLen[i] = min(maxLen[2*mid-i], right-i)
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

// 前缀树/字典树/单词查找树
// 另类解读：如果将字符串长度视作定值的话，trie 树是一种 O(n) 排序，O(1) 查询的数据结构
//          这点上和哈希表很像，但是 trie 树可以在路径上保存信息，从而能做到一些哈希表做不到的前缀操作
// https://oi-wiki.org/string/trie/
// 另见 strings_index_trie.go
// NOTE: 为保证连续性，分隔符可取 'Z'+1 或 'z'+1
// 模板题 https://leetcode-cn.com/problems/implement-trie-prefix-tree/
// 前缀和后缀搜索 https://leetcode-cn.com/problems/prefix-and-suffix-search/
// 回文对（配合 Manacher 可以做到线性复杂度） https://leetcode-cn.com/problems/palindrome-pairs/
// LC 套题（推荐困难难度的题） https://leetcode-cn.com/tag/trie/
type trieNode struct {
	son    [26]*trieNode // 2
	dupCnt int
	val    int
	// val 也可以是个 string、[]int 或 map，此时 dupCnt == len(val)
}

func (o *trieNode) empty() bool {
	for _, son := range o.son {
		if son != nil {
			return false
		}
	}
	return true
}

type trie struct{ root *trieNode }

func newTrie() *trie {
	// init with a root (empty string)
	return &trie{&trieNode{}}
}

func (*trie) ord(c byte) byte { return c - 'a' }
func (*trie) chr(v byte) byte { return v + 'a' }

func (t *trie) put(s []byte, val int) {
	o := t.root
	for _, c := range s {
		c = t.ord(c)
		if o.son[c] == nil {
			o.son[c] = &trieNode{}
		}
		o = o.son[c]
		//o.dupCnt++ // 经过节点 o 的字符串个数
		//o.val = val // 更新 s 的所有前缀的值
	}
	o.dupCnt++
	o.val = val
}

func (t *trie) get(s []byte) *trieNode {
	o := t.root
	for _, c := range s {
		o = o.son[t.ord(c)]
		if o == nil {
			return nil
		}
	}
	if o.dupCnt == 0 {
		return nil
	} // s 只是某个字符串的前缀
	return o
}

// s 必须在 trie 中存在
func (t *trie) del(s []byte) {
	fa := make([]*trieNode, len(s)+1)
	o := t.root
	for i, c := range s {
		fa[i] = o
		o = o.son[t.ord(c)]
		//o.dupCnt--
	}
	o.dupCnt--
	if o.dupCnt == 0 {
		for i := len(s) - 1; i >= 0; i-- {
			f := fa[i]
			f.son[t.ord(s[i])] = nil
			if !f.empty() {
				break
			}
		}
	}
}

// 在 trie 中寻找字典序最小的以 p 为前缀的字符串
// 若没有，返回 nil, 0
func (t *trie) minPrefix(p []byte) (s []byte, node *trieNode) {
	o := t.root
	for _, c := range p {
		o = o.son[t.ord(c)]
		if o == nil {
			return
		}
	}
	// trie 中存在字符串 s，使得 p 是 s 的前缀

	for o.dupCnt == 0 {
		for i, son := range o.son {
			if son != nil {
				s = append(s, t.chr(byte(i)))
				o = son
				break
			}
		}
	}
	return s, o
}

// rank 和 kth 分别求小于 s 的字符串个数和第 k 小
// 此时 o.val 保存子树字符串个数
func (t *trie) rank(s []byte) (k int) {
	o := t.root
	for _, c := range s {
		c = t.ord(c)
		for _, son := range o.son[:c] {
			if son != nil {
				k += son.val
			}
		}
		o = o.son[c]
		if o == nil {
			return
		}
	}
	//k += o.val // 加上这句表示小于等于 s 的字符串个数
	return
}

func (t *trie) kth(k int) (s []byte) {
	o := t.root
outer:
	for {
		for i, son := range o.son {
			if son != nil {
				if k >= son.val {
					k -= son.val
				} else {
					o = son
					s = append(s, t.chr(byte(i)))
					continue outer
				}
			}
		}
		return
	}
}

// 01-trie：val 与树上所有数中的最大异或值
// 也可以说这是一颗（所有叶节点深度都相同的）二叉树
// 参考《算法竞赛进阶指南》0x16
// 模板题：数组中两个数的最大异或值 https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/
// 模板题：树上最长异或路径 https://www.luogu.com.cn/problem/P4551
func (t *trie) maxXor(val int) (ans int) {
	bits := [31]byte{}
	for i := range bits {
		bits[i] = byte(val >> uint(30-i) & 1)
	}

	o := t.root
	for i, b := range bits {
		if o.son[b^1] != nil {
			ans |= 1 << uint(30-i)
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 也可以用哈希表做，效率是一样的
func findMaximumXOR(a []int) (ans int) {
	n := len(a)
	for i := 30; i >= 0; i-- {
		ans <<= 1
		prefixes := make(map[int]bool, n)
		for _, v := range a {
			prefixes[v>>uint(i)] = true
		}
		tmp := ans + 1
		for p := range prefixes {
			if prefixes[tmp^p] {
				ans = tmp
				break
			}
		}
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
