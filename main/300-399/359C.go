package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

// https://space.bilibili.com/206214
type node59 struct {
	lr       [2]*node59
	priority uint
	key      int64
	value    int
	subCnt   int
}

func (o *node59) size() int {
	if o != nil {
		return o.subCnt
	}
	return 0
}

func (o *node59) maintain() { o.subCnt = 1 + o.lr[0].size() + o.lr[1].size() }

func (o *node59) rotate(d int) *node59 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap59 struct {
	rd   uint
	root *node59
}

func (t *treap59) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap59) size() int { return t.root.size() }

func (t *treap59) _put(o *node59, key int64, value int) *node59 {
	if o == nil {
		return &node59{priority: t.fastRand(), key: key, value: value, subCnt: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, value)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.value += value
	}
	o.maintain()
	return o
}

func (t *treap59) put(key int64, value int) { t.root = t._put(t.root, key, value) }

func (t *treap59) _delete(o *node59, key int64) *node59 {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		if o.lr[1] == nil {
			return o.lr[0]
		}
		if o.lr[0] == nil {
			return o.lr[1]
		}
		d = 0
		if o.lr[0].priority > o.lr[1].priority {
			d = 1
		}
		o = o.rotate(d)
		o.lr[d] = t._delete(o.lr[d], key)
	}
	o.maintain()
	return o
}

func (t *treap59) delete(key int64) { t.root = t._delete(t.root, key) }

func (o *node59) cmp(a int64) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0
	}
	return 1
}

func (t *treap59) min() (min *node59) {
	for o := t.root; o != nil; o = o.lr[0] {
		min = o
	}
	return
}

func CF359C(_r io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	pow := func(x, n int64) (res int64) {
		x %= mod
		res = 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}
	in := bufio.NewReader(_r)
	var n, x int
	s := int64(0)
	Fscan(in, &n, &x)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}

	t := &treap59{rd: uint(time.Now().UnixNano())/2 + 1}
	for _, v := range a {
		t.put(s-v, 1)
	}

	ans := int64(1)
	for t.size() > 0 {
		top := t.min()
		k, e := top.value, top.key
		if k%x > 0 {
			ans = pow(int64(x), min(s, e))
			break
		}
		e2 := int64(0)
		for k%x == 0 {
			k /= x
			e2++
		}
		t.delete(e)
		t.put(e+e2, k)
	}
	Fprint(out, ans)
}

//func main() { CF359C(os.Stdin, os.Stdout) }
