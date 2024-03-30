package copypasta

/* 回文自动机（回文树）  Palindrome Automaton (PAM) / EerTree
原论文 https://arxiv.org/pdf/1506.04862.pdf
https://en.wikipedia.org/wiki/Palindrome_tree
https://oi-wiki.org/string/pam/

如果我们既知道前缀信息（trie），又知道后缀信息（fail），就可以做字符串匹配
对于回文计数类问题（比如求一个字符串有多少个本质不同的回文子串，每个又分别出现多少次），Manacher 就心有余而力不足了
刚好，PAM 的每个节点就表示一个本质不同回文串

定理：一个长为 n 的字符串，其本质不同回文串不超过 n 个
证明关键：一个回文串 t 的回文后缀，同时也是 t 的回文前缀，所以每次新增一个字符，至多增加一个本质不同回文串，更短的回文子串必然在前面出现过
详细证明：https://oi-wiki.org/string/pam/#%E5%AE%9A%E7%90%86

todo http://zhylj.cc/index.php/archives/26/
 https://baobaobear.github.io/post/20200416-pam/
 《字符串算法选讲》-金策
 https://zhuanlan.zhihu.com/p/92874690
 https://ac.nowcoder.com/courses/cover/live/738
 https://arxiv.org/pdf/1506.04862v2.pdf
 题单 https://www.luogu.com.cn/training/9372#information
https://hihocoder.com/problemset/problem/1602

模板题 https://www.luogu.com.cn/problem/P5496
本质不同回文子串的出现次数 https://www.luogu.com.cn/problem/P3649
todo 最小回文划分 https://codeforces.com/problemset/problem/932/G 2900
 能否划分成三段回文 LC1745 https://leetcode.cn/problems/palindrome-partitioning-iv/
 最长双回文串（另一种做法是 Manacher）https://www.luogu.com.cn/problem/P4555
 https://codeforces.com/gym/104787/problem/C
 [SHOI2011] 双倍回文 https://www.luogu.com.cn/problem/P4287
 https://www.luogu.com.cn/problem/P4762
*/
const pamNodeSize = 26

type pamNode struct {
	son [pamNodeSize]*pamNode
	// S(o.fail) 是 S(o) 的最大 border，也是 S(o) 的最长回文前缀/后缀（注意回文串的 border 也是一个回文串）
	fail *pamNode
	len  int

	// 节点编号从大到小，就是 fail 树的拓扑序
	nodeID int
}

type pam struct {
	s []byte

	evenRoot *pamNode // len = 0
	oddRoot  *pamNode // len = -1
	last     *pamNode

	nodeCnt int
}

func newPam(s string) *pam {
	evenRoot := &pamNode{len: 0, nodeID: 0}
	oddRoot := &pamNode{len: -1, nodeID: 1} // len=-1 使得奇根不可能失配
	oddRoot.fail = evenRoot
	evenRoot.fail = oddRoot

	t := &pam{
		s:        make([]byte, 0, len(s)),
		evenRoot: evenRoot,
		oddRoot:  oddRoot,
		last:     evenRoot, // 也可以是 oddRoot
		nodeCnt:  2,
	}
	for _, b := range s {
		t.insert(byte(b))
	}
	return t
}

func (pam) ord(c byte) byte { return c - 'a' }

// PAM 的构造，实际上就是求每个前缀的最长回文后缀
// 方法：枚举前一个位置的回文后缀，即 fail 链
func (t *pam) getFail(o *pamNode) *pamNode {
	i := len(t.s) - 1
	// 看 S(o) 的最长回文后缀的前一个字母是不是 t.s[i]
	for i-o.len-1 < 0 || t.s[i-o.len-1] != t.s[i] {
		o = o.fail
	}
	return o
}

func (t *pam) insert(b byte) {
	t.s = append(t.s, b)
	o := t.getFail(t.last)
	j := t.ord(b)
	if o.son[j] == nil {
		newNode := &pamNode{len: o.len + 2, fail: t.getFail(o.fail).son[j], nodeID: t.nodeCnt}
		t.nodeCnt++
		if newNode.fail == nil { // 没有真回文后缀
			newNode.fail = t.evenRoot
		}
		//newNode.num = newNode.fail.num + 1 // 子串右端点为 i=len(t.s)-1 的回文子串个数
		o.son[j] = newNode
	}
	t.last = o.son[j]
	// t.last.cnt++ // 统计当前节点对应多少个位置的最长回文后缀
	// 每个节点对应的回文串在整个字符串中的出现次数 = fail 树的子树和
}
