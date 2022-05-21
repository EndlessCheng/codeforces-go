package copypasta

import (
	. "fmt"
	"sort"
	"strings"
)

/* 后缀自动机 Suffix automaton (SAM)
将字符串的所有子串压缩后的结果
【推荐】https://oi-wiki.org/string/sam/ 证明了 SAM 的状态数不会超过 2n-1 (n>=2)，最坏情况下为 abbb...bbb
【推荐】可视化工具：SAM Drawer https://yutong.site/sam/
https://en.wikipedia.org/wiki/Suffix_automaton
https://www.bilibili.com/video/av756051240/
https://baobaobear.github.io/post/20200220-sam/
https://ouuan.github.io/post/%E5%90%8E%E7%BC%80%E8%87%AA%E5%8A%A8%E6%9C%BAsam%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/
https://codeforces.com/blog/entry/20861
《后缀自动机》，陈立杰
《后缀自动机在字典树上的拓展》，刘研绎
《后缀自动机及其应用》，张天扬
炫酷后缀树魔术 https://eternalalexander.github.io/2019/10/31/%E5%90%8E%E7%BC%80%E6%A0%91%E7%AE%80%E4%BB%8B/
后缀平衡树 https://www.luogu.com.cn/blog/CHiCO/CHiCO-Suffix-Balance-Tree

后缀自动机一·基本概念  http://hihocoder.com/problemset/problem/1441
后缀自动机二·重复旋律5 http://hihocoder.com/problemset/problem/1445
后缀自动机三·重复旋律6 http://hihocoder.com/problemset/problem/1449
后缀自动机四·重复旋律7 http://hihocoder.com/problemset/problem/1457
后缀自动机五·重复旋律8 http://hihocoder.com/problemset/problem/1465
后缀自动机六·重复旋律9 http://hihocoder.com/problemset/problem/1466

模板题：子串出现次数 https://www.luogu.com.cn/problem/P3804
多串最长公共子串（另见后缀数组）SPOJ LCS2 https://www.luogu.com.cn/problem/SP1812 https://loj.ac/p/171 LC1923/周赛248D https://leetcode-cn.com/problems/longest-common-subpath/
第 k 小子串（也可以用后缀数组做，见题解区）SPOJ SUBLEX https://www.luogu.com.cn/problem/SP7258 TJOI15 弦论 https://www.luogu.com.cn/problem/P3975
动态本质不同子串个数（也可以用后缀数组做，见题解区）https://www.luogu.com.cn/problem/P4070
区间本质不同子串个数（与 LCT 结合）https://www.luogu.com.cn/problem/P6292
todo https://codeforces.com/problemset/problem/235/C
*/

// 如果超时/超内存，改用预分配内存池 + func init() { debug.SetGCPercent(-1) } 的写法
// 如果仍然超内存且环境为 64 位，则需要把指针改成 int32 下标的写法 https://www.luogu.com.cn/record/76046834 https://www.luogu.com.cn/record/76047438
const mx int = 1e7

var samNodes [2 * mx]*node

type next [26]*node // map[int]*node

type node struct { // 也叫 state
	//      len 为该节点（endpos 等价类）中最长的子串长度
	// fa.len+1 为该节点（endpos 等价类）中最短的子串长度
	// 等价类大小为 len-fa.len
	// 等价类中的每个子串都是其最长子串的一系列连续后缀，即长度组成了区间 [fa.len+1,len]
	// 这一系列连续后缀之后更短的后缀，要去 fa, fa.fa, fa.fa.fa, ... 上找
	fa  *node
	ch  next
	len int

	// EXTRA: 所有 node->node.fa 的反向边 rev 构成一棵以 sam.nodes[0] 为根的树（sam.nodes[0] 表示空串）
	// 性质：（结合 AABABA 的 rev 理解 https://public.noi.top/image/1583315246012.png）
	// 1. 每个节点中的子串，是【紧跟】其儿子节点的子串的后缀
	// 2. 每个节点中的子串，都是以该节点为根的子树中的所有子串的后缀
	// 3. 两个节点的 LCA 对应着这两个节点的子串的最长公共后缀
	// 4. 叶子节点一定包含原串的前缀（根据性质 1 反证可得），但原串的前缀不一定在叶子节点中（比如前缀 A 是前缀 AA 的后缀）
	// 5. 由于【子串】等价于【前缀的后缀】，因此求子串在原串中的出现次数，
	//    可以先通过在 SAM 上找到该子串所处的节点，然后求以该节点为根的子树中，多有少个包含原串前缀的节点（性质 2）
	//    这可以通过在 rev 上统计子树信息来预处理得到
	rev []*node

	cnt           int
	debugSuffixes []string
}

type sam struct {
	nodes []*node
	last  *node
}

func newSam() *sam {
	m := &sam{}
	m.last = m.newNode(nil, next{}, 0)
	return m
}

func (m *sam) newNode(fa *node, _ch next, length int) *node {
	// 如果 next 是 map 则需要 clone
	//ch := make(next, len(_ch))
	//for c, o := range _ch {
	//	ch[c] = o
	//}
	//_ch = ch
	o := &node{fa: fa, ch: _ch, len: length}
	m.nodes = append(m.nodes, o)
	return o
}

func (m *sam) append(c int) {
	last := m.newNode(m.nodes[0], next{}, m.last.len+1)
	last.cnt = 1 // 或者标记这是个前缀节点
	for o := m.last; o != nil; o = o.fa {
		p := o.ch[c]
		if p == nil {
			o.ch[c] = last
			continue
		}
		if o.len+1 == p.len {
			last.fa = p
		} else {
			np := m.newNode(p.fa, p.ch, o.len+1)
			p.fa = np
			for ; o != nil && o.ch[c] == p; o = o.fa {
				o.ch[c] = np
			}
			last.fa = np
		}
		break
	}
	m.last = last
}

func (m *sam) buildSam(s string) {
	for _, b := range s {
		m.append(int(b - 'a')) // todo 'A'
	}
}

func (m *sam) buildRev() {
	for _, o := range m.nodes[1:] {
		o.fa.rev = append(o.fa.rev, o)
	}
}

func (m *sam) dfs(v *node) {
	for _, w := range v.rev {
		m.dfs(w)
		// ...

	}
}

// 多串 LCS 则去掉代码中的注释
func (m *sam) lcs(s string) (ans int) {
	//for _, o := range m.nodes {
	//	o.maxCommon = 0
	//}
	root := m.nodes[0]
	o, common := root, 0
	for _, b := range s {
		// 下面的结构形式十分类似 KMP
		b -= 'a'
		for o != root && o.ch[b] == nil {
			o = o.fa
			common = o.len
		}
		if o.ch[b] != nil {
			o = o.ch[b]
			common++
			if common > ans {
				ans = common
			}
		}
		//o.maxCommon = max(o.maxCommon, common)
	}
	//m.dfs(root)
	//for _, o := range m.nodes {
	//	o.ans = min(o.ans, o.maxCommon) // newNode 时 ans 初始化成 len
	//}
	return
}

// 返回 s 在 sam 中的最长前缀
// 特别地，若 s 在 sam 中，则返回 s 的长度
// JSOI12 https://www.luogu.com.cn/problem/P5231
func (m *sam) longestPrefix(s string) int {
	o := m.nodes[0]
	for i, b := range s {
		b -= 'a'
		if o.ch[b] == nil {
			return i
		}
		o = o.ch[b]
	}
	return len(s)
}

// debug 用
func (m *sam) unzipSAM() {
	root := m.nodes[0]
	// 如果 append 了新的元素导致 rev 树的结构变化，可以直接重建
	for _, o := range m.nodes {
		o.rev = nil
		o.debugSuffixes = nil
	}
	m.buildRev()

	var makeSuf func(*node, string)
	makeSuf = func(v *node, s string) {
		v.debugSuffixes = append(v.debugSuffixes, s)
		for i, w := range v.ch {
			if w != nil {
				makeSuf(w, s+string(byte('a'+i))) // 'A'
			}
		}
	}
	makeSuf(root, "")

	var dfs func(*node, int)
	dfs = func(v *node, depth int) {
		isPrefixNode := v.cnt == 1
		for _, w := range v.rev {
			dfs(w, depth+1)
			v.cnt += w.cnt
		}

		sort.Slice(v.debugSuffixes, func(i, j int) bool { return len(v.debugSuffixes[i]) > len(v.debugSuffixes[j]) })

		// 以 rev 树的形式打印所有【本质不同子串】（打印结果需要从下往上看）
		// 方括号表示原串的前缀，圆括号表示非前缀
		// 如果 SAM 构建正确，那么 debugSuffixes 必然是 debugSuffixes[0] 的一系列连续后缀
		if depth > 0 {
			var suf string
			if isPrefixNode { // 前缀节点
				suf = Sprintf("[%s]", v.debugSuffixes[0]) // 由于上面从长到短排序了，这里最长的一定是前缀
				if len(v.debugSuffixes) > 1 {
					suf += Sprintf(" (%s)", strings.Join(v.debugSuffixes[1:], " "))
				}
			} else { // 非前缀节点
				suf = Sprintf("(%s)", strings.Join(v.debugSuffixes, " "))
			}
			Print(strings.Repeat("    ", depth-1))
			Println(v.cnt, suf)
		}
	}
	dfs(root, 0)
}

// 广义 SAM
// https://oi-wiki.org/string/general-sam/
//
// todo https://www.luogu.com.cn/problem/P6139
//  https://codeforces.com/problemset/problem/1437/G
