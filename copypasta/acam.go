package copypasta

import (
	. "fmt"
	"strings"
)

/* AC 自动机   Aho–Corasick algorithm (ACAM) / Deterministic Finite Automaton (DFA)

如果我们既能知道前缀信息，又能知道后缀信息，就可以做字符串匹配。
建议先运行 acam_test.go 感受下 AC 自动机是怎么用的

https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
https://en.wikipedia.org/wiki/Deterministic_finite_automaton
https://oi-wiki.org/string/ac-automaton/
应用 https://cp-algorithms.com/string/aho_corasick.html
https://ac.nowcoder.com/study/live/738/4/1
https://zhuanlan.zhihu.com/p/533603249
https://www.cnblogs.com/sclbgw7/p/9260756.html

https://www.luogu.com.cn/problem/P3808
https://www.luogu.com.cn/problem/P3796
https://www.luogu.com.cn/problem/P5357
LC1032 https://leetcode.cn/problems/stream-of-characters/
LC1408 https://leetcode.cn/problems/string-matching-in-an-array/
todo 最长前缀查询 https://www.luogu.com.cn/problem/P5231
 NOI11 阿狸的打字机 https://www.luogu.com.cn/problem/P2414
 https://www.acwing.com/solution/content/25473/
 https://www.acwing.com/solution/content/54646/
 https://codeforces.com/problemset/problem/163/E
 https://codeforces.com/problemset/problem/1437/G
 https://codeforces.com/problemset/problem/963/D
 AC 自动机 fail 树 DFS 序上建可持久化线段树 https://codeforces.com/problemset/problem/547/E 2800
 - 后缀自动机 next 指针 DAG 图上跑 SG 函数 https://www.jisuanke.com/contest/1209/problems A 题 - A String Game https://www.jisuanke.com/problem/A1623
 - 注：楼教主新男人八题 https://www.zhihu.com/question/269890748
 - 注：https://codeforces.com/blog/entry/68292?#comment-526002
 LC30 串联所有单词的子串 https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/
 ? LC616 给字符串添加加粗标签 https://leetcode-cn.com/problems/add-bold-tag-in-string/
 LC2781 最长合法子字符串的长度 https://leetcode.cn/problems/length-of-the-longest-valid-substring/solution/aczi-dong-ji-onjie-fa-wu-shi-chang-du-10-47dy/
 https://codeforces.com/contest/1801/problem/G 3400
*/
const acamNodeSize = 26

type acamNode struct {
	son [acamNodeSize]*acamNode
	cnt int //（子树中）完整字符串的个数
	idx int

	// 当 o.son[i] 不能匹配文本串 text 中的某个字符时，o.fail.son[i] 即为下一个待匹配节点
	fail *acamNode
	last *acamNode // 后缀链接（suffix link），用来快速跳到一定是模式串末尾的位置
}

type acam struct {
	patterns []string // 额外保存，方便 debug

	root    *acamNode
	nodeCnt uint32

	inDeg map[*acamNode]int
}

func newACAM(patterns []string) *acam {
	t := &acam{
		patterns: patterns,
		root:     &acamNode{},
		nodeCnt:  1,
		inDeg:    map[*acamNode]int{},
	}
	for i, s := range patterns {
		t.put(s, i+1) // 注意这里 +1 了
	}
	t.buildFail()
	return t
}

func (acam) ord(c rune) rune { return c - 'a' }

// 插入字符串 s，附带值 idx
func (t *acam) put(s string, idx int) {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		if o.son[b] == nil {
			t.nodeCnt++
			newNode := &acamNode{}
			o.son[b] = newNode
			t.inDeg[newNode] = 0
		}
		o = o.son[b]
		//o.cnt++ // 写法一：统计 o 对应的字符串是多少个完整字符串的前缀
	}
	o.cnt++ // 写法二：统计 o 上有多少个完整字符串
	o.idx = idx
	//o.end = true
}

func (t *acam) buildFail() {
	t.root.fail = t.root
	t.root.last = t.root
	q := make([]*acamNode, 0, t.nodeCnt)
	for i, son := range t.root.son[:] {
		if son == nil {
			t.root.son[i] = t.root
		} else {
			son.fail = t.root // 第一层的失配指针，都指向 ∅
			son.last = t.root
			q = append(q, son)
		}
	}
	// BFS
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		f := o.fail
		//o.end = o.end || f.end // o 是否为某个模式串的末尾
		for i, son := range o.son[:] {
			if son == nil {
				o.son[i] = f.son[i] // 虚拟子节点 o.son[i]，和 o.fail.son[i] 是同一个 
				continue
			}
			son.fail = f.son[i] // 下一个匹配位置
			t.inDeg[son.fail]++ // fail 树上的从 son 到 son.fail 的边
			if son.fail.cnt > 0 {
				son.last = son.fail
			} else {
				son.last = son.fail.last
			}
			q = append(q, son)
		}
	}
}

// 有多少个下标不同的模式串在文本串 text 里出现过
// https://www.luogu.com.cn/problem/P3808
// https://www.luogu.com.cn/record/136447022
func (t *acam) sumCountAllPatterns(text string) (cnt int) {
	o := t.root
	for _, b := range text {
		o = o.son[t.ord(b)]
		// 遍历 fail 链（fail 树上的从 o 到 root 的路径）
		// 由于只找模式串，用 last 快速跳 fail
		for f := o; f != t.root && f.cnt != -1; f = f.last {
			cnt += f.cnt
			f.cnt = -1 // 访问标记
		}
	}
	return
}

// 返回一个 pos 列表，其中 pos[i] 表示 patterns[i] 的【首字母】在文本串 text 的所有位置（未找到时为空）
// patterns 为模式串列表（互不相同），下标从 1 开始
// 如果只求个数，更快的做法见后面
func (t *acam) acSearch(text string) [][]int {
	pos := make([][]int, len(t.patterns))
	o := t.root
	for i, b := range text {
		o = o.son[t.ord(b)]
		// 注：如果可以进入 for 循环，表示当前匹配到了一个（尽可能长的）模式串，其余更短的要在 fail 链上找
		// 遍历 fail 链（fail 树上的从 o 到 root 的路径）
		// 由于只找模式串，用 last 快速跳 fail
		f := o
		if o.idx == 0 {
			f = o.last
		}
		for ; f != t.root; f = f.last {
			pIdx := f.idx - 1
			// 如果改为记录 i，则表示 patterns[pIdx] 的【末尾字母】在 text 的位置
			pos[pIdx] = append(pos[pIdx], i-len(t.patterns[pIdx])+1)
		}
	}
	return pos
}

// 返回一个 cnt 列表，其中 cnt[i] 表示 patterns[i] 的在文本串 text 的出现次数（未找到时为 0）
// patterns 为模式串列表（互不相同），下标从 1 开始
// https://www.luogu.com.cn/problem/P5357
// https://www.luogu.com.cn/record/136429060
func (t *acam) acSearchCount(text string) []int {
	// 【注意】调用前把 put 中的 o.cnt++ 去掉！
	o := t.root
	for _, b := range text {
		o = o.son[t.ord(b)]
		// 本来应该像上面那样一路找到 t.root，但这样太慢了
		// 可以先打个标记，然后在 fail 树上跑拓扑序一起统计
		o.cnt++
	}

	cnt := make([]int, len(t.patterns))
	deg := t.inDeg
	q := make([]*acamNode, 0, len(deg)+1)
	for v, d := range deg {
		if d == 0 {
			q = append(q, v)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		// 如果 v 不是某个模式串的末尾字母，则 v.idx = 0
		if v.idx > 0 {
			cnt[v.idx-1] = v.cnt
		}
		w := v.fail // 注意不能用 last，会漏掉中间打上的 cnt 标记
		w.cnt += v.cnt
		if deg[w]--; deg[w] == 0 {
			q = append(q, w)
		}
	}
	return cnt
}

func (t *acam) debug(text string) {
	Println("text:", text)
	Println("patterns (下面简称 p)")
	for i, p := range t.patterns {
		Printf("%d: %s\n", i, p)
	}

	o := t.root
	for i, b := range text {
		o = o.son[t.ord(b)]
		_f := o
		if o.idx == 0 {
			_f = o.last
		}
		cnt := 0
		for ; _f != t.root; _f = _f.last {
			cnt++
		}
		if cnt == 0 {
			continue
		}

		Println()
		Println(text)
		Print(strings.Repeat(" ", i))
		Printf("^ i=%d\n", i)
		Println("找到", cnt, "个模式串")

		// 用 last
		f := o
		if o.idx == 0 {
			f = o.last
		}
		for ; f != t.root; f = f.last {
			pIdx := f.idx - 1
			Printf("[FAST] p[%d]=%s\n", pIdx, t.patterns[pIdx])
		}

		// 只用 fail 指针，不用 last
		for f := o; f != t.root; f = f.fail {
			if f.idx == 0 {
				Println("[SLOW] skip")
				continue
			}
			pIdx := f.idx - 1
			Printf("[SLOW] p[%d]=%s\n", pIdx, t.patterns[pIdx])
		}
	}
}
