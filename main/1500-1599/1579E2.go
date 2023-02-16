package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

// https://space.bilibili.com/206214
type node79 struct {
	lr       [2]*node79
	priority uint
	key      int
	keyCnt   int
	subCnt   int
}

func (o *node79) size() int {
	if o != nil {
		return o.subCnt
	}
	return 0
}

func (o *node79) maintain() {
	o.subCnt = o.keyCnt + o.lr[0].size() + o.lr[1].size()
}

func (o *node79) rotate(d int) *node79 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap79 struct {
	rd   uint
	root *node79
}

func (t *treap79) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap79) size() int { return t.root.size() }

func (t *treap79) _put(o *node79, key int) *node79 {
	if o == nil {
		o = &node79{priority: t.fastRand(), key: key, keyCnt: 1}
	} else if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.keyCnt++
	}
	o.maintain()
	return o
}

func (t *treap79) put(key int) { t.root = t._put(t.root, key) }

func (t *treap79) _delete(o *node79, key int) *node79 {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		if o.keyCnt > 1 {
			o.keyCnt--
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
	}
	o.maintain()
	return o
}

func (t *treap79) delete(key int) { t.root = t._delete(t.root, key) }

func (o *node79) cmp(a int) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0
	}
	return 1
}

func (t *treap79) rank(key int) (kth int) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			o = o.lr[0]
		case c > 0:
			kth += o.lr[0].size() + o.keyCnt
			o = o.lr[1]
		default:
			kth += o.lr[0].size()
			return
		}
	}
	return
}

func CF1579E2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n, v int
	t := &treap79{rd: uint(time.Now().UnixNano())/2 + 1}
	for Fscan(in, &T); T > 0; T-- {
		ans := int64(0)
		Fscan(in, &n, &v)
		t.root = nil
		t.put(v)
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			ans += int64(min(t.rank(v), i-t.rank(v+1)))
			t.put(v)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1579E2(os.Stdin, os.Stdout) }
