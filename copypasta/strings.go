package copypasta

import (
	"index/suffixarray"
	"reflect"
	"unsafe"
)

// TIPS: 若处理原串比较困难，不妨考虑下反转后的串 https://codeforces.ml/contest/873/problem/F

func stringCollection() {
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	// 字符串哈希
	// https://oi-wiki.org/string/hash/
	// 利用 set 可以求出固定长度的不同子串个数
	// 模板题 https://www.luogu.com.cn/problem/P3370
	// 最长重复子串（二分哈希）https://leetcode-cn.com/problems/longest-duplicate-substring/
	// 题目推荐 https://cp-algorithms.com/string/string-hashing.html#toc-tgt-7
	// TODO 二维 hash
	// TODO anti-hash: 最好不要自然溢出 https://codeforces.ml/blog/entry/4898 https://codeforces.ml/blog/entry/60442
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
		cnt := 0
		for i := 1; i < n; i++ {
			b := s[i]
			for cnt > 0 && s[cnt] != b {
				cnt = maxMatchLengths[cnt-1]
			}
			if s[cnt] == b {
				cnt++
			}
			maxMatchLengths[i] = cnt
		}
		return maxMatchLengths
	}
	// search pattern from text, return all start positions
	kmpSearch := func(text, pattern string) (pos []int) {
		maxMatchLengths := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		cnt := 0
		for i := range text {
			b := text[i]
			for cnt > 0 && pattern[cnt] != b {
				cnt = maxMatchLengths[cnt-1]
			}
			if pattern[cnt] == b {
				cnt++
			}
			if cnt == lenP {
				pos = append(pos, i-lenP+1)
				cnt = maxMatchLengths[cnt-1]
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
	// https://codeforces.com/blog/entry/12143
	// http://manacher-viz.s3-website-us-east-1.amazonaws.com
	// https://oi-wiki.org/string/manacher/#manacher
	// https://cp-algorithms.com/string/manacher.html
	// 模板题 https://www.luogu.com.cn/problem/P3805
	var maxLen []int
	manacher := func(origin []byte) int {
		min := func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}
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
				// 取 min 的原因：记点 i 关于 mid 的对称点为 i'，
				// 若以 i' 为中心的回文串范围超过了以 mid 为中心的回文串的范围
				// (此时有 i + len[2*mid-i] >= right，注意 len 是包括中心的半长度)
				// 则 len[i] 应取 right - i (总不能超过边界吧)
				maxLen[i] = min(maxLen[2*mid-i], right-i)
			} else {
				maxLen[i] = 1
			}
			for s[i+maxLen[i]] == s[i-maxLen[i]] {
				maxLen[i]++
			}
			mx := maxLen[i]
			if ans < mx {
				ans = mx
			}
			if right < i+mx {
				mid = i
				right = i + mx
			}
		}
		return ans - 1
	}
	// 判断 [l,r] 是否为回文串，范围 0<=l<=r<n
	isP := func(l, r int) bool { return maxLen[l+r+2] >= r-l+1 }

	// 后缀数组
	// 讲解+例题+套题 https://oi-wiki.org/string/sa/
	// 由于 height 数组的性质，经常需要和单调栈/单调队列结合
	// SA-IS 与 DC3 的效率对比 https://riteme.site/blog/2016-6-19/sais.html#5
	// NOTE: Go1.13 开始使用 SA-IS 算法
	// 题目推荐 https://cp-algorithms.com/string/suffix-array.html#toc-tgt-11
	// 题目总结《后缀数组——处理字符串的有力工具》
	// 模板题 https://www.luogu.com.cn/problem/P3809
	// 可重叠最长重复子串 https://leetcode-cn.com/problems/longest-duplicate-substring/
	// 不可重叠最长重复子串 http://poj.org/problem?id=1743（可参考《算法与实现》p.223 以及 https://oi-wiki.org/string/sa/#_14）
	// 可重叠的至少出现 k 次的最长重复子串 http://poj.org/problem?id=3261（height 上的滑动窗口最小值）
	// 重复次数最多的连续重复子串 http://poj.org/problem?id=3693
	// todo []int 的后缀数组
	suffixArray := func(s []byte) {
		n := len(s)
		//sa := *(*[]int)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").UnsafeAddr()))
		sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))

		rank := make([]int, n) // rank[i] 表示 s[i:] 在 sa 中的位置
		for i := range rank {
			rank[sa[i]] = i
		}
		height := make([]int, n) // height[i] = lcp(s[sa[i]:], s[sa[i-1]:])
		h := 0
		for i, sai := range rank {
			if h > 0 {
				h--
			}
			if sai > 0 {
				for j := int(sa[sai-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
				}
			}
			height[sai] = h
		}

		// 构建 height 的 ST 表……
		lcp := func(i, j int) (res int) {
			ri, rj := rank[sa[i]], rank[sa[j]]
			if ri > rj {
				ri, rj = rj, ri
			}
			//res := stQuery(ri+1, rj) // 左闭右闭
			return
		}

		longestDupSubstring := func() []byte {
			maxP, maxH := 0, 0
			for i, h := range height {
				if h > maxH {
					maxP, maxH = i, h
				}
			}
			return s[sa[maxP] : int(sa[maxP])+maxH]
		}

		// debug
		for i, h := range height {
			suffix := string(s[sa[i]:])
			if h == 0 {
				println(" ", suffix)
			} else {
				println(h, suffix)
			}
		}

		_ = []interface{}{lcp, longestDupSubstring}
	}

	_ = []interface{}{
		initPowP, calcHash,
		kmpSearch, calcMinPeriod,
		smallestRepresentation,
		manacher, isP,
		suffixArray,
	}
}

// 前缀树/字典树/单词查找树
// 另类解读：如果将字符串长度视作定值的话，trie 树是一种 O(n) 排序，O(1) 查询的数据结构
//          这点上和哈希表很像，但是 trie 树可以在路径上保存信息，从而能做到一些哈希表做不到的前缀操作
// https://oi-wiki.org/string/trie/
// https://www.quora.com/q/threadsiiithyderabad/Tutorial-on-Trie-and-example-problems
// 另见 strings_index_trie.go
// NOTE: 为保证连续性，分隔符可取 'Z'+1 或 'z'+1
// 模板题 https://leetcode-cn.com/problems/implement-trie-prefix-tree/
// 前缀和后缀搜索 https://leetcode-cn.com/problems/prefix-and-suffix-search/
// 回文对（配合 Manacher 可以做到线性复杂度） https://leetcode-cn.com/problems/palindrome-pairs/
// LC 套题（推荐困难难度的题） https://leetcode-cn.com/tag/trie/
// todo https://codeforces.ml/contest/455/problem/B
type trieNode struct {
	son    [26]*trieNode
	dupCnt int
	val    int // val 也可以是个 []int，此时 dupCnt == len(val)

	// 当 o.son[i] 不能匹配 text 中的某个字符时，o.fail 即为下一个应该查找的结点
	fail *trieNode
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
		bits[i] = byte(val >> (30 - i) & 1)
	}

	o := t.root
	for i, b := range bits {
		if o.son[b^1] != nil {
			ans |= 1 << (30 - i)
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
			prefixes[v>>i] = true
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

// EXTRA: Aho–Corasick algorithm
// https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
// 推荐 https://zhuanlan.zhihu.com/p/80325757
// 推荐 https://www.cnblogs.com/nullzx/p/7499397.html
// https://oi-wiki.org/string/ac-automaton/
// EXTRA: https://cp-algorithms.com/string/aho_corasick.html
// 模板题 https://leetcode-cn.com/problems/stream-of-characters/
// 模板题 https://www.luogu.com.cn/problem/P3808 https://www.luogu.com.cn/problem/P3796
// https://codeforces.com/problemset/problem/963/D
func (t *trie) buildDFA() {
	q := []*trieNode{}
	for _, son := range t.root.son {
		if son != nil {
			q = append(q, son)
			son.fail = t.root
		}
	}
	for len(q) > 0 {
		var o *trieNode
		o, q = q[0], q[1:]
		for i, son := range o.son {
			if son != nil {
				q = append(q, son)
				// 沿着失配边不断往上查找，直到找到一个匹配的前缀。未找到时指向根节点
				for f := o.fail; ; f = f.fail {
					if f == nil {
						son.fail = t.root
						break
					}
					if fs := f.son[i]; fs != nil {
						son.fail = fs
						break
					}
				}
			}
		}
	}
}

// 返回 text 中所有模式串的所有位置（未找到时对应数组为空）
// patterns 为模式串数组（无重复元素），为方便起见，数组从 1 开始
// TODO 后缀链接优化：只算出现次数可以做到 O(len(text))
func (t *trie) acSearch(text string, patterns []string) [][]int {
	pos := make([][]int, len(patterns))
	o := t.root
	for i := range text {
		c := t.ord(text[i])
		for ; o != t.root && o.son[c] == nil; o = o.fail {
		}
		o = o.son[c]
		if o == nil {
			o = t.root
		}
		for f := o; f != t.root; f = f.fail {
			if pid := f.val; pid != 0 {
				pos[pid] = append(pos[pid], i-len(patterns[pid])+1)
			}
		}
	}
	return pos
}

// 可持久化 trie
// TODO https://oi-wiki.org/ds/persistent-trie/
// 模板题（最大异或和） https://www.luogu.com.cn/problem/P4735

// Suffix automaton (SAM)
// https://en.wikipedia.org/wiki/Suffix_automaton
// todo https://codeforces.com/blog/entry/20861
// TODO https://oi-wiki.org/string/sam/
// TODO https://cp-algorithms.com/string/suffix-automaton.html
//《后缀自动机》，陈立杰
//《后缀自动机在字典树上的拓展》，刘研绎
//《后缀自动机及其应用》，张天扬
// 模板题 https://www.luogu.com.cn/problem/P3804
