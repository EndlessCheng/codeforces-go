package copypasta

// 允许重复元素 set

type msNode struct {
	lr       [2]*msNode
	priority uint
	key      int
	dupCnt   int
}

func (o *msNode) rotate(d int) *msNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type msTreap struct {
	rd   uint
	root *msNode
}

func newMultiSetTreap() *msTreap {
	return &msTreap{rd: 1}
}

func (t *msTreap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *msTreap) compare(a, b int) int {
	switch {
	case a < b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func (t *msTreap) _put(o *msNode, key int) *msNode {
	if o == nil {
		return &msNode{priority: t.fastRand(), key: key, dupCnt: 1}
	}
	if cmp := t.compare(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], key)
		if o.lr[cmp].priority > o.priority {
			o = o.rotate(cmp ^ 1)
		}
	} else {
		o.dupCnt++
	}
	return o
}

func (t *msTreap) put(key int) { t.root = t._put(t.root, key) }

func (t *msTreap) _delete(o *msNode, key int) *msNode {
	if o == nil {
		return nil
	}
	if cmp := t.compare(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._delete(o.lr[cmp], key)
	} else {
		if o.dupCnt > 1 {
			o.dupCnt--
		} else {
			if o.lr[1] == nil {
				return o.lr[0]
			}
			if o.lr[0] == nil {
				return o.lr[1]
			}
			cmp2 := 0
			if o.lr[0].priority > o.lr[1].priority {
				cmp2 = 1
			}
			o = o.rotate(cmp2)
			o.lr[cmp2] = t._delete(o.lr[cmp2], key)
		}
	}
	return o
}

func (t *msTreap) delete(key int) { t.root = t._delete(t.root, key) }
