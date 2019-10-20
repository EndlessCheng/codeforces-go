package copypasta

import . "fmt"

type rbKeyType int   // *custom* 图方便可以全局替换
type rbValueType int // *custom* 图方便可以全局替换
type rbColor bool

const red, black rbColor = true, false

type rbNode struct {
	lr    [2]*rbNode
	sz    int
	msz   int
	key   rbKeyType
	value rbValueType
	c     rbColor // 指向父节点的颜色（根节点为黑）
}

func (o *rbNode) isRed() bool {
	if o == nil {
		return false
	}
	return o.c == red
}

// d=0: left
// d=1: right
func (o *rbNode) rotate(d int) *rbNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	x.c = o.c
	o.c = red
	// x.msz = o.msz; o.pushUp()
	o.pushUp()
	x.pushUp()
	return x
}

func (o *rbNode) flipColors() {
	o.c = red
	o.lr[0].c = black
	o.lr[1].c = black
}

func (o *rbNode) pushUp() {
	sz := 1
	msz := int(o.value)
	if ol := o.lr[0]; ol != nil {
		sz += ol.sz
		msz += ol.msz
	}
	if or := o.lr[1]; or != nil {
		sz += or.sz
		msz += or.msz
	}
	o.sz = sz
	o.msz = msz
}

type rbTree struct {
	root       *rbNode
	comparator func(a, b rbKeyType) int
}

func newRBTree() *rbTree {
	// 设置如下返回值是为了方便使用 rbNode 中的 lr 数组
	return &rbTree{comparator: func(a, b rbKeyType) int {
		if a < b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	}}
}

func (t *rbTree) _put(o *rbNode, key rbKeyType, value rbValueType) *rbNode {
	if o == nil {
		return &rbNode{sz: 1, msz: 1, key: key, value: value, c: red}
	}
	if cmp := t.comparator(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], key, value)
		if o.lr[1].isRed() && !o.lr[0].isRed() {
			o = o.rotate(0)
		}
		if o.lr[0].isRed() && o.lr[0].lr[0].isRed() {
			o = o.rotate(1)
		}
		if o.lr[0].isRed() && o.lr[1].isRed() {
			o.flipColors()
		}
	} else {
		o.value = value
		//o.value += value
	}
	o.pushUp()
	return o
}

func (t *rbTree) put(key rbKeyType, value rbValueType) {
	t.root = t._put(t.root, key, value)
	t.root.c = black
}

//

func (c rbColor) String() string {
	if c == red {
		return "☀"
	}
	return "🌙"
}

func (o *rbNode) String() string {
	var s string
	if o.value == 1 {
		s = Sprintf("%v", o.key)
	} else {
		s = Sprintf("%v(%v)", o.key, o.value)
	}
	s += Sprintf("[sz:%d,msz:%d,c:%v]", o.sz, o.msz, o.c)
	return s
}

func (o *rbNode) draw(prefix string, isTail bool, str *string) {
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

func (t *rbTree) String() string {
	if t.root == nil {
		return "BST (empty)\n"
	}
	str := "BST\n"
	t.root.draw("", true, &str)
	return str
}
