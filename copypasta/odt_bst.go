package copypasta

// https://codeforces.com/problemset/problem/558/E
// 代码 https://codeforces.com/problemset/submission/558/117163317
// https://codeforces.com/problemset/problem/915/E
// 代码 https://codeforces.com/problemset/submission/915/117158161
// todo https://www.luogu.com.cn/problem/P5350
//      https://www.luogu.com.cn/problem/P5586

type odtNode struct {
	tpNode
	l, r int
}

func (t *treap) put1(l, r int, val tpValueType) {}
func (t *treap) floor(int) (_ *odtNode)         { return }
func (t *treap) next(int) (_ *odtNode)          { return }

func (t *treap) split(mid int) {
	if o := t.floor(mid); o.l < mid && mid <= o.r {
		r, val := o.r, o.val
		o.r = mid - 1
		t.put1(mid, r, val)
	}
}

func (t *treap) prepare(l, r int) {
	t.split(l)
	t.split(r + 1)
}

func (t *treap) merge(l, r int, value tpValueType) {
	t.prepare(l, r)
	for o := t.next(l); o != nil && o.l <= r; o = t.next(o.l) {
		t.delete(tpKeyType(o.l))
	}
	o := t.floor(l)
	o.r = r
	o.val = value
}

// https://codeforces.com/problemset/problem/558/E
func (t *treap) sortBytes(l, r int, inc bool) {
	t.prepare(l, r)
	// 整段删除重建
	cnt := [26]int{}
	for o := t.floor(l); o != nil && o.l <= r; o = t.next(o.l) {
		cnt[o.val] += o.r - o.l + 1
		t.delete(tpKeyType(o.l))
	}
	if inc {
		for i, c := range cnt {
			if c > 0 {
				t.put1(l, l+c-1, tpValueType(i))
				l += c
			}
		}
	} else {
		for i, c := range cnt {
			if c > 0 {
				t.put1(r-c+1, r, tpValueType(i))
				r -= c
			}
		}
	}
}
