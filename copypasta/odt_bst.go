package copypasta

type odtNode struct {
	tpNode
	l, r int
}

func (t *treap) floor(key int) (floor *odtNode) { return }
func (t *treap) next(key int) (next *odtNode)   { return }

func (t *treap) split(mid int) {
	if o := t.floor(mid); o.l < mid && mid <= o.r {
		//r, value := o.r, o.value
		//o.r = mid - 1
		//t.put(mid, r, value)
	}
}

func (t *treap) prepare(l, r int) {
	t.split(l)
	t.split(r + 1)
}

func (t *treap) merge(l, r int, value tpValueType) {
	//t.prepare(l, r)
	for o := t.next(l); o != nil && o.l <= r; o = t.next(o.l) {
		//t.delete(o.l)
	}
	o := t.floor(l)
	o.r = r
	o.value = value
}
