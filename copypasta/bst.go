package copypasta

import . "fmt"

type tKeyType int
type tValueType int

type tnode struct {
	lr    [2]*tnode
	sz    int
	msz   int
	key   tKeyType
	value tValueType
}

func (o *tnode) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *tnode) mSize() int {
	if o != nil {
		return o.msz
	}
	return 0
}

func (o *tnode) pushUp() {
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

type bst struct {
	root *tnode
}

// 设置如下返回值是为了方便使用 tNode 中的 lr 数组
func (t *bst) compare(a, b tKeyType) int {
	switch {
	case a < b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func (t *bst) size() int   { return t.root.size() }
func (t *bst) empty() bool { return t.root == nil }

func (t *bst) get(key tKeyType) *tnode {
	for o := t.root; o != nil; {
		if cmp := t.compare(key, o.key); cmp >= 0 {
			o = o.lr[cmp]
		} else {
			return o
		}
	}
	return nil
}

// max <= key
// return nil if not found
// same like --upper_bound in C++ STL
func (t *bst) floor(key tKeyType) (floor *tnode) {
	for o := t.root; o != nil; {
		switch cmp := t.compare(key, o.key); {
		case cmp == 0:
			o = o.lr[0]
		case cmp > 0:
			floor = o
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

// 前驱（小于 key，且最大的数）
func (t *bst) prev(key tKeyType) (prev *tnode) {
	// 另一种写法
	// rank, _ := t.mRank(key)
	// return t.mSelect(rank - 1)
	for o := t.root; o != nil; {
		if cmp := t.compare(key, o.key); cmp <= 0 {
			o = o.lr[0]
		} else {
			prev = o
			o = o.lr[1]
		}
	}
	return
}

// min >= key
// return nil if not found
// same like lower_bound in C++ STL
func (t *bst) ceiling(key tKeyType) (ceiling *tnode) {
	for o := t.root; o != nil; {
		switch cmp := t.compare(key, o.key); {
		case cmp == 0:
			ceiling = o
			o = o.lr[0]
		case cmp > 0:
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

// 后继（大于 key，且最小的数)
func (t *bst) next(key tKeyType) (next *tnode) {
	// 另一种写法
	// rank, o := t.mRank(key)
	// if o != nil {
	// 	 rank += int(o.value)
	// }
	// return t.mSelect(rank)
	for o := t.root; o != nil; {
		if cmp := t.compare(key, o.key); cmp != 0 {
			o = o.lr[1]
		} else {
			next = o
			o = o.lr[0]
		}
	}
	return
}

func (t *bst) inRange(l, r int) bool {
	o := t.ceiling(tKeyType(l))
	return o != nil && int(o.key) <= r
}

// 小于 key 的键的数量
func (t *bst) mRank(key tKeyType) (cnt int, o *tnode) {
	for o = t.root; o != nil; {
		switch cmp := t.compare(key, o.key); {
		case cmp == 0:
			o = o.lr[0]
		case cmp > 0:
			cnt += int(o.value) + o.lr[0].mSize()
			o = o.lr[1]
		default:
			cnt += o.lr[0].mSize()
			return
		}
	}
	return
}

// 排名为 k 的节点 o（即有 k 个键小于 o.key）
func (t *bst) mSelect(k int) (o *tnode) {
	//if k < 0 {
	//	return
	//}
	for o = t.root; o != nil; {
		switch ls := o.lr[0].mSize(); {
		case k < ls:
			o = o.lr[0]
		case k > ls:
			k -= int(o.value) + ls
			if k < 0 {
				return
			}
			o = o.lr[1]
		default:
			return
		}
	}
	return
}

func (t *bst) min() (min *tnode) {
	for o := t.root; o != nil; o = o.lr[0] {
		min = o
	}
	return
}

func (t *bst) max() (max *tnode) {
	for o := t.root; o != nil; o = o.lr[1] {
		max = o
	}
	return
}

func (t *bst) keys() (keys []tKeyType) {
	var o *tnode
	q := []*tnode{t.root}
	for len(q) > 0 {
		o, q = q[0], q[1:]
		if o == nil {
			continue
		}
		keys = append(keys, o.key)
		q = append(q, o.lr[0], o.lr[1])
	}
	//sort.Ints(keys)
	return
}

//

func (o *tnode) String() string {
	var s string
	if o.value == 1 {
		s = Sprintf("%v", o.key)
	} else {
		s = Sprintf("%v(%v)", o.key, o.value)
	}
	s += Sprintf("[sz:%d,msz:%d]", o.sz, o.msz)
	return s
}

func (o *tnode) draw(prefix string, isTail bool, str *string) {
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

func (t *bst) String() string {
	if t.root == nil {
		return "BST (empty)\n"
	}
	str := "BST\n"
	t.root.draw("", true, &str)
	return str
}
