package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
	"time"
)

// https://space.bilibili.com/206214
type node94[K comparable] struct {
	son      [2]*node94[K]
	priority uint
	key      K
	keyCnt   int
	subSize  int
	keySum   int
	subSum   int
}

func (o *node94[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *node94[K]) getSum() int {
	if o != nil {
		return o.subSum
	}
	return 0
}

func (o *node94[K]) maintain() {
	o.subSize = o.keyCnt + o.son[0].size() + o.son[1].size()
	o.subSum = o.keySum + o.son[0].getSum() + o.son[1].getSum()
}

func (o *node94[K]) rotate(d int) *node94[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap94[K comparable] struct {
	rd         uint
	root       *node94[K]
	comparator func(a, b K) int
	keyToInt   func(key K) int
}

func (t *treap94[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap94[K]) size() int   { return t.root.size() }
func (t *treap94[K]) empty() bool { return t.size() == 0 }

func (t *treap94[K]) _put(o *node94[K], key K, num int) *node94[K] {
	if o == nil {
		if num < 0 {
			panic(num)
		}
		o = &node94[K]{priority: t.fastRand(), key: key, keyCnt: num, keySum: t.keyToInt(key) * num}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 {
			o.keyCnt += num
			if o.keyCnt < 0 {
				panic(o.keyCnt)
			}
			o.keySum += t.keyToInt(key) * num
		} else {
			d := 0
			if c > 0 {
				d = 1
			}
			o.son[d] = t._put(o.son[d], key, num)
			if o.son[d].priority > o.priority {
				o = o.rotate(d ^ 1)
			}
		}
	}
	o.maintain()
	return o
}

// num=1 添加一个 key
// num=-1 移除一个 key
func (t *treap94[K]) put(key K, num int) { t.root = t._put(t.root, key, num) }

func (t *treap94[K]) min() *node94[K] { return t.kth(0) }
func (t *treap94[K]) max() *node94[K] { return t.kth(t.size() - 1) }

func (t *treap94[K]) lowerBoundIndex(key K) (kth int) {
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

func (t *treap94[K]) upperBoundIndex(key K) (kth int) {
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

func (t *treap94[K]) kth(k int) (o *node94[K]) {
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

func (t *treap94[K]) prev(key K) *node94[K] { return t.kth(t.lowerBoundIndex(key) - 1) }
func (t *treap94[K]) next(key K) *node94[K] { return t.kth(t.upperBoundIndex(key)) }

func (t *treap94[K]) find(key K) *node94[K] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

// 返回 [0,k) 的元素和
func (t *treap94[K]) preSum(k int) (sum int) {
	if k < 0 {
		//return
		panic(k)
	}
	if k > t.root.size() {
		//return t.root.getSum()
		panic(k)
	}
	for o := t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
			continue
		}
		sum += o.son[0].getSum()
		k -= leftSize
		if k <= o.keyCnt {
			sum += t.keyToInt(o.key) * k
			break
		}
		sum += o.keySum
		k -= o.keyCnt
		o = o.son[1]
	}
	return
}

func (t *treap94[K]) sumLess(x K) int         { return t.preSum(t.lowerBoundIndex(x)) }
func (t *treap94[K]) sumLessEqual(x K) int    { return t.preSum(t.upperBoundIndex(x)) }
func (t *treap94[K]) sumGreater(x K) int      { return t.root.getSum() - t.preSum(t.upperBoundIndex(x)) }
func (t *treap94[K]) sumGreaterEqual(x K) int { return t.root.getSum() - t.preSum(t.lowerBoundIndex(x)) }

func newTreap94() *treap94[int] {
	return &treap94[int]{
		rd:         uint(time.Now().UnixNano()),
		comparator: cmp.Compare[int],
		keyToInt:   func(key int) int { return key },
	}
}

func p5094(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	type pair struct{ v, x int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v, &a[i].x)
	}

	slices.SortFunc(a, func(a, b pair) int { return a.v - b.v })
	t := newTreap94()
	for _, p := range a {
		i := t.lowerBoundIndex(p.x)
		ans += p.v * (t.preSum(t.size()) - t.preSum(i)*2 - (t.size()-i*2)*p.x)
		t.put(p.x, 1)
	}
	Fprint(out, ans)
}

//func main() { p5094(bufio.NewReader(os.Stdin), os.Stdout) }
