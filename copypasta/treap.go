package copypasta

import . "fmt"

/* Treap=Tree+Heap
https://oi-wiki.org/ds/treap/
https://en.wikipedia.org/wiki/Treap
部分代码参考刘汝佳实现，见 https://github.com/klb3713/aoapc-book/blob/master/TrainingGuide/bookcodes/ch3/la5031.cpp
模板题 https://www.luogu.com.cn/problem/P3369
题目推荐 https://cp-algorithms.com/data_structures/treap.html#toc-tgt-8

耗时大约是红黑树（父节点实现）的 1.2 倍
*/

// 用 GoLand 的话强烈建议加入到 Live Templates 中，比赛时直接敲快捷键
type tpKeyType int
type tpValueType int

type tpNode struct {
	lr       [2]*tpNode
	priority uint // max heap
	sz       int
	msz      int
	key      tpKeyType
	val      tpValueType // dupCnt for multiset
}

// 设置如下返回值是为了方便使用 tpNode 中的 lr 数组
func (o *tpNode) cmp(b tpKeyType) int {
	switch {
	case b < o.key:
		return 0 // 左儿子
	case b > o.key:
		return 1 // 右儿子
	default:
		return -1
	}
}

func (o *tpNode) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *tpNode) mSize() int {
	if o != nil {
		return o.msz
	}
	return 0
}

func (o *tpNode) pushUp() {
	o.sz = 1 + o.lr[0].size() + o.lr[1].size()
	o.msz = int(o.val) + o.lr[0].mSize() + o.lr[1].mSize()
}

// 旋转，并维护子树大小
// d=0: left
// d=1: right
func (o *tpNode) rotate(d int) *tpNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	// x.sz = o.sz; x.msz = o.msz; o.pushUp()
	o.pushUp()
	x.pushUp()
	return x
}

type treap struct {
	rd   uint
	root *tpNode
}

func newTreap() *treap { return &treap{rd: 1} }

// https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
// https://en.wikipedia.org/wiki/Xorshift
func (t *treap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

// 先和二叉搜索树的插入一样，先把要插入的点插入到一个叶子上，并随机分配一个优先级，
// 然后跟维护堆一样，如果当前节点的优先级比根大就旋转，如果当前节点是根的左儿子就右旋如果当前节点是根的右儿子就左旋
func (t *treap) _put(o *tpNode, key tpKeyType, val tpValueType) *tpNode {
	if o == nil {
		return &tpNode{priority: t.fastRand(), sz: 1, msz: 1, key: key, val: val}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, val)
		// 如果当前节点的优先级比根大就旋转
		if o.lr[d].priority > o.priority {
			// 如果当前节点是根的左儿子就右旋，反之左旋
			o = o.rotate(d ^ 1)
		}
	} else {
		//o.val = val
		o.val += val
	}
	o.pushUp()
	return o
}

func (t *treap) put(key tpKeyType, val tpValueType) { t.root = t._put(t.root, key, val) }

// 因为 Treap 满足堆性质，所以只需要把要删除的节点旋转到叶节点上，然后直接删除就可以了
// 具体的方法就是每次找到优先级最大的儿子，向与其相反的方向旋转，这样要删除的节点会不断下降直到叶节点，然后直接删除
func (t *treap) _delete(o *tpNode, key tpKeyType) *tpNode {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		if o.val > 1 {
			o.val--
		} else {
			if o.lr[1] == nil {
				return o.lr[0]
			}
			if o.lr[0] == nil {
				return o.lr[1]
			}
			// o 有两颗子树，把优先级高的子树旋转到根，然后递归在另一颗子树中删除 o
			d = 0
			if o.lr[0].priority > o.lr[1].priority {
				d = 1
			}
			o = o.rotate(d)
			o.lr[d] = t._delete(o.lr[d], key)
		}
	}
	o.pushUp()
	return o
}

func (t *treap) delete(key tpKeyType) { t.root = t._delete(t.root, key) }

//

func (o *tpNode) String() string {
	var s string
	if o.val == 1 {
		s = Sprintf("%v", o.key)
	} else {
		s = Sprintf("%v(%v)", o.key, o.val)
	}
	s += Sprintf("[sz:%d,msz:%d]", o.sz, o.msz)
	return s
}

func (o *tpNode) draw(prefix string, isTail bool, str *string) {
	if o.lr[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		o.lr[1].draw(newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += o.String() + "\n"
	if o.lr[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		o.lr[0].draw(newPrefix, true, str)
	}
}

func (t *treap) String() string {
	if t.root == nil {
		return "BST (empty)\n"
	}
	str := "BST\n"
	t.root.draw("", true, &str)
	return str
}
