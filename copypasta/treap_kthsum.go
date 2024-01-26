package copypasta

import "time"

/* 维护前 k 小元素和的 treap
支持添加删除元素

https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-ii/
https://atcoder.jp/contests/abc306/tasks/abc306_e
https://atcoder.jp/contests/abc287/tasks/abc287_g

*/
type nodeSum struct {
	lr       [2]*nodeSum
	priority uint
	key      int
	keyCnt   int
	subSize  int
	keySum   int
	subSum   int
}

func (o *nodeSum) cmp(a int) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0
	}
	return 1
}

func (o *nodeSum) getSize() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeSum) getSum() int {
	if o != nil {
		return o.subSum
	}
	return 0
}

func (o *nodeSum) maintain() {
	o.subSize = o.keyCnt + o.lr[0].getSize() + o.lr[1].getSize()
	o.subSum = o.keySum + o.lr[0].getSum() + o.lr[1].getSum()
}

func (o *nodeSum) rotate(d int) *nodeSum {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapSum struct {
	rd   uint
	root *nodeSum
}

func (t *treapSum) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapSum) _put(o *nodeSum, key, num int) *nodeSum {
	if o == nil {
		o = &nodeSum{priority: t.fastRand(), key: key, keyCnt: num, keySum: key * num}
	} else if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, num)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.keyCnt += num
		o.keySum += key * num
	}
	o.maintain()
	return o
}

// num=1 表示添加一个 key
// num=-1 表示移除一个 key
func (t *treapSum) put(key, num int) { t.root = t._put(t.root, key, num) }

// 返回前 k 小数的和（k 从 1 开始）
func (t *treapSum) kth(k int) (sum int) {
	if k > t.root.getSize() {
		panic(-1)
	}
	for o := t.root; o != nil; {
		if ls := o.lr[0].getSize(); k < ls {
			o = o.lr[0]
		} else {
			sum += o.lr[0].getSum()
			k -= ls
			if k <= o.keyCnt {
				sum += o.key * k
				return
			}
			sum += o.keySum
			k -= o.keyCnt
			o = o.lr[1]
		}
	}
	return
}

func newTreapSum() *treapSum { 
	return &treapSum{rd: uint(time.Now().UnixNano())/2 + 1} 
}
