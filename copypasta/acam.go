package copypasta

import (
	. "fmt"
	"strings"
)

/* AC 自动机（多模式串匹配）   Aho–Corasick Automaton (ACAM) / Deterministic Finite Automaton (DFA)

Python/Java/C++ 的实现见（方法三）https://leetcode.cn/problems/construct-string-with-minimum-cost/solution/hou-zhui-shu-zu-by-endlesscheng-32h9/

可视化 https://wiomoc.github.io/aho-corasick-viz/ 蓝色是 fail 指针，绿色是 last 指针
      https://brunorb.com/aho-corasick/
      https://daniel.lawrence.lu/blog/y2014m03d25/

如果我们既知道前缀信息（trie），又知道后缀信息（fail），就可以做字符串匹配：
前缀的后缀就是子串，只要遍历到所有前缀，对每个前缀做「后缀匹配」，就完成了字符串匹配（统计子串出现次数）
推荐用 https://codeforces.com/problemset/problem/547/E 来理解
举例：对于 "ababab" 来说，把它的每个前缀节点都 +1，站在 fail 树的视角上看，相当于把（其中）一条全为 "a" 和一条全为 "b" 的树链上的【部分】节点 +1
统计 "ab" / "ba" 这些子串的出现次数，即「后缀匹配」的成功次数，只需要计算 "ab" / "ba" 在 fail 树上的子树点权和即可
这可以用树状数组维护

https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
https://en.wikipedia.org/wiki/Deterministic_finite_automaton
https://oi-wiki.org/string/ac-automaton/
应用 https://cp-algorithms.com/string/aho_corasick.html
https://ac.nowcoder.com/study/live/738/4/1
https://zhuanlan.zhihu.com/p/533603249
https://www.cnblogs.com/sclbgw7/p/9260756.html
todo 题单 https://www.luogu.com.cn/training/9372
         https://www.luogu.com.cn/training/53971
         https://www.cnblogs.com/alex-wei/p/Common_String_Theory_Theory_automaton_related.html
         https://ac.nowcoder.com/acm/contest/29086

Trie 图：合并 trie 树和 fail 树（求联集）
从 root 出发 DFS，跳过 end=true 的节点，如果找到环，那么就可以构造一个无限长的文本串，它不包含任何模式串
注意 buildFail 的时候要加上 o.end = o.end || f.end 这句话，
因为一个模式串可能包含其它模式串，比如有 "abcd" 和 "bc" 这两个模式串，要在 "abc" 的 "c" 上也标记 end=true
判环可以用 DFS + 三色标记法，见 graph.go 的「有向图的环」
[POI2000] 病毒 https://www.luogu.com.cn/problem/P2444

TIPS: 动态增删模式串，同时询问「查询所有模式串在文本串中的出现次数」，可以改为离线（先把所有模式串加入 AC 自动机）
当对一个节点增删 end 标记时，如果只对这一个节点修改，那么询问就需要遍历整个 fail 链，太慢了
换个思路：改为 end 标记会对它的 fail 子树全部 +1/-1，这样就可以在遍历文本串时单点询问了
这可以用【DFS 序 + 差分树状数组】实现，见 graph_tree.go 和 fenwick_tree.go
https://codeforces.com/problemset/problem/163/E 2800
- 弱化版 https://ac.nowcoder.com/acm/problem/14612
- 牛客这题也可以用分治 https://ac.nowcoder.com/acm/contest/view-submission?submissionId=53548785
在线做法：二进制分组 https://codeforces.com/problemset/problem/710/F 2400

https://www.luogu.com.cn/problem/P3808
https://www.luogu.com.cn/problem/P3796
每个模式串在文本串中的出现次数 https://www.luogu.com.cn/problem/P5357 双倍经验 https://www.luogu.com.cn/problem/P3966
LC1032 https://leetcode.cn/problems/stream-of-characters/
https://leetcode.cn/problems/construct-string-with-minimum-cost/
- https://ac.nowcoder.com/acm/contest/58568/F
todo LC1408 模式串之间的包含关系 https://leetcode.cn/problems/string-matching-in-an-array/
- https://leetcode.cn/problems/string-matching-in-an-array/submissions/484231678/
结合线段树优化 DP https://www.luogu.com.cn/problem/P7456
结合数位 DP https://ac.nowcoder.com/acm/problem/20366
- dfs 传参用 node *acamNode，记忆化的时候用 node.nodeID 代替，这样可以用数组而不是 map 记忆化，效率提高 10 倍
《AC 自动机 fail 树 DFS 序上建可持久化线段树》https://codeforces.com/problemset/problem/547/E 2800
- 其实不需要可持久化线段树，差分树状数组就行：s[i] 的每个前缀都可以在它的 fail 子树中的所有模式串内
- 包含 s[k] 的状态节点一定位于 s[k] 末尾节点的子树内
- 注：后缀自动机 next 指针 DAG 图上跑 SG 函数 https://www.jisuanke.com/contest/1209/problems A String Game https://www.jisuanke.com/problem/A1623
- 注：楼教主新男人八题 https://www.zhihu.com/question/269890748
- 注：https://codeforces.com/blog/entry/68292?#comment-526002
todo 最长前缀查询 https://www.luogu.com.cn/problem/P5231
 NOI11 阿狸的打字机 https://www.luogu.com.cn/problem/P2414
 https://www.luogu.com.cn/problem/P3121
 https://www.luogu.com.cn/problem/P3041
 https://www.luogu.com.cn/problem/P4052
 https://www.luogu.com.cn/problem/P3311
 https://www.luogu.com.cn/problem/P2292
 https://www.luogu.com.cn/problem/P5840
 二进制分组 https://codeforces.com/problemset/problem/710/F 2400
 https://codeforces.com/problemset/problem/1202/E 2400
 https://codeforces.com/problemset/problem/696/D 2500
 https://codeforces.com/problemset/problem/963/D 2500
 https://codeforces.com/problemset/problem/1437/G 2600
 https://codeforces.com/problemset/problem/1739/F 2600
 https://codeforces.com/problemset/problem/1207/G 2700
 https://codeforces.com/problemset/problem/163/E 2800
 https://codeforces.com/problemset/problem/1801/G 3400
 https://codeforces.com/gym/102511/problem/G ICPC Final 2019 G
 LC30 串联所有单词的子串 https://leetcode.cn/problems/substring-with-concatenation-of-all-words/
 ? LC616 给字符串添加加粗标签 https://leetcode.cn/problems/add-bold-tag-in-string/
 LC2781 最长合法子字符串的长度 https://leetcode.cn/problems/length-of-the-longest-valid-substring/solution/aczi-dong-ji-onjie-fa-wu-shi-chang-du-10-47dy/
 https://www.acwing.com/solution/content/25473/
 https://www.acwing.com/solution/content/54646/
*/

// 如果 MLE 请把指针替换成 uint32，代码见 https://codeforces.com/contest/163/submission/233981400

// 注意，AC 自动机中有两个「匹配」：
// fail: 失配时，移动到了字典树上的某个节点，但这个节点不一定是模式串的末尾
// last: 移动到模式串的末尾

const acamNodeSize = 26

type acamNode struct {
	son [acamNodeSize]*acamNode
	// 当 o.son[i] 不能匹配文本串 text 中的某个字符时，o.fail.son[i] 即为下一个待匹配节点
	// 特别地，如果 fail == root，则表示没有找到能匹配的模式串
	fail *acamNode
	last *acamNode // 后缀链接（suffix link），用来快速跳到一定是模式串末尾的节点

	idx int // 保存的信息也可以是 isEnd，字符串长度（节点深度）等
	cnt int //（子树中）完整字符串的个数

	nodeID int // 用于构建 fail 树
}

type gInfo struct{ l, r int } // [l,r]

type acam struct {
	patterns []string // 额外保存，方便 debug

	root    *acamNode
	nodeCnt int

	g     [][]int // fail 树
	gInfo []gInfo
	dfn   int

	inDeg map[*acamNode]int // 求拓扑序时有用
}

func newACAM(patterns []string) *acam {
	ac := &acam{
		patterns: patterns,
		root:     &acamNode{},
		nodeCnt:  1,
		inDeg:    map[*acamNode]int{},
	}
	for i, s := range patterns {
		ac.put(s, i+1) // 注意这里 +1 了
	}
	ac.buildFail()
	return ac
}

func (acam) ord(c rune) rune { return c - 'a' }

// 插入字符串 s，附带值 idx
func (ac *acam) put(s string, idx int) {
	o := ac.root
	for _, b := range s {
		b = ac.ord(b)
		if o.son[b] == nil {
			newNode := &acamNode{nodeID: ac.nodeCnt}
			o.son[b] = newNode
			//ac.inDeg[newNode] = 0
			ac.nodeCnt++
		}
		o = o.son[b]
		//o.cnt++ // 写法一：统计 o 对应的字符串是多少个完整字符串的前缀
	}
	o.cnt++ // 写法二：统计 o 上有多少个完整字符串
	o.idx = idx
	//o.end = true
}

// 层序遍历这棵字典树，得到一棵 fail 树
// 层序遍历可以保证：当节点出队时，其 fail 指针一定指向了最终的位置，无需写个 for 循环计算
func (ac *acam) buildFail() {
	ac.g = make([][]int, ac.nodeCnt) // fail 树
	ac.root.fail = ac.root
	ac.root.last = ac.root
	q := make([]*acamNode, 0, ac.nodeCnt)
	for i, son := range ac.root.son[:] {
		if son == nil {
			ac.root.son[i] = ac.root
		} else {
			son.fail = ac.root // 第一层的失配指针，都指向根节点 ∅
			son.last = ac.root
			ac.g[son.fail.nodeID] = append(ac.g[son.fail.nodeID], son.nodeID) // fail 树
			q = append(q, son)
		}
	}
	// BFS
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		for i, son := range o.son[:] {
			if son == nil {
				// 虚拟子节点 o.son[i]，和 o.fail.son[i] 是同一个 
				// 方便失配时直接跳到下一个可能匹配的位置（但不一定是某个模式串的最后一个字母）
				// 点评：这句话非常关键，相当于把 KMP 中的内层循环给消掉了，可以直接算出失配位置
				o.son[i] = o.fail.son[i]
				continue
			}
			// 计算失配位置
			son.fail = o.fail.son[i]
			ac.g[son.fail.nodeID] = append(ac.g[son.fail.nodeID], son.nodeID) // fail 树
			//ac.inDeg[son.fail]++
			if son.fail.cnt > 0 {
				son.last = son.fail
			} else {
				// 沿着 last 往上走，可以直接跳到一定是某个模式串末尾的节点（如果跳到 root 表示没有匹配）
				son.last = son.fail.last
			}
			q = append(q, son)
		}
	}
	// 跑完 BFS 后，我们得到了一棵 fail 树，只不过这棵树上的边都是从下往上指的
}

func (ac *acam) _buildDFN(v int) {
	ac.dfn++
	ac.gInfo[v].l = ac.dfn
	for _, w := range ac.g[v] {
		ac._buildDFN(w)
	}
	ac.gInfo[v].r = ac.dfn
}

func (ac *acam) buildDFN() {
	ac.gInfo = make([]gInfo, len(ac.g))
	ac._buildDFN(ac.root.nodeID)

	// 利用差分树状数组可以实现：添加删除模式串/查询有多少模式串在文本串中出现过
	// 见 https://codeforces.com/contest/163/submission/233925639
	//
	//bit := make(fenwick, ac.dfn+2)
	//
	//p := ac.gi[nodeIDs[i]]
	//bit.update(p.dfn, p.dfn+p.size, 1) // 左闭右开    1 是添加，-1 是删除
	//
	//cnt := 0
	//o := ac.root
	//for _, b := range text {
	//	o = o.son[ac.ord(b)]
	//	cnt += bit.pre(ac.gi[o.nodeID].dfn)
	//}
}

// 有多少个下标不同的模式串在文本串 text 里出现过
// https://www.luogu.com.cn/problem/P3808
// https://www.luogu.com.cn/record/136447022
func (ac *acam) sumCountAllPatterns(text string) (cnt int) {
	o := ac.root
	for _, b := range text {
		o = o.son[ac.ord(b)] // 如果没有匹配相当于移动到 fail 的 son[t.ord(b)]
		// 遍历 fail 链（fail 树上的从 o 到 root 的路径）
		for match := o; match != ac.root && match.cnt != -1; match = match.last {
			cnt += match.cnt
			match.cnt = -1 // 访问标记
		}
	}
	return
}

// 返回一个 pos 列表，其中 pos[i] 表示 patterns[i] 的【首字母】在文本串 text 的所有位置（未找到时为空）
// patterns 为模式串列表（互不相同），下标从 1 开始
// 如果只求个数，更快的做法见后面
func (ac *acam) acSearch(text string) [][]int {
	pos := make([][]int, len(ac.patterns))
	o := ac.root
	for i, b := range text {
		o = o.son[ac.ord(b)] // 如果没有匹配相当于移动到 fail 的 son[t.ord(b)]
		// 如果可以进入 for 循环，表示当前匹配到了一个尽可能长的模式串，其余更短的模式串要在 fail 链（last 链）上找
		for match := o; match != ac.root; match = match.last {
			if match.idx == 0 { // 注：只有 o 可能会触发 if，其余 last 不会触发 if
				continue
			}
			pIdx := match.idx - 1
			// 如果改为记录 i，则表示 patterns[pIdx] 的【末尾字母】在 text 的位置
			pos[pIdx] = append(pos[pIdx], i-len(ac.patterns[pIdx])+1)
		}
	}
	return pos
}

// 返回一个 cnt 列表，其中 cnt[i] 表示 patterns[i] 的在文本串 text 的出现次数（未找到时为 0）
// patterns 为模式串列表（互不相同），下标从 1 开始
// https://www.luogu.com.cn/problem/P5357 https://www.luogu.com.cn/problem/P3966
// https://www.luogu.com.cn/record/136429060
func (ac *acam) acSearchCount(text string) []int {
	// 【注意】调用前把 put 中的 o.cnt++ 去掉！
	o := ac.root
	for _, b := range text {
		o = o.son[ac.ord(b)]
		// 本来应该像上面那样一路找到 t.root，但这样太慢了
		// 可以先打个标记，然后在 fail 树上跑拓扑序一起统计
		o.cnt++
	}

	cnt := make([]int, len(ac.patterns))
	deg := ac.inDeg
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

func (ac *acam) debug(text string) {
	Println("text:", text)
	Println("patterns (下面简称 p)")
	for i, p := range ac.patterns {
		Printf("%d: %s\n", i, p)
	}

	o := ac.root
	for i, b := range text {
		o = o.son[ac.ord(b)]
		cnt := 0
		for f := o; f != ac.root; f = f.fail {
			if f.idx > 0 {
				cnt++
			}
		}
		if cnt == 0 {
			continue
		}

		Println()
		Println(text)
		Print(strings.Repeat(" ", i))
		Printf("^ i=%d\n", i)
		Println("找到", cnt, "个模式串")

		for f := o; f != ac.root; f = f.fail {
			if f.idx == 0 {
				//Println("skip")
				continue
			}
			pIdx := f.idx - 1
			Printf("p[%d]=%s\n", pIdx, ac.patterns[pIdx])
		}
	}
}
