package copypasta

// 无重复 set

type stpNode struct {
	lr       [2]*stpNode
	priority uint
	key      int
}

func (o *stpNode) rotate(d int) *stpNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type sTreap struct {
	rd   uint
	root *stpNode
}

func newSimpleTreap() *sTreap {
	return &sTreap{rd: 1}
}

func (t *sTreap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *sTreap) compare(a, b int) int {
	switch {
	case a < b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func (t *sTreap) _put(o *stpNode, key int) *stpNode {
	if o == nil {
		return &stpNode{priority: t.fastRand(), key: key}
	}
	cmp := t.compare(key, o.key)
	o.lr[cmp] = t._put(o.lr[cmp], key)
	if o.lr[cmp].priority > o.priority {
		o = o.rotate(cmp ^ 1)
	}
	return o
}

func (t *sTreap) put(key int) { t.root = t._put(t.root, key) }

func (t *sTreap) _delete(o *stpNode, key int) *stpNode {
	if o == nil {
		return nil
	}
	if cmp := t.compare(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._delete(o.lr[cmp], key)
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
	return o
}

func (t *sTreap) delete(key int) { t.root = t._delete(t.root, key) }
