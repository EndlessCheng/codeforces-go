package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"time"
)

// https://github.com/EndlessCheng
type nodeS[K comparable] struct {
	son      [2]*nodeS[K]
	priority uint
	key      K
	subSize  int
}

func (o *nodeS[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeS[K]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *nodeS[K]) rotate(d int) *nodeS[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapS[K comparable] struct {
	rd         uint
	root       *nodeS[K]
	comparator func(a, b K) int
}

func (t *treapS[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapS[K]) size() int   { return t.root.size() }
func (t *treapS[K]) empty() bool { return t.size() == 0 }

func (t *treapS[K]) _put(o *nodeS[K], key K) *nodeS[K] {
	if o == nil {
		o = &nodeS[K]{priority: t.fastRand(), key: key}
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

func (t *treapS[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *treapS[K]) _delete(o *nodeS[K], key K) *nodeS[K] {
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

func (t *treapS[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapS[K]) min() *nodeS[K] { return t.kth(0) }
func (t *treapS[K]) max() *nodeS[K] { return t.kth(t.size() - 1) }

func (t *treapS[K]) lowerBoundIndex(key K) (kth int) {
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

func (t *treapS[K]) upperBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else {
			kth += o.son[0].size() + 1
			break
		}
	}
	return
}

func (t *treapS[K]) kth(k int) (o *nodeS[K]) {
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

func (t *treapS[K]) prev(key K) *nodeS[K] { return t.kth(t.lowerBoundIndex(key) - 1) }
func (t *treapS[K]) next(key K) *nodeS[K] { return t.kth(t.upperBoundIndex(key)) }

func (t *treapS[K]) find(key K) *nodeS[K] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

func run2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ k, d int }
		var a, b []pair
		ans := 0
		for i := 0; i < n; i++ {
			Fscan(in, &k, &l, &r)
			if l > r {
				ans += r
				a = append(a, pair{k, l - r})
			} else {
				ans += l
				b = append(b, pair{n - k, r - l})
			}
		}
		f := func(a []pair) {
			t := &treapS[int]{
				rd:         uint(time.Now().UnixNano())/2 + 1,
				comparator: func(a, b int) int { return a - b },
			}
			for i := 1; i <= n; i++ {
				t.put(i)
			}
			sort.Slice(a, func(i, j int) bool { return a[i].d > a[j].d })
			for _, p := range a {
				o := t.kth(t.upperBoundIndex(p.k) - 1)
				if o != nil {
					ans += p.d
					t.delete(o.key)
				}
			}
		}
		f(a)
		f(b)
		Fprintln(out, ans)
	}
}

//func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
