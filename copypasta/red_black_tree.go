package copypasta

import . "fmt"

type rbKeyType int
type rbValueType int
type rbColor bool

const red, black rbColor = true, false

type rbNode struct {
	ch    [2]*rbNode
	sz    int
	msz   int
	key   rbKeyType
	value rbValueType
	c     rbColor // 指向父节点的颜色（根节点为黑）
}

func (o *rbNode) isRed() bool {
	return o != nil && o.c == red
}

func (o *rbNode) pushUp() {
	sz := 1
	msz := int(o.value)
	if ol := o.ch[0]; ol != nil {
		sz += ol.sz
		msz += ol.msz
	}
	if or := o.ch[1]; or != nil {
		sz += or.sz
		msz += or.msz
	}
	o.sz = sz
	o.msz = msz
}

// d=0：左旋，返回 o 的右儿子
// d=1：右旋，返回 o 的左儿子
func (o *rbNode) rotate(d int) *rbNode {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	x.c = x.ch[d].c
	x.ch[d].c = red
	// 或者写成
	// x.msz = o.msz; o.pushUp()
	o.pushUp()
	x.pushUp()
	return x
}

func (o *rbNode) flipColors() {
	o.c = !o.c
	o.ch[0].c = !o.ch[0].c
	o.ch[1].c = !o.ch[1].c
}

func (o *rbNode) moveRedLeft() *rbNode {
	o.flipColors()
	if o.ch[1].ch[0].isRed() {
		o.ch[1] = o.ch[1].rotate(1)
		o = o.rotate(0)
		o.flipColors()
	}
	return o
}

func (o *rbNode) moveRedRight() *rbNode {
	o.flipColors()
	if o.ch[0].ch[0].isRed() {
		o = o.rotate(1)
		o.flipColors()
	}
	return o
}

func (o *rbNode) balance() *rbNode {
	if o.ch[1].isRed() {
		o = o.rotate(0)
	}
	if o.ch[0].isRed() && o.ch[0].ch[0].isRed() {
		o = o.rotate(1)
	}
	if o.ch[0].isRed() && o.ch[1].isRed() {
		o.flipColors()
	}
	o.pushUp()
	return o
}

func (o *rbNode) min() *rbNode {
	for o.ch[0] != nil {
		o = o.ch[0]
	}
	return o
}

func (o *rbNode) deleteMin() *rbNode {
	if o.ch[0] == nil {
		return nil
	}
	if !o.ch[0].isRed() && !o.ch[0].ch[0].isRed() {
		o = o.moveRedLeft()
	}
	o.ch[0] = o.ch[0].deleteMin()
	return o.balance()
}

type rbTree struct {
	root *rbNode
}

func (t *rbTree) compare(a, b rbKeyType) int {
	switch {
	case a < b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func (t *rbTree) _put(o *rbNode, key rbKeyType, value rbValueType) *rbNode {
	if o == nil {
		return &rbNode{sz: 1, msz: 1, key: key, value: value, c: red}
	}
	if cmp := t.compare(key, o.key); cmp >= 0 {
		o.ch[cmp] = t._put(o.ch[cmp], key, value)
		if o.ch[1].isRed() && !o.ch[0].isRed() {
			o = o.rotate(0)
		}
		if o.ch[0].isRed() && o.ch[0].ch[0].isRed() {
			o = o.rotate(1)
		}
		if o.ch[0].isRed() && o.ch[1].isRed() {
			o.flipColors()
		}
	} else {
		o.value = value
	}
	o.pushUp()
	return o
}

func (t *rbTree) put(key rbKeyType, value rbValueType) {
	t.root = t._put(t.root, key, value)
	t.root.c = black
}

func (t *rbTree) get(key rbKeyType) *rbNode {
	for o := t.root; o != nil; {
		if cmp := t.compare(key, o.key); cmp >= 0 {
			o = o.ch[cmp]
		} else {
			return o
		}
	}
	return nil
}

func (t *rbTree) getStack(key rbKeyType) (stack []*rbNode) {
	for o := t.root; o != nil; {
		stack = append(stack, o)
		if cmp := t.compare(key, o.key); cmp >= 0 {
			o = o.ch[cmp]
		} else {
			return
		}
	}
	return nil
}

func (t *rbTree) _delete(o *rbNode, key rbKeyType) *rbNode {
	if cmp := t.compare(key, o.key); cmp == 0 {
		if !o.ch[0].isRed() && !o.ch[0].ch[0].isRed() {
			o = o.moveRedLeft()
		}
		o.ch[0] = t._delete(o.ch[0], key)
	} else {
		if o.ch[0].isRed() {
			o = o.rotate(1)
		}
		if t.compare(key, o.key) == -1 && o.ch[1] == nil {
			return nil
		}
		if !o.ch[1].isRed() && !o.ch[1].ch[0].isRed() {
			o = o.moveRedRight()
		}
		if t.compare(key, o.key) == -1 {
			x := o.ch[1].min()
			o.key = x.key
			o.value = x.value
			o.ch[1] = o.ch[1].deleteMin()
		} else {
			o.ch[1] = t._delete(o.ch[1], key)
		}
	}
	return o.balance()
}

func (t *rbTree) delete(key rbKeyType) {
	var o *rbNode
	if stack := t.getStack(key); stack != nil {
		stack, o = stack[:len(stack)-1], stack[len(stack)-1]
		if o.value > 1 {
			o.value--
			o.pushUp()
			for len(stack) > 0 {
				stack, o = stack[:len(stack)-1], stack[len(stack)-1]
				o.pushUp()
			}
			return
		}
	}
	if !t.root.ch[0].isRed() && !t.root.ch[1].isRed() {
		t.root.c = red
	}
	t.root = t._delete(t.root, key)
	if t.root != nil {
		t.root.c = black
	}
}

// 其余通用方法见 bst.go

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
	if o.ch[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		o.ch[1].draw(newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += o.String() + "\n"
	if o.ch[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		o.ch[0].draw(newPrefix, true, str)
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
