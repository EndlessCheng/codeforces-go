package copypasta

/*
伸展树 splay
https://oi-wiki.org/ds/splay/
https://www.cnblogs.com/cjyyb/p/7499020.html
普通平衡树 https://www.luogu.com.cn/problem/P3369 https://www.luogu.com.cn/problem/P6136
文艺平衡树 https://www.luogu.com.cn/problem/P3391
*/

// 参考刘汝佳的实现，即不使用父节点的方案

type spKeyType int
type spValueType int

type spNode struct {
	lr  [2]*spNode
	sz  int
	key spKeyType
	val spValueType
}

// 设置如下返回值是为了方便使用 spNode 中的 lr 数组
func (o *spNode) cmpSz(k int) int8 {
	switch d := k - o.lr[0].size() - 1; {
	case d < 0:
		return 0 // 左儿子
	case d > 0:
		return 1 // 右儿子
	default:
		return -1
	}
}

func (o *spNode) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

// 对于取名叫 maintain 还是 pushUp，由于操作的对象是当前节点，个人认为取名 maintain 更为准确
func (o *spNode) maintain() {
	o.sz = 1 + o.lr[0].size() + o.lr[1].size()
}

func (o *spNode) pushDown() {
	// custom ...

}

// 旋转，并维护子树大小
// d=0：左旋，返回 o 的右儿子
// d=1：右旋，返回 o 的左儿子
func (o *spNode) rotate(d int8) *spNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	// x.sz = o.sz; x.msz = o.msz; o.maintain()
	o.maintain()
	x.maintain()
	return x
}

// 将子树 o 的第 k 小节点伸展到 o，返回该节点
// k 必须为正
// 注：其中的 if d==1 {...} 可以写成 k -= d * (o.lr[0].size() + 1)
func (o *spNode) splay(k int) (kth *spNode) {
	d := o.cmpSz(k)
	if d == -1 {
		return o
	}
	o.pushDown()
	if d > 0 {
		k -= o.lr[0].size() + 1
	}
	c := o.lr[d]
	c.pushDown()
	d2 := c.cmpSz(k)
	if d2 >= 0 {
		if d2 > 0 {
			k -= c.lr[0].size() + 1
		}
		c.lr[d2] = c.lr[d2].splay(k)
		if d2 == d {
			o = o.rotate(d ^ 1)
		} else {
			o.lr[d] = c.rotate(d)
		}
	}
	return o.rotate(d ^ 1)
}

// 分裂子树 o，把 o 的前 k 小个节点放在 left 子树，其他的放在 right 子树（left 节点为 o 的第 k 小节点）
// 0 < k <= o.size()，取等号时 right 为 nil
func (o *spNode) split(k int) (left, right *spNode) {
	left = o.splay(k)
	right = left.lr[1]
	left.lr[1] = nil
	left.maintain()
	return
}

// 把子树 right 合并进子树 o，返回合并前 o 的最大节点
// 子树 o 的所有元素比子树 right 中的小
// o != nil
func (o *spNode) merge(right *spNode) *spNode {
	o = o.splay(o.size())
	o.lr[1] = right
	o.maintain()
	return o
}

type splay struct{ root *spNode }
