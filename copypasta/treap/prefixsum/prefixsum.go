package multiset

import (
	"cmp"
	"time"
)

/*
基于 multiset
把 treap 当作一个有序数组
额外维护这个数组的前缀和

LC3013 https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-ii/
LC3321 https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-ii/
https://www.codechef.com/problems/OKLAMA
https://atcoder.jp/contests/abc306/tasks/abc306_e
https://atcoder.jp/contests/abc287/tasks/abc287_g
*/

type node[K comparable] struct {
	son      [2]*node[K]
	priority uint
	key      K
	keyCnt   int
	subSize  int
	keySum   int
	subSum   int
}

func (o *node[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *node[K]) getSum() int {
	if o != nil {
		return o.subSum
	}
	return 0
}

func (o *node[K]) maintain() {
	o.subSize = o.keyCnt + o.son[0].size() + o.son[1].size()
	o.subSum = o.keySum + o.son[0].getSum() + o.son[1].getSum()
}

func (o *node[K]) rotate(d int) *node[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap[K comparable] struct {
	rd         uint
	root       *node[K]
	comparator func(a, b K) int
	keyToInt   func(key K) int
}

func (t *treap[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap[K]) size() int   { return t.root.size() }
func (t *treap[K]) empty() bool { return t.size() == 0 }

func (t *treap[K]) _put(o *node[K], key K, num int) *node[K] {
	if o == nil {
		if num < 0 {
			panic(num)
		}
		o = &node[K]{priority: t.fastRand(), key: key, keyCnt: num, keySum: t.keyToInt(key) * num}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 { // 相等
			o.keyCnt += num
			if o.keyCnt < 0 {
				panic(o.keyCnt)
			}
			o.keySum += t.keyToInt(key) * num
		} else {
			// < 0 去左边，> 0 去右边
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

// num=1 表示添加一个 key
// num=-1 表示移除一个 key
func (t *treap[K]) put(key K, num int) { t.root = t._put(t.root, key, num) }

func (t *treap[K]) min() *node[K] { return t.kth(0) }
func (t *treap[K]) max() *node[K] { return t.kth(t.size() - 1) }

// 把 treap 当作一个有序数组，返回第一个 >= key 的数的下标，若不存在，返回 size()
// 等价于 < key 的元素个数
func (t *treap[K]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + o.keyCnt
			o = o.son[1]
		} else { // 相等
			kth += o.son[0].size()
			break
		}
	}
	return
}

// 把 treap 当作一个有序数组，返回第一个 > key 的数的下标，若不存在，返回 size()
// 等价于 <= key 的元素个数
func (t *treap[K]) upperBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + o.keyCnt
			o = o.son[1]
		} else { // 相等
			kth += o.son[0].size() + o.keyCnt
			break
		}
	}
	return
}

// 把 treap 当作一个有序数组，返回下标为 k 的 node（k 从 0 开始）
// 也就是第 k+1 小：node.key 是最小的满足「有至少 k+1 个元素 <= node.key」的元素
func (t *treap[K]) kth(k int) (o *node[K]) {
	if k < 0 || k >= t.root.size() {
		return // NOTE: check nil
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
	return // NOTE: check nil
}

// 求 v 的前驱（小于 key 的最大元素）
func (t *treap[K]) prev(key K) *node[K] { return t.kth(t.lowerBoundIndex(key) - 1) }

// 求 v 的后继（大于 key 的最小元素）
func (t *treap[K]) next(key K) *node[K] { return t.kth(t.upperBoundIndex(key)) }

// <= key 可以用 t.kth(t.upperBoundIndex(key)-1)
// >= key 可以用 t.kth(t.lowerBoundIndex(key))

func (t *treap[K]) find(key K) *node[K] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

// 把 treap 当作一个有序数组，返回左闭右开区间 [0,k) 的元素和
// 特别地，k = size() 表示整棵树的元素和
func (t *treap[K]) preSum(k int) (sum int) {
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

// < x 的元素和
func (t *treap[K]) sumLess(x K) int { return t.preSum(t.lowerBoundIndex(x)) }

// <= x 的元素和
func (t *treap[K]) sumLessEqual(x K) int { return t.preSum(t.upperBoundIndex(x)) }

// > x 的元素和
func (t *treap[K]) sumGreater(x K) int { return t.root.getSum() - t.preSum(t.upperBoundIndex(x)) }

// >= x 的元素和
func (t *treap[K]) sumGreaterEqual(x K) int { return t.root.getSum() - t.preSum(t.lowerBoundIndex(x)) }

func newTreap() *treap[int] {
	return &treap[int]{
		rd:         uint(time.Now().UnixNano())/2 + 1,
		comparator: cmp.Compare[int],
		keyToInt:   func(key int) int { return key },
	}
}

func newTreapWith[K comparable](comparator func(a, b K) int, keyToInt func(key K) int) *treap[K] {
	return &treap[K]{
		rd:         uint(time.Now().UnixNano())/2 + 1,
		comparator: comparator,
		keyToInt:   keyToInt,
	}
}

// 具体用法见本目录下的 prefixsum_example.go

// 旧版写法见外面的 treap_kthsum.go
