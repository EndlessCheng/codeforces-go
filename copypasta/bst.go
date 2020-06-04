package copypasta

import . "fmt"

// 二叉树常用函数

// 用 GoLand 的话强烈建议加入到 Live Templates 中，比赛时直接敲快捷键
type tKeyType int
type tValueType int

type bstNode struct {
	lr    [2]*bstNode
	sz    int
	msz   int
	key   tKeyType
	value tValueType
}

// 设置如下返回值是为了方便使用 bstNode 中的 lr 数组
func (o *bstNode) cmp(b tKeyType) int8 {
	switch {
	case b < o.key:
		return 0
	case b > o.key:
		return 1
	default:
		return -1
	}
}

func (o *bstNode) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *bstNode) mSize() int {
	if o != nil {
		return o.msz
	}
	return 0
}

type bst struct {
	root *bstNode
}

func (t *bst) size() int   { return t.root.size() }
func (t *bst) empty() bool { return t.root == nil }

func (t *bst) get(key tKeyType) *bstNode {
	for o := t.root; o != nil; {
		if c := o.cmp(key); c >= 0 {
			o = o.lr[c]
		} else {
			return o
		}
	}
	return nil
}

// max <= key
// return nil if not found
// same like --upper_bound in C++ STL
func (t *bst) floor(key tKeyType) (floor *bstNode) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.lr[0]
		case c > 0:
			floor = o
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

// min >= key
// return nil if not found
func (t *bst) lowerBound(key tKeyType) (lb *bstNode) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			lb = o
			o = o.lr[0]
		case c > 0:
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

// 前驱（小于 key，且最大的数）
func (t *bst) prev(key tKeyType) (prev *bstNode) {
	// 另一种写法
	// rank, _ := t.mRank(key)
	// return t.mSelect(rank - 1)
	for o := t.root; o != nil; {
		if o.cmp(key) <= 0 {
			o = o.lr[0]
		} else {
			prev = o
			o = o.lr[1]
		}
	}
	return
}

// 后继（大于 key，且最小的数)
func (t *bst) next(key tKeyType) (next *bstNode) {
	// 另一种写法
	// rank, o := t.mRank(key)
	// if o != nil {
	// 	 rank += int(o.value)
	// }
	// return t.mSelect(rank)
	for o := t.root; o != nil; {
		if o.cmp(key) != 0 {
			o = o.lr[1]
		} else {
			next = o
			o = o.lr[0]
		}
	}
	return
}

func (t *bst) hasValueInRange(l, r int) bool {
	o := t.lowerBound(tKeyType(l))
	return o != nil && int(o.key) <= r
}

// 小于 key 的键的数量
func (t *bst) mRank(key tKeyType) (cnt int) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.lr[0]
		case c > 0:
			cnt += int(o.value) + o.lr[0].mSize()
			o = o.lr[1]
		default:
			cnt += o.lr[0].mSize()
			return
		}
	}
	return
}

// >= key 的键的数量
func (t *bst) lowerCount(key tKeyType) (cnt int) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			cnt += int(o.value) + o.lr[1].mSize()
			o = o.lr[0]
		case c > 0:
			o = o.lr[1]
		default:
			cnt += int(o.value) + o.lr[1].mSize()
			return
		}
	}
	return
}

// 排名为 k 的节点 o（即有 k 个键小于 o.key）
func (t *bst) mSelect(k int) (o *bstNode) {
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

func (t *bst) min() (min *bstNode) {
	for o := t.root; o != nil; o = o.lr[0] {
		min = o
	}
	return
}

func (t *bst) max() (max *bstNode) {
	for o := t.root; o != nil; o = o.lr[1] {
		max = o
	}
	return
}

func (t *bst) keys() (keys []tKeyType) {
	var o *bstNode
	q := []*bstNode{t.root}
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

func (o *bstNode) String() string {
	var s string
	if o.value == 1 {
		s = Sprintf("%v", o.key)
	} else {
		s = Sprintf("%v(%v)", o.key, o.value)
	}
	s += Sprintf("[sz:%d,msz:%d]", o.sz, o.msz)
	return s
}

func (o *bstNode) draw(prefix string, isTail bool, str *string) {
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
