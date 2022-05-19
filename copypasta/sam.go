package copypasta

/* 后缀自动机 Suffix automaton (SAM)
【推荐】https://oi-wiki.org/string/sam/
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

模板题 https://www.luogu.com.cn/problem/P3804
后缀自动机一·基本概念 http://hihocoder.com/problemset/problem/1441
后缀自动机二·重复旋律5 http://hihocoder.com/problemset/problem/1445
后缀自动机三·重复旋律6 http://hihocoder.com/problemset/problem/1449
后缀自动机四·重复旋律7 http://hihocoder.com/problemset/problem/1457
后缀自动机五·重复旋律8 http://hihocoder.com/problemset/problem/1465
后缀自动机六·重复旋律9 http://hihocoder.com/problemset/problem/1466

多串最长公共子串（另见后缀数组）SPOJ LCS2 https://www.luogu.com.cn/problem/SP1812 https://loj.ac/p/171 LC1923/周赛248D https://leetcode-cn.com/problems/longest-common-subpath/
第 k 小子串（也可以用后缀数组做，见题解区）SPOJ SUBLEX https://www.luogu.com.cn/problem/SP7258 TJOI15 弦论 https://www.luogu.com.cn/problem/P3975
动态本质不同子串个数（也可以用后缀数组做，见题解区）https://www.luogu.com.cn/problem/P4070
区间本质不同子串个数（与 LCT 结合）https://www.luogu.com.cn/problem/P6292
todo https://codeforces.com/problemset/problem/235/C
*/

// 如果超时/超内存，改用预分配内存池 + func init() { debug.SetGCPercent(-1) } 的写法
// 如果仍然超内存且环境为 64 位，则需要把指针改成下标的写法 https://www.luogu.com.cn/record/76046834 https://www.luogu.com.cn/record/76047438
const mx int = 1e7

var samNodes [2 * mx]*node

type next [26]*node // map[int]*node

type node struct {
	fa  *node
	ch  next
	len int

	rev []*node // EXTRA: 所有 fa 的反向边构成一棵以 sam.nodes[0] 为根的树
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
		m.append(int(b - 'a')) // 'A'
	}
}

func (m *sam) buildRev() {
	for _, o := range m.nodes[1:] {
		o.fa.rev = append(o.fa.rev, o)
	}
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
