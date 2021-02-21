package copypasta

/* 左偏树
https://en.wikipedia.org/wiki/Leftist_tree
代码参考 https://oi-wiki.org/ds/leftist-tree/

模板题 https://www.luogu.com.cn/problem/P3377
      https://www.luogu.com.cn/problem/P2713
      https://www.luogu.com.cn/problem/P1456
*/

// 默认写法为小根堆

type ltNode struct {
	lc, rc, fa *ltNode
	s, val, id int
}

func newLeftistTreeNode(id, val int) *ltNode {
	o := &ltNode{s: 1, val: val, id: id}
	o.fa = o
	return o
}

func (o *ltNode) sVal() int {
	if o != nil {
		return o.s
	}
	return 0
}

func (o *ltNode) findRoot() *ltNode {
	if o.fa != o {
		o.fa = o.fa.findRoot() // 路径压缩
	}
	return o.fa
}

func (o *ltNode) _merge(p *ltNode) *ltNode {
	if p == nil {
		return o
	}
	if o == nil {
		return p
	}
	if o.val > p.val || o.val == p.val && o.id > p.id { // 大根堆改成 <
		o, p = p, o
	}
	o.rc = o.rc._merge(p)
	if o.lc.sVal() < o.rc.sVal() {
		o.lc, o.rc = o.rc, o.lc
	}
	o.s = o.rc.sVal() + 1
	return o
}

// 注：push 一个节点就相当于 merge 这个节点
func (o *ltNode) merge(p *ltNode) {
	o = o.findRoot()
	p = p.findRoot()
	if o == p {
		return
	}
	q := o._merge(p)
	o.fa = q
	p.fa = q
}

func (o *ltNode) pop() *ltNode {
	o = o.findRoot()
	p := o.lc._merge(o.rc)
	o.fa = p
	if o.lc != nil {
		o.lc.fa = p
	}
	if o.rc != nil {
		o.rc.fa = p
	}
	return o
}
