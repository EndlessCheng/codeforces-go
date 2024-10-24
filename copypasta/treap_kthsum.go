package copypasta

import "time"

// 已废弃！请移步新版本泛型代码 ./treap/prefixsum/prefixsum.go

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

func newTreapSum() *treapSum {
	return &treapSum{rd: uint(time.Now().UnixNano())/2 + 1}
}

// <= size 的元素个数，元素和
// LC3245 https://leetcode.cn/problems/alternating-groups-iii/
func (t *treapSum) cntSum(size int) (cnt, sum int) {
	for o := t.root; o != nil; {
		c := o.cmp(size)
		if c == 0 { // size 小，去左子树找
			o = o.lr[0]
			continue
		}
		// 左子树 + 自己
		cnt += o.lr[0].getSize() + o.keyCnt
		sum += o.lr[0].getSum() + o.keySum
		if c < 0 { // 相等
			return
		}
		o = o.lr[1]
	}
	return
}

// 返回前 k 小数的和（k 从 1 开始）
// https://www.codechef.com/problems/OKLAMA
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

// 从大到小，计算凑出 need 至少需要多少个数
// cmp 需要改成 >
// 来自 https://codeforces.com/contest/1978/problem/D 的麻烦写法
func (t *treapSum) rank(need int) (cnt int) {
	for o := t.root; o != nil; {
		if o.lr[0].getSum() >= need {
			o = o.lr[0]
		} else {
			need -= o.lr[0].getSum()
			cnt += o.lr[0].getSize()
			if o.keyCnt*o.key >= need {
				cnt += (need + o.key - 1) / o.key
				return
			}
			need -= o.keyCnt * o.key
			o = o.lr[1]
		}
	}
	panic(-1)
}
