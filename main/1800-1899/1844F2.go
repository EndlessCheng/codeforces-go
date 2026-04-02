package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"slices"
	"time"
)

type nodeS44[K comparable] struct {
	son      [2]*nodeS44[K]
	priority uint
	key      K
	subSize  int
}

func (o *nodeS44[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeS44[K]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *nodeS44[K]) rotate(d int) *nodeS44[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapS44[K comparable] struct {
	rd         uint
	root       *nodeS44[K]
	comparator func(a, b K) int
}

func (t *treapS44[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapS44[K]) size() int   { return t.root.size() }
func (t *treapS44[K]) empty() bool { return t.size() == 0 }

func (t *treapS44[K]) _put(o *nodeS44[K], key K) *nodeS44[K] {
	if o == nil {
		o = &nodeS44[K]{priority: t.fastRand(), key: key}
	} else {
		c := t.comparator(key, o.key)
		if c != 0 {
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

func (t *treapS44[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *treapS44[K]) _delete(o *nodeS44[K], key K) *nodeS44[K] {
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
	o.maintain()
	return o
}

func (t *treapS44[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapS44[K]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else {
			kth += o.son[0].size()
			break
		}
	}
	return
}

func (t *treapS44[K]) kth(k int) (o *nodeS44[K]) {
	if k < 0 || k >= t.root.size() {
		return
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + 1
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return
}

func newSetWith44[K comparable](comp func(a, b K) int) *treapS44[K] {
	return &treapS44[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: comp,
	}
}

func cf1844F2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &c)
		a := make([]int, n+2)
		l := make([]int, n+2)
		r := make([]int, n+2)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}

		var ans []int
		if c >= 0 {
			slices.Sort(a[1:n+1])
			ans = a
		} else {
			slices.SortFunc(a[1:n+1], func(a, b int) int { return b - a })
			ans = make([]int, n+1)
			ans[1] = a[1]
			r[1] = 2
			l[n+1] = n
			for i := 2; i <= n; i++ {
				l[i] = i - 1
				r[i] = i + 1
			}

			type pair struct{ x, y int }
			s := newSetWith44[pair](func(a, b pair) int {return cmp.Or(a.x-b.x, a.y-b.y)})
			for i := 3; i < n; i++ {
				if a[r[i]]-a[l[i]] >= c {
					s.put(pair{a[i], i})
				}
			}

			for i := 2; i <= n; i++ {
				pos := r[1]
				idx := s.lowerBoundIndex(pair{ans[i-1] + c, 0})
				if idx < s.size() {
					pos = s.kth(idx).key.y
				}
				ans[i] = a[pos]
				s.delete(pair{a[pos], pos})
				x := l[pos]
				y := r[pos]
				r[x] = y
				l[y] = x
				if r[x] > n || a[r[x]]-a[l[x]] < c {
					s.delete(pair{a[x], x})
				}
				if l[y] < 2 || a[r[y]]-a[l[y]] < c {
					s.delete(pair{a[y], y})
				}
			}
		}

		for i := 1; i <= n; i++ {
			Fprint(out, ans[i], " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1844F2(bufio.NewReader(os.Stdin), os.Stdout) }
