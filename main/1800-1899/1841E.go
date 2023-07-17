package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

// https://space.bilibili.com/206214
type node41 struct {
	lr       [2]*node41
	priority uint
	key      int
	subCnt   int
}

func (o *node41) size() int {
	if o != nil {
		return o.subCnt // 汇总
	}
	return 0
}

func (o *node41) maintain() { o.subCnt = 1 + o.lr[0].size() + o.lr[1].size() }

func (o *node41) rotate(d int) *node41 {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap41 struct {
	rd   uint
	root *node41
}

func (t *treap41) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap41) size() int { return t.root.size() }

func (t *treap41) _put(o *node41, key int) *node41 {
	if o == nil {
		return &node41{priority: t.fastRand(), key: key, subCnt: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		// 相等
	}
	o.maintain()
	return o
}

func (t *treap41) put(key int) { t.root = t._put(t.root, key) }

func (t *treap41) _delete(o *node41, key int) *node41 {
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

func (t *treap41) delete(key int) { t.root = t._delete(t.root, key) }

func (o *node41) cmp(a int) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0
	}
	return 1
}

func (t *treap41) lowerBound(key int) (lb *node41) {
	for o := t.root; o != nil; {
		switch c := o.cmp(key); {
		case c == 0:
			lb = o
			o = o.lr[0]
		case c > 0:
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

func (t *treap41) prev(key int) (prev *node41) {
	for o := t.root; o != nil; {
		if o.cmp(key) <= 0 {
			o = o.lr[0]
		} else {
			prev = o
			o = o.lr[1]
		}
	}
	return
}

func CF1841E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var T, n, v int
	var m int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		col := make([][]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			col[v] = append(col[v], i)
		}

		cnt := make([]int, n+1)
		cnt[n] = n
		t :=  &treap41{rd: uint(time.Now().UnixNano())/2 + 1}
		t.put(-1)
		t.put(n)
		for i := n; i > 0; i-- {
			for _, j := range col[i] {
				r := t.lowerBound(j).key
				l := t.prev(r).key
				cnt[r-l-1] -= i
				cnt[j-l-1] += i
				cnt[r-j-1] += i
				t.put(j)
			}
		}

		ans := int64(0)
		Fscan(in, &m)
		for i := int64(n); i > 1 && m > 1; i-- {
			c := min(m/i, int64(cnt[i]))
			ans += c * (i - 1)
			m -= c * i
			cnt[i] -= int(c)
			if 1 < m && m < i && cnt[i] > 0 {
				ans += m - 1
				break
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1841E(os.Stdin, os.Stdout) }
