package copypasta

/* 替罪羊树
https://en.wikipedia.org/wiki/Scapegoat_tree

https://people.ksp.sk/~kuko/gnarley-trees/Scapegoat.html
lazy insert: let the tree grow and from time to time, when a subtree gets too imbalanced,
             rebuild the whole subtree from scratch into a perfectly balanced tree
lazy delete: just mark the node for deletion; when n/2 nodes are marked,
             rebuild the whole tree and throw away all deleted nodes

推荐：常见平衡树的α值+复杂度分析 https://riteme.site/blog/2016-4-6/scapegoat.html
https://oi-wiki.org/ds/sgt/
【可视化】https://www.bilibili.com/video/BV1sP4y1i7WB/
https://zhuanlan.zhihu.com/p/21263304
https://zhuanlan.zhihu.com/p/180545164
https://taodaling.github.io/blog/2019/04/19/%E6%9B%BF%E7%BD%AA%E7%BE%8A%E6%A0%91/

模板题 https://www.luogu.com.cn/problem/P3369 https://www.luogu.com.cn/problem/P6136
https://codeforces.com/contest/455/problem/D
*/

type sgtNode struct {
	lr  [2]*sgtNode
	key int
	val int // 如果是可重集，val 也可以当作删除标记
	sz  int
	del bool // 删除标记
}

func (o *sgtNode) cmp(b int) int {
	switch {
	case b < o.key:
		return 0 // 左儿子
	case b > o.key:
		return 1 // 右儿子
	default:
		return -1
	}
}

func (o *sgtNode) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *sgtNode) maintain() {
	o.sz = o.lr[0].size() + o.lr[1].size()
	if !o.del { // o.val > 0
		o.sz++
	}
}

func (o *sgtNode) nodes() []*sgtNode {
	nodes := make([]*sgtNode, 0, o.size())
	var f func(*sgtNode)
	f = func(o *sgtNode) {
		if o == nil {
			return
		}
		f(o.lr[0])
		if !o.del { // o.val > 0
			nodes = append(nodes, o)
		}
		f(o.lr[1])
	}
	f(o)
	return nodes
}

func buildSGT(nodes []*sgtNode) *sgtNode {
	if len(nodes) == 0 {
		return nil
	}
	m := len(nodes) / 2
	o := nodes[m]
	o.lr[0] = buildSGT(nodes[:m])
	o.lr[1] = buildSGT(nodes[m+1:])
	o.maintain()
	return o
}

func (o *sgtNode) rebuild() *sgtNode { return buildSGT(o.nodes()) }

type scapegoatTree struct {
	root   *sgtNode
	delCnt int
}

func (t *scapegoatTree) _put(o *sgtNode, key, val int) *sgtNode {
	if o == nil {
		return &sgtNode{key: key, val: val, sz: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, val)
	} else {
		if o.del { // o.val == 0
			o.del = false
			t.delCnt--
		}
		//o.val = val
		o.val += val
	}
	o.maintain()
	if sz := o.size() * 3; o.lr[0].size()*4 > sz || o.lr[1].size()*4 > sz { // alpha=3/4
		return o.rebuild()
	}
	return o
}

func (t *scapegoatTree) put(key, val int) { t.root = t._put(t.root, key, val) }

func (t *scapegoatTree) _delete(o *sgtNode, key int) {
	if o == nil {
		return
	}
	if d := o.cmp(key); d >= 0 {
		t._delete(o.lr[d], key)
	} else if !o.del { // o.val > 0
		//o.val--
		o.del = true
		t.delCnt++
	}
	o.maintain()
}

func (t *scapegoatTree) delete(key int) {
	t._delete(t.root, key)
	if t.delCnt > t.root.size()/2 {
		t.root = t.root.rebuild()
		t.delCnt = 0
	}
}

// 其余和 BST 有关的方法见 bst.go
// 注意求前驱后继的写法
