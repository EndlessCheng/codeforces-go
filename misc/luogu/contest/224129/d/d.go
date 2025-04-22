package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"os"
	"time"
)

// https://github.com/EndlessCheng
type nodeMS[K comparable] struct {
	son      [2]*nodeMS[K]
	priority uint
	key      K
	keyCnt   int
	subSize  int
}

func (o *nodeMS[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeMS[K]) maintain() {
	o.subSize = o.keyCnt + o.son[0].size() + o.son[1].size()
}

func (o *nodeMS[K]) rotate(d int) *nodeMS[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapMS[K comparable] struct {
	rd         uint
	root       *nodeMS[K]
	comparator func(a, b K) int
}

func (t *treapMS[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapMS[K]) size() int   { return t.root.size() }
func (t *treapMS[K]) empty() bool { return t.size() == 0 }

func (t *treapMS[K]) _put(o *nodeMS[K], key K) *nodeMS[K] {
	if o == nil {
		o = &nodeMS[K]{priority: t.fastRand(), key: key, keyCnt: 1}
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

func (t *treapMS[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *treapMS[K]) _delete(o *nodeMS[K], key K) *nodeMS[K] {
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

func (t *treapMS[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapMS[K]) min() *nodeMS[K] { return t.kth(0) }
func (t *treapMS[K]) max() *nodeMS[K] { return t.kth(t.size() - 1) }

func (t *treapMS[K]) lowerBoundIndex(key K) (kth int) {
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

func (t *treapMS[K]) upperBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + o.keyCnt
			o = o.son[1]
		} else {
			kth += o.son[0].size() + o.keyCnt
			break
		}
	}
	return
}

func (t *treapMS[K]) kth(k int) (o *nodeMS[K]) {
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

func (t *treapMS[K]) prev(key K) *nodeMS[K] { return t.kth(t.lowerBoundIndex(key) - 1) }
func (t *treapMS[K]) next(key K) *nodeMS[K] { return t.kth(t.upperBoundIndex(key)) }

func (t *treapMS[K]) find(key K) *nodeMS[K] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

func newMultiset[K cmp.Ordered]() *treapMS[K] {
	return &treapMS[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: cmp.Compare[K],
	}
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	buf := make([]byte, 4096)
	_i, _n := 0, 0
	rc := func() byte {
		if _i == _n {
			_n, _ = in.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	rd := func() (x int) {
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n, m, k := rd(), rd(), rd()
	inE := make([][]int, m-k+2)
	outE := make([][]int, m-k+2)
	for ; n > 0; n-- {
		l, r, w := rd(), rd(), rd()
		if r-l+1 < k {
			continue
		}
		inE[r-k+1] = append(inE[r-k+1], w) // 最后可以游玩的分钟
		outE[l-1] = append(outE[l-1], w)
	}

	t := newMultiset[int]()
	f := make([]int, m+2)
	for i := m - k + 1; i > 0; i-- {
		for _, w := range inE[i] {
			t.put(w)
		}
		for _, w := range outE[i] {
			t.delete(w)
		}
		f[i] = f[i+1]
		if !t.empty() {
			f[i] = max(f[i], t.max().key+f[i+k])
		}
	}
	Fprintln(out, f[1])
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
