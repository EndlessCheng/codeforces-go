package copypasta

/* 动态树 link/cut tree (LCT)
https://en.wikipedia.org/wiki/Link/cut_tree
https://oi-wiki.org/ds/lct/

将一棵树剖分成由若干实边组成的实链，实链与实链之间用虚边相连
一棵 splay 对应一条实链，其中序遍历就是该实链从顶部往下的路径
这样一条实链上某个节点的父/子节点就对应着该节点在 splay 上的前驱/后继

https://www.cnblogs.com/flashhu/p/8324551.html
https://www.luogu.com.cn/blog/command-block/lct-xiao-ji
https://codeforces.com/blog/entry/80383

TIPS: 若要修改一个点，可以将这个点 splay 上来修改后 maintain，这样就不需要考虑修改这个点对父节点的影响了
TIPS: 对于卡常的题目，用 findRoot 判断连通性的逻辑可以用并查集代替

模板题 https://www.luogu.com.cn/problem/P3690
魔法森林（维护最大边权）https://www.luogu.com.cn/problem/P2387 AC 代码 https://www.luogu.com.cn/record/46975435 去掉 link 和 cut 的多余判断后 https://www.luogu.com.cn/record/46977629
最小差值生成树 https://www.luogu.com.cn/problem/P4234 https://codeforces.com/edu/course/2/lesson/7/2/practice/contest/289391/problem/F
todo https://ac.nowcoder.com/acm/contest/4643/F 题解 https://ac.nowcoder.com/discuss/387703
 动态修改 + 任意两点 LCA 的权值的期望 https://codeforces.com/problemset/problem/482/E
*/

type lctNode struct {
	lr   [2]*lctNode
	fa   *lctNode
	mi   *lctNode
	id   int
	v, s int
	flip bool
}

func (o *lctNode) sum() int {
	if o != nil {
		return o.s
	}
	return 0
}

func (o *lctNode) maintain() {
	o.s = o.v ^ o.lr[0].sum() ^ o.lr[1].sum()
}

// EXTRA: 最小差值生成树（见 graph.minDiffMST）
func (o *lctNode) maintainMin() {
	var n int // 节点个数

	o.mi = o
	for _, ch := range o.lr {
		if ch != nil && ch.mi.id >= n && (o.mi.id < n || ch.mi.id < o.mi.id) {
			o.mi = ch.mi
		}
	}
}

func (o *lctNode) doFlip() {
	if o != nil {
		o.flip = !o.flip
	}
}

func (o *lctNode) pushDown() {
	if o.flip {
		o.lr[0].doFlip()
		o.lr[1].doFlip()
		o.lr[0], o.lr[1] = o.lr[1], o.lr[0]
		o.flip = false
	}
}

// true 表示 o 是实链的顶点，即 o 为树的根节点或 o 和 o.fa 之间为虚边
// false 则表示 o 和 o.fa 之间为实边
func (o *lctNode) isRoot() bool {
	return o.fa == nil || o != o.fa.lr[0] && o != o.fa.lr[1]
}

func (o *lctNode) rotate() {
	y := o.fa
	z := y.fa
	if !y.isRoot() {
		if z.lr[0] == y {
			z.lr[0] = o
		} else {
			z.lr[1] = o
		}
	}
	o.fa = z
	d := 0
	if y.lr[1] == o {
		d = 1
	}
	y.lr[d] = o.lr[d^1]
	if o.lr[d^1] != nil {
		o.lr[d^1].fa = y
	}
	o.lr[d^1] = y
	y.fa = o
	y.maintain()
	o.maintain()
}

// 将 o 旋转到 splay 的根
// 由于直接从 o 出发旋转到根，所以要先把标记从根传递到 o 的子节点（注意不是 o），再旋转
func (o *lctNode) splay() {
	s := []*lctNode{o}
	for x := o; !x.isRoot(); x = x.fa {
		s = append(s, x.fa)
	}
	for i := len(s) - 1; i >= 0; i-- {
		s[i].pushDown()
	}
	for !o.isRoot() {
		// 直线：旋转 y 再旋转 o
		// 折线：旋转两次 o
		if y := o.fa; !y.isRoot() {
			if y.lr[0] == o == (y.fa.lr[0] == y) {
				y.rotate()
			} else {
				o.rotate()
			}
		}
		o.rotate()
	}
}

// 建立一条从树根到 o 的实链（极大实链，即 o 位于该实链末尾，下同），然后将 o 旋转到 splay 的根
func (o *lctNode) access() {
	x := o
	for y := (*lctNode)(nil); o != nil; o = o.fa {
		o.splay()
		o.lr[1] = y // 首次循环：o 与其之后建立虚边；剩余循环：路径上的虚边变成实边
		o.maintain()
		y = o
	}
	x.splay()
}

// 建立一条从树根到 o 的实链，然后将 o 变成树的根节点
func (o *lctNode) makeRoot() {
	o.access()
	o.flip = !o.flip // 将整条实链反向，o 就变成了树的根节点
}

// 建立一条从树根到 o 的实链，然后将树根旋转到 splay 的根，并返回树根
func (o *lctNode) findRoot() *lctNode {
	o.access()
	for o.lr[0] != nil {
		//o.pushDown() // 可以省略，后续 splay 时会传递标记
		o = o.lr[0]
	}
	o.splay()
	return o
}

// 将 o 变成树的根节点，然后建立（分离出，即 split）一条从 o 到 p 的实链，且将 p 旋转到 splay 的根
// 方便做区间操作
func (o *lctNode) split(p *lctNode) {
	o.makeRoot()
	p.access()
}

// 将 o 变成树的根节点，若 o 和 p 之间不连通，则在 o 和 p 之间连一条虚边
func (o *lctNode) link(p *lctNode) {
	o.makeRoot()
	if p.findRoot() != o { // 保证不连通时可以去掉，或者用并查集代替
		o.fa = p
	}
}

// 将 o 变成树的根节点，若 o 和 p 之间存在边，则删除该边
func (o *lctNode) cut(p *lctNode) {
	o.makeRoot()
	// 由于 o 为根节点，若 p 是 o 的后继，p 必无左子树
	if p.findRoot() == o && p.lr[0] == nil && p.fa == o {
		p.fa = nil
		o.lr[1] = nil
		//o.maintain() // 可以省略，后续 splay 时会维护成正确的值，下同
	}
}

// 题目保证 o 和 p 连通时的写法
func (o *lctNode) mustCut(p *lctNode) {
	o.split(p)
	o.fa = nil
	p.lr[0] = nil
	//p.maintain()
}
