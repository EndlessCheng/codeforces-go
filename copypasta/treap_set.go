package copypasta

// 无重复元素 set

type sNode struct {
	lr       [2]*sNode
	priority uint
	key      int
}

func (o *sNode) rotate(d int) *sNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type sTreap struct {
	rd   uint
	root *sNode
}

func newSetTreap() *sTreap {
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

func (t *sTreap) _put(o *sNode, key int) *sNode {
	if o == nil {
		return &sNode{priority: t.fastRand(), key: key}
	}
	if cmp := t.compare(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], key)
		if o.lr[cmp].priority > o.priority {
			o = o.rotate(cmp ^ 1)
		}
	}
	return o
}

func (t *sTreap) put(key int) { t.root = t._put(t.root, key) }

func (t *sTreap) _delete(o *sNode, key int) *sNode {
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

func (t *sTreap) ceiling(key int) (ceiling *sNode) {
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

func (t *sTreap) hasValueInRange(l, r int) bool {
	o := t.ceiling(l)
	return o != nil && o.key <= r
}
