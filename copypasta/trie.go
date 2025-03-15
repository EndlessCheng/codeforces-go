package copypasta

/* 前缀树/字典树/单词查找树
适用于多串前缀/后缀匹配
另类解读：如果将字符串长度视作定值 L 的话，trie 树是一种 O(nL) 排序，O(L) 查询的数据结构
https://oi-wiki.org/string/trie/
https://www.quora.com/q/threadsiiithyderabad/Tutorial-on-Trie-and-example-problems
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TrieST.java.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TrieSET.java.html
https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TST.java.html

注：由于用的是指针写法，必要时禁止 GC，能加速不少
func init() { debug.SetGCPercent(-1) }

模板题 LC208 https://leetcode.cn/problems/implement-trie-prefix-tree/
最长匹配后缀 https://leetcode.cn/problems/longest-common-suffix-queries/
前后缀同时匹配 LC745 https://leetcode.cn/problems/prefix-and-suffix-search/
            LC3045 https://leetcode.cn/problems/count-prefix-and-suffix-pairs-ii/
- 把 (s[i], s[n-1-i]) 插入字典树
LC527 https://leetcode.cn/problems/word-abbreviation/
LC336 https://leetcode.cn/problems/palindrome-pairs/ 回文对（配合 Manacher 可以做到线性复杂度）
https://leetcode.cn/problems/re-space-lcci/ 与 DP 结合
https://codeforces.com/problemset/problem/271/D 1800 满足一些条件的本质不同子串个数
https://codeforces.com/problemset/problem/455/B 1900
https://codeforces.com/problemset/problem/113/B 2000 满足一些条件的本质不同子串个数
https://codeforces.com/problemset/problem/514/C 2000
https://codeforces.com/problemset/problem/965/E 2200 与贪心堆结合
https://codeforces.com/problemset/problem/557/E 2300
https://atcoder.jp/contests/abc273/tasks/abc273_e 深刻理解
https://atcoder.jp/contests/abc353/tasks/abc353_e
*/
type trieNode struct {
	son [26]*trieNode
	cnt int // trieNode 对应的完整字符串的个数
	sum int // 子树 cnt 之和
	val int // 额外存储的信息
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
	return &trie{&trieNode{}}
}

func (trie) ord(c rune) rune { return c - 'a' }
func (trie) chr(v byte) byte { return v + 'a' }

// 插入字符串 s，附带值 val，返回插入后字符串末尾对应的节点
func (t *trie) put(s string, val int) *trieNode {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
		o.sum++ // 子树 cnt 之和（o 对应的字符串是多少个完整字符串的前缀）
	}
	o.cnt++ // o 对应的完整字符串的个数
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
		// 统计从 root 到 o 的路径
		sum += o.cnt
		// ...
		for _, child := range o.son {
			f(child, sum)
		}
	}
	f(t.root, 0)
}

// 字符串 s 与字典树中字符串的最长公共前缀
// 返回最后一个匹配的节点（最长公共前缀），以及是否找到 s
func (t *trie) find(s string) (*trieNode, bool) {
	o := t.root
	for _, b := range s {
		nxt := o.son[t.ord(b)]
		if nxt == nil {
			return o, false
		}
		o = nxt
	}
	return o, o.cnt != 0
}

// 删除字符串 s，返回字符串末尾对应的节点
// LC1804 https://leetcode.cn/problems/implement-trie-ii-prefix-tree/
func (t *trie) delete(s string) *trieNode {
	fa := make([]*trieNode, len(s))
	o := t.root
	for i, b := range s {
		fa[i] = o
		o = o.son[t.ord(b)]
		if o == nil {
			return nil
		}
	}
	o.cnt--
	if o.cnt == 0 {
		for i := len(s) - 1; i >= 0; i-- {
			f := fa[i]
			f.son[t.ord(rune(s[i]))] = nil // 完全删除节点
			if !f.empty() {
				break
			}
		}
	}
	return o
}

// 求小于 s 的字符串个数
func (t *trie) rank(s string) (k int) {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		// 累加在 b 之前的子树大小
		for _, son := range o.son[:b] {
			if son != nil {
				k += son.sum
			}
		}
		o = o.son[b]
		if o == nil {
			return
		}
		k += o.cnt // 以 b 结尾的字符串个数
	}
	// 上面算的是小于等于 s 的字符串个数
	// 要算小于 s 的字符串个数，要把恰好等于 s 的字符串个数减掉
	k -= o.cnt
	return
}

// 求第 k 小（k 从 0 开始）
// 相当于 <= s 的字符串至少有 k+1 个
// ！需要保证 trie 中至少有 k+1 个字符串
// https://codeforces.com/problemset/problem/557/E
func (t *trie) kth(k int) (s []byte) {
	o := t.root
outer:
	for {
		for i, son := range o.son {
			if son == nil {
				continue
			}
			// 子树 son 中的字符串都比答案小
			if k >= son.sum {
				k -= son.sum
				continue
			}
			s = append(s, 'a'+byte(i))
			k -= son.cnt
			if k < 0 {
				return
			}
			o = son
			continue outer
		}
		panic("k is too large (make sure k starts from 0)")
	}
}

// 结合 rank 和 kth，可以求出一个字符串的前驱和后继
// 见 bst.go 中的 prev 和 next

// 返回字符串 s 在 trie 中的前缀个数
// https://codeforces.com/gym/101628/problem/K
// https://www.acwing.com/problem/content/144/
func (t *trie) countPrefixOfString(s string) (cnt int) {
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
// 此时 o.endCnt 保存子树字符串个数
// https://codeforces.com/gym/101628/problem/K
// LC1804 https://leetcode.cn/problems/implement-trie-ii-prefix-tree/
func (t *trie) countStringHasPrefix(p string) int {
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
// - 注：这题还可以用后缀数组+前缀和二分来做到 O(nlogn)
// https://codeforces.com/problemset/problem/113/B
func (t *trie) countDistinctSubstring(s string) (cnt int) {
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

// 扩展：见 acam.go 和 pam.go
