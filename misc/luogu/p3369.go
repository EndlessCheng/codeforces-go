package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"time"
)

// https://space.bilibili.com/206214
type node3369 struct {
	son      [2]*node3369
	priority uint
	key      int32
	keyCnt   int
	subSize  int
}

func (o *node3369) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *node3369) maintain() { o.subSize = o.keyCnt + o.son[0].size() + o.son[1].size() }

func (o *node3369) rotate(d int) *node3369 {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap3369 struct {
	rd   uint
	root *node3369
}

func (t *treap3369) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap3369) size() int { return t.root.size() }

func (t *treap3369) _put(o *node3369, key int32) *node3369 {
	if o == nil {
		return &node3369{priority: t.fastRand(), key: key, keyCnt: 1, subSize: 1}
	}
	if c := o.cmp(key); c != 0 {
		d := (c + 1) / 2
		o.son[d] = t._put(o.son[d], key)
		if o.son[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else { // 相等
		o.keyCnt++
	}
	o.maintain()
	return o
}

func (t *treap3369) put(key int32) { t.root = t._put(t.root, key) }

func (t *treap3369) _delete(o *node3369, key int32) *node3369 {
	if o == nil {
		return nil
	}
	if c := o.cmp(key); c != 0 {
		d := (c + 1) / 2
		o.son[d] = t._delete(o.son[d], key)
	} else { // 相等
		if o.keyCnt > 1 {
			o.keyCnt--
		} else { // 删除
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

func (t *treap3369) delete(key int32) { t.root = t._delete(t.root, key) }

func newTreap3369() *treap3369 { return &treap3369{rd: uint(time.Now().UnixNano())/2 + 1} }

func (t *treap3369) min() *node3369 { return t.kth(0) }
func (t *treap3369) max() *node3369 { return t.kth(t.size() - 1) }

// < key 的元素个数
// 如果统计 > 改 cmp
func (t *treap3369) lowerBoundIndex(key int32) (kth int) {
	for o := t.root; o != nil; {
		c := o.cmp(key)
		if c == -1 {
			o = o.son[0]
		} else if c == 1 {
			kth += o.son[0].size() + o.keyCnt // 1 <-> o.keyCnt
			o = o.son[1]
		} else { // 相等
			kth += o.son[0].size()
			break
		}
	}
	return
}

// 第 k 小：有 k 个元素小于 o.key（k 从 0 开始）
// 也可以把 treap 当作一个有序数组，返回下标为 k 的 node
// 如果统计 > 改 cmp
func (t *treap3369) kth(k int) (o *node3369) {
	if k < 0 || k >= t.root.size() {
		return // NOTE: check nil
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + o.keyCnt // 1 <-> o.keyCnt
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return // NOTE: check nil
}

// -1 去左边，1 去右边
func (o *node3369) cmp(key int32) int {
	return cmp.Compare(key, o.key)
}

func p3369(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var q, op int
	var v int32
	Fscan(in, &q)

	t := newTreap3369()
	for range q {
		Fscan(in, &op, &v)
		switch op {
		case 1:
			t.put(v)
		case 2:
			t.delete(v)
		case 3:
			Fprintln(out, t.lowerBoundIndex(v)+1)
		case 4:
			Fprintln(out, t.kth(int(v-1)).key)
		case 5:
			Fprintln(out, t.kth(t.lowerBoundIndex(v)-1).key)
		default:
			Fprintln(out, t.kth(t.lowerBoundIndex(v+1)).key)
		}
	}
}

//func main() { p3369(bufio.NewReader(os.Stdin), os.Stdout) }
