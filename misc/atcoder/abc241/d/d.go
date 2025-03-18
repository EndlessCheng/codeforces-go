package main

import (
	"bufio"
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

func (t *treapMS[K]) size() int { return t.root.size() }

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

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	t := &treapMS[int]{
		rd:         uint(time.Now().UnixNano())/2 + 1,
		comparator: func(a, b int) int { return a - b },
	}
	var q, op, x, k int
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op, &x)
		if op == 1 {
			t.put(x)
			continue
		}
		Fscan(in, &k)
		if k > t.size() {
			Fprintln(out, -1)
			continue
		}
		if op == 2 {
			o := t.kth(t.lowerBoundIndex(x+1) - k)
			if o == nil {
				Fprintln(out, -1)
			} else {
				Fprintln(out, o.key)
			}
		} else {
			o := t.kth(t.lowerBoundIndex(x) + k - 1)
			if o == nil {
				Fprintln(out, -1)
			} else {
				Fprintln(out, o.key)
			}
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
