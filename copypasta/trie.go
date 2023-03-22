package copypasta

/* 前缀树/字典树/单词查找树
另类解读：如果将字符串长度视作定值的话，trie 树是一种 O(n) 排序，O(1) 查询的数据结构
https://oi-wiki.org/string/trie/
https://www.quora.com/q/threadsiiithyderabad/Tutorial-on-Trie-and-example-problems
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TrieST.java.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TrieSET.java.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TST.java.html

注：由于用的是指针写法，必要时禁止 GC，能加速不少
func init() { debug.SetGCPercent(-1) }

模板题 LC208 https://leetcode-cn.com/problems/implement-trie-prefix-tree/
前缀和后缀搜索 周赛62D/LC745 https://leetcode-cn.com/problems/prefix-and-suffix-search/
https://codeforces.com/contest/514/problem/C
回文对（配合 Manacher 可以做到线性复杂度）LC336 https://leetcode-cn.com/problems/palindrome-pairs/
与 DP 结合 https://leetcode-cn.com/problems/re-space-lcci/
与贪心堆结合 https://codeforces.com/problemset/problem/965/E
todo https://codeforces.com/contest/455/problem/B

深刻理解 https://atcoder.jp/contests/abc273/tasks/abc273_e
*/
type trieNode struct {
	son [26]*trieNode
	cnt int //（子树中）完整字符串的个数
	val int // []int

	isEnd bool

	// AC 自动机：当 o.son[i] 不能匹配文本串 text 中的某个字符时，o.fail 即为下一个应该查找的结点
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

func (trie) ord(c byte) byte { return c - 'a' }
func (trie) chr(v byte) byte { return v + 'a' }

// 插入字符串 s，附带值 val，返回插入后字符串末尾对应的节点
func (t *trie) put(s []byte, val int) *trieNode {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
		//o.cnt++ // 写法一：统计 o 对应的字符串是多少个完整字符串的前缀
	}
	o.cnt++ // 写法二：统计 o 上有多少个完整字符串
	o.val = val
	return o
}

// 字典树 DFS（模板）
// LC2416 https://leetcode.cn/problems/sum-of-prefix-scores-of-strings/
func (t *trie) dfs() {
	var f func(*trieNode, int)
	f = func(o *trieNode, sum int) {
		if o == nil {
			return
		}
		// 统计从根到 o 的路径
		sum += o.cnt //

		for _, child := range o.son {
			f(child, sum)
		}
	}
	f(t.root, 0)
}

// 最长连续单词链
// https://leetcode.com/discuss/interview-question/2255835/Google-or-Onsite-or-Longest-Chain-Words
func (t *trie) longestChainWords() (ans int) {
	var f func(*trieNode, int)
	f = func(o *trieNode, cnt int) {
		if o == nil {
			return
		}
		if o.isEnd {
			cnt++
			if cnt > ans {
				ans = cnt
			}
		} else {
			cnt = 0
		}
		for _, child := range o.son {
			f(child, cnt)
		}
	}
	f(t.root, 0)
	return
}

// 查找字符串 s
func (t *trie) find(s []byte) *trieNode {
	o := t.root
	for _, b := range s {
		o = o.son[t.ord(b)]
		// 未找到 s，且 s 不是任何字符串的前缀
		if o == nil {
			return nil
		}
		//sum += o.cnt
	}
	// 未找到 s，但是 s 是某个字符串的前缀
	if o.cnt == 0 { // 已删除
		return nil
	}
	return o
}

// 删除字符串 s，返回字符串末尾对应的节点
// LC1804 https://leetcode-cn.com/problems/implement-trie-ii-prefix-tree/
func (t *trie) delete(s []byte) *trieNode {
	fa := make([]*trieNode, len(s))
	o := t.root
	for i, b := range s {
		fa[i] = o
		o = o.son[t.ord(b)]
		if o == nil {
			return nil
		}
		//o.cnt-- // 对应 put 的写法
	}
	o.cnt--
	if o.cnt == 0 {
		for i := len(s) - 1; i >= 0; i-- {
			f := fa[i]
			f.son[t.ord(s[i])] = nil
			if !f.empty() {
				break
			}
		}
	}
	return o
}

// 求小于 s 的字符串个数
// 此时 o.cnt 保存子树完整字符串个数
func (t *trie) rank(s []byte) (k int) {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		for _, son := range o.son[:b] {
			if son != nil {
				k += son.cnt
			}
		}
		o = o.son[b]
		if o == nil {
			return
		}
	}
	//k += o.cnt // 这样写就是小于或等于 s 的字符串个数
	return
}

// 求第 k 小（k 从 0 开始，相当于有 k 个字符串小于返回的字符串 s）
// 此时 o.cnt 保存子树完整字符串个数
func (t *trie) kth(k int) (s []byte) {
	o := t.root
outer:
	for {
		for i, son := range o.son[:] {
			if son != nil {
				if k < son.cnt {
					o = son
					s = append(s, t.chr(byte(i)))
					continue outer
				}
				k -= son.cnt
			}
		}
		return
	}
}

// 结合 rank 和 kth，可以求出一个字符串的前驱和后继
// 见 bst.go 中的 prev 和 next

// 返回字符串 s 在 trie 中的前缀个数
// https://www.acwing.com/problem/content/144/
// https://codeforces.com/gym/101628/problem/K
func (t *trie) countPrefixOfString(s []byte) (cnt int) {
	o := t.root
	for _, b := range s {
		o = o.son[t.ord(b)]
		if o == nil {
			return
		}
		cnt += o.cnt
	}
	return
}

// 返回 trie 中前缀为 p 的字符串个数
// 此时 o.cnt 保存子树字符串个数
// https://codeforces.com/gym/101628/problem/K
// LC1804 https://leetcode-cn.com/problems/implement-trie-ii-prefix-tree/
func (t *trie) countStringHasPrefix(p []byte) int {
	o := t.root
	for _, b := range p {
		o = o.son[t.ord(b)]
		if o == nil {
			return 0
		}
	}
	return o.cnt
}

// s 的本质不同子串数量 O(n^2)
// 做法是插入每个后缀，统计节点数。但题目往往会带上额外的条件
// https://codeforces.com/problemset/problem/271/D
//     注：这题还可以用后缀数组+前缀和二分来做到 O(nlogn)
func (t *trie) countDistinctSubstring(s []byte) (cnt int) {
	for i := range s {
		o := t.root
		for _, b := range s[i:] {
			b = t.ord(b)
			if o.son[b] == nil {
				o.son[b] = &trieNode{}
				cnt++
			}
			o = o.son[b]
		}
	}
	return
}

// EXTRA: AC 自动机 Aho–Corasick algorithm / Deterministic Finite Automaton (DFA)
// https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
// https://en.wikipedia.org/wiki/Deterministic_finite_automaton
// 基础实现 https://zhuanlan.zhihu.com/p/80325757
// 基础实现 https://www.cnblogs.com/nullzx/p/7499397.html
// 改进实现 https://oi-wiki.org/string/ac-automaton/
// 应用 https://cp-algorithms.com/string/aho_corasick.html
//
// 模板题
// LC1032 https://leetcode-cn.com/problems/stream-of-characters/
// https://www.luogu.com.cn/problem/P3808
// https://www.luogu.com.cn/problem/P3796
// todo 最长前缀查询 https://www.luogu.com.cn/problem/P5231
// todo https://www.luogu.com.cn/problem/P5357 二次加强版
//  NOI11 阿狸的打字机 https://www.luogu.com.cn/problem/P2414
//  https://www.acwing.com/solution/content/25473/
//  https://www.acwing.com/solution/content/54646/
//
// todo https://codeforces.com/problemset/problem/1437/G
// todo https://codeforces.com/problemset/problem/963/D
// todo LC30 串联所有单词的子串 https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/
// todo ? LC616 给字符串添加加粗标签 https://leetcode-cn.com/problems/add-bold-tag-in-string/
func (t *trie) buildDFA() {
	q := []*trieNode{}
	for _, son := range t.root.son[:] {
		if son != nil {
			son.fail = t.root
			q = append(q, son)
		}
	}
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		if o.fail == nil {
			o.fail = t.root
		}
		for i, son := range o.son[:] {
			if son != nil {
				son.fail = o.fail.son[i]
				q = append(q, son)
			} else {
				o.son[i] = o.fail.son[i]
			}
		}
	}
}

// 有多少个（编号）不同的模式串在文本串 text 里出现过
func (t *trie) sumCountAllPatterns(text []byte) (cnt int) {
	o := t.root
	for _, b := range text {
		o = o.son[t.ord(b)]
		if o == nil {
			o = t.root
			continue
		}
		for f := o; f != nil && f.val > 0; f = f.fail {
			cnt += f.val
			f.val = 0
		}
	}
	return
}

// 返回所有模式串 patterns 的开头在文本串 text 的所有位置（未找到时对应数组为空）
// patterns 为模式串数组（无重复元素），为方便起见，patterns 从 1 开始
func (t *trie) acSearch(text []byte, patterns [][]byte) [][]int {
	pos := make([][]int, len(patterns))
	o := t.root
	for i, b := range text {
		o = o.son[t.ord(b)]
		if o == nil {
			o = t.root
			continue
		}
		for f := o; f != nil; f = f.fail {
			if pid := f.val; pid != 0 {
				pos[pid] = append(pos[pid], i-len(patterns[pid])+1) // 也可以只记录 i，代表模式串末尾在文本的位置
			}
		}
	}
	return pos
}

// EXTRA: 可持久化字典树
// 注意为了拷贝一份 trieNode，这里的接收器不是指针
// https://oi-wiki.org/ds/persistent-trie/
// roots := make([]*trieNode, n+1)
// roots[0] = &trieNode{}
// roots[i+1] = roots[i].put(s)
func (o trieNode) put(s []byte) *trieNode {
	if len(s) == 0 {
		o.cnt++
		return &o
	}
	b := s[0] - 'a' //
	if o.son[b] == nil {
		o.son[b] = &trieNode{}
	}
	o.son[b] = o.son[b].put(s[1:])
	//o.maintain()
	return &o
}

// EXTRA: 回文自动机（回文树） PAM  Eertree
// todo https://oi-wiki.org/string/pam/
//  https://baobaobear.github.io/post/20200416-pam/
//  《字符串算法选讲》-金策
//  https://zhuanlan.zhihu.com/p/92874690
//  https://arxiv.org/pdf/1506.04862v2.pdf
//
// 模板题 https://www.luogu.com.cn/problem/P5496
// todo 本质不同回文子串个数 https://hihocoder.com/problemset/problem/1602
//  回文子串出现次数 https://www.luogu.com.cn/problem/P3649
//  最小回文划分 https://codeforces.com/problemset/problem/932/G
//  能否划分成三段回文 LC1745 https://leetcode-cn.com/problems/palindrome-partitioning-iv/
//  最长双回文串（另一种做法是 Manacher）https://www.luogu.com.cn/problem/P4555
