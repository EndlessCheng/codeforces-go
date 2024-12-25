package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type nodeMS00[K comparable] struct {
	son      [2]*nodeMS00[K]
	priority uint
	key      K
	keyCnt   int
	subSize  int
}

func (o *nodeMS00[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeMS00[K]) maintain() {
	o.subSize = o.keyCnt + o.son[0].size() + o.son[1].size()
}

func (o *nodeMS00[K]) rotate(d int) *nodeMS00[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapMS00[K comparable] struct {
	rd         uint
	root       *nodeMS00[K]
	comparator func(a, b K) int
}

func (t *treapMS00[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapMS00[K]) empty() bool { return t.root.size() == 0 }

func (t *treapMS00[K]) _put(o *nodeMS00[K], key K) *nodeMS00[K] {
	if o == nil {
		o = &nodeMS00[K]{priority: t.fastRand(), key: key, keyCnt: 1}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 {
			o.keyCnt++
		} else {
			d := 0
			if c > 0 {
				d = 1
			}
			o.son[d] = t._put(o.son[d], key)
			if o.son[d].priority > o.priority {
				o = o.rotate(d ^ 1)
			}
		}
	}
	o.maintain()
	return o
}

func (t *treapMS00[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *treapMS00[K]) _delete(o *nodeMS00[K], key K) *nodeMS00[K] {
	if o == nil {
		return nil
	}
	if c := t.comparator(key, o.key); c != 0 {
		d := 0
		if c > 0 {
			d = 1
		}
		o.son[d] = t._delete(o.son[d], key)
	} else {
		if o.keyCnt > 1 {
			o.keyCnt--
		} else {
			if o.son[1] == nil {
				return o.son[0]
			}
			if o.son[0] == nil {
				return o.son[1]
			}
			d := 0
			if o.son[0].priority > o.son[1].priority {
				d = 1
			}
			o = o.rotate(d)
			o.son[d] = t._delete(o.son[d], key)
		}
	}
	o.maintain()
	return o
}

func (t *treapMS00[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapMS00[K]) min() *nodeMS00[K] { return t.kth(0) }

func (t *treapMS00[K]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + o.keyCnt
			o = o.son[1]
		} else {
			kth += o.son[0].size()
			break
		}
	}
	return
}

func (t *treapMS00[K]) kth(k int) (o *nodeMS00[K]) {
	if k < 0 || k >= t.root.size() {
		return
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + o.keyCnt
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return
}

func newMultiset00[K cmp.Ordered]() *treapMS00[K] {
	return &treapMS00[K]{
		rd:         1,
		comparator: cmp.Compare[K],
	}
}

const stNodeDefaultVal00 = 1e9

var emptyStNode00 = &stNode00{val: stNodeDefaultVal00}

func init() {
	emptyStNode00.lo = emptyStNode00
	emptyStNode00.ro = emptyStNode00
}

type stNode00 struct {
	lo, ro    *stNode00
	l, r, val int
}

func (o *stNode00) update(i int, val int) {
	if o.l == o.r {
		o.val = val
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.lo == emptyStNode00 {
			o.lo = &stNode00{lo: emptyStNode00, ro: emptyStNode00, l: o.l, r: m, val: stNodeDefaultVal00}
		}
		o.lo.update(i, val)
	} else {
		if o.ro == emptyStNode00 {
			o.ro = &stNode00{lo: emptyStNode00, ro: emptyStNode00, l: m + 1, r: o.r, val: stNodeDefaultVal00}
		}
		o.ro.update(i, val)
	}
	o.val = min(o.lo.val, o.ro.val)
}

func (o *stNode00) query(l int) int {
	if o == emptyStNode00 || l > o.r {
		return stNodeDefaultVal00
	}
	if l <= o.l {
		return o.val
	}
	return min(o.lo.query(l), o.ro.query(l))
}

func newStRoot00(l, r int) *stNode00 {
	return &stNode00{lo: emptyStNode00, ro: emptyStNode00, l: l, r: r, val: stNodeDefaultVal00}
}

func cf2000H(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 4e6 + 1
	var T, n, v, m int
	var op string
	for Fscan(in, &T); T > 0; T-- {
		t := newStRoot00(1, mx)
		gap := map[int]*treapMS00[int]{}
		put := func(l, r int) {
			k := r - l - 1
			if k == 0 {
				return
			}
			if gap[k] == nil {
				gap[k] = newMultiset00[int]()
			}
			gap[k].put(l + 1)
			t.update(k, gap[k].min().key)
		}
		del := func(l, r int) {
			k := r - l - 1
			if k == 0 {
				return
			}
			t2 := gap[k]
			t2.delete(l + 1)
			if t2.empty() {
				t.update(k, 1e9)
			} else {
				t.update(k, t2.min().key)
			}
		}

		set := newMultiset00[int]()
		set.put(0)
		pre := 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			set.put(v)
			put(pre, v)
			pre = v
		}
		set.put(mx)
		put(pre, mx)

		for Fscan(in, &m); m > 0; m-- {
			Fscan(in, &op, &v)
			if op == "+" {
				i := set.lowerBoundIndex(v)
				l, r := set.kth(i-1).key, set.kth(i).key
				set.put(v)
				del(l, r)
				put(l, v)
				put(v, r)
			} else if op == "-" {
				i := set.lowerBoundIndex(v)
				l, r := set.kth(i-1).key, set.kth(i+1).key
				set.delete(v)
				del(l, v)
				del(v, r)
				put(l, r)
			} else {
				Fprint(out, t.query(v), " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2000H(bufio.NewReader(os.Stdin), os.Stdout) }
