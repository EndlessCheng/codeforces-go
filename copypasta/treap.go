package copypasta

import (
	"cmp"
	. "fmt"
	"strings"
	"time"
)

// 已废弃！请移步新版本泛型代码 ./treap/

/* 树堆 treap=tree+heap
本质上属于笛卡尔树，见 cartesian_tree.go

https://oi-wiki.org/ds/treap/
https://en.wikipedia.org/wiki/Treap
复杂度证明 http://www.cs.cmu.edu/afs/cs/academic/class/15210-s12/www/lectures/lecture16.pdf
部分代码参考刘汝佳实现，见 https://github.com/klb3713/aoapc-book/blob/master/TrainingGuide/bookcodes/ch3/la5031.cpp
额外维护子树和的写法见 https://codeforces.com/contest/1398/submission/119651187
todo Merging treaps https://codeforces.com/blog/entry/108601

模板题 https://www.luogu.com.cn/problem/P3369
      https://www.luogu.com.cn/problem/P6136
https://atcoder.jp/contests/abc241/tasks/abc241_d
题目推荐 https://cp-algorithms.com/data_structures/treap.html#toc-tgt-8
https://codeforces.com/problemset/problem/85/D 较为复杂的维护
https://atcoder.jp/contests/abc245/tasks/abc245_e 离线+lowerbound+delete
https://atcoder.jp/contests/abc356/tasks/abc356_f 维护关键位置
https://atcoder.jp/contests/aising2020/tasks/aising2020_e
*/

// 用 GoLand 的话强烈建议加入到 Live Templates 中，比赛时直接敲快捷键
type tpKeyType int
type tpValueType int

type tpNode struct {
	lr       [2]*tpNode
	priority uint // max heap
	sz       int
	msz      int
	key      tpKeyType
	val      tpValueType // dupCnt for multiset
}

// 设置如下返回值是为了方便使用 tpNode 中的 lr 数组
func (o *tpNode) cmp(a tpKeyType) int {
	b := o.key
	if a == b {
		return -1
	}
	if a < b {
		return 0 // 左儿子
	}
	return 1 // 右儿子
}

func (o *tpNode) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *tpNode) mSize() int {
	if o != nil {
		return o.msz
	}
	return 0
}

// 对于取名叫 maintain 还是 pushUp，由于操作的对象是当前节点，个人认为取名 maintain 更为准确
func (o *tpNode) maintain() {
	o.sz = 1 + o.lr[0].size() + o.lr[1].size()
	o.msz = int(o.val) + o.lr[0].mSize() + o.lr[1].mSize()
}

// 旋转，并维护子树大小
// d=0：左旋，返回 o 的右儿子
// d=1：右旋，返回 o 的左儿子
func (o *tpNode) rotate(d int) *tpNode {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	// x.sz = o.sz; x.msz = o.msz; o.maintain()
	o.maintain()
	x.maintain()
	return x
}

type treap struct {
	rd   uint
	root *tpNode
}

// 也可以直接设 rd 为 1
func newTreap() *treap {
	return &treap{rd: uint(time.Now().UnixNano())}
}

// https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
// https://en.wikipedia.org/wiki/Xorshift
// 任何 Go 版本都通用的写法
func (t *treap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

// 插入一键值对，返回插入后优先级最大的节点
// 先和二叉搜索树的插入一样，先把要插入的点插入到一个叶子上，并随机分配一个优先级，
// 然后跟维护堆一样，如果当前节点的优先级比根大就旋转，如果当前节点是根的左儿子就右旋如果当前节点是根的右儿子就左旋
func (t *treap) _put(o *tpNode, key tpKeyType, val tpValueType) *tpNode {
	if o == nil {
		return &tpNode{priority: t.fastRand(), sz: 1, msz: 1, key: key, val: val}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key, val)
		// 优先级比根大就旋转
		if o.lr[d].priority > o.priority {
			// 是根的左儿子就右旋，反之左旋
			o = o.rotate(d ^ 1)
		}
	} else {
		//o.val = val
		o.val += val
	}
	o.maintain()
	return o
}

func (t *treap) put(key tpKeyType, val tpValueType) { t.root = t._put(t.root, key, val) }

// 删除一个键，返回删除后优先级最大的节点，若无节点返回 nil
// 因为 treap 满足堆性质，所以只需要把要删除的节点旋转到叶节点上，然后直接删除就可以了
// 具体的方法就是每次找到优先级最大的儿子，向与其相反的方向旋转，这样要删除的节点会不断下降直到叶节点，然后直接删除
func (t *treap) _delete(o *tpNode, key tpKeyType) *tpNode {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		if o.val > 1 {
			o.val--
		} else {
			if o.lr[1] == nil {
				return o.lr[0]
			}
			if o.lr[0] == nil {
				return o.lr[1]
			}
			// o 有两棵子树，把优先级高的子树旋转到根，然后递归在另一棵子树中删除 o
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

func (t *treap) delete(key tpKeyType) { t.root = t._delete(t.root, key) }

// 其余通用方法，例如 get/find/prev/next/min/max 见 bst.go

//

func (o *tpNode) String() (s string) {
	//return strconv.Itoa(int(o.key))
	if o.val == 1 {
		s = Sprintf("%v", o.key)
	} else {
		s = Sprintf("%v(%v)", o.key, o.val)
	}
	s += Sprintf("[sz:%d,msz:%d]", o.sz, o.msz)
	return
}

/* 逆时针旋转 90° 打印这棵树：根节点在最左侧，右子树在上侧，左子树在下侧

效果如下（只打印 key）

Root
│           ┌── 95
│       ┌── 94
│   ┌── 90
│   │   │           ┌── 89
│   │   │       ┌── 88
│   │   │       │   └── 87
│   │   │       │       └── 81
│   │   │   ┌── 74
│   │   └── 66
└── 62
    │           ┌── 59
    │       ┌── 58
    │       │   └── 56
    │       │       └── 47
    │   ┌── 45
    └── 40
        │       ┌── 37
        │   ┌── 28
        └── 25
            │           ┌── 18
            │       ┌── 15
            │   ┌── 11
            └── 6
                └── 0

*/
func (o *tpNode) draw(treeSB, prefixSB *strings.Builder, isTail bool) {
	prefix := prefixSB.String()
	if o.lr[1] != nil {
		newPrefixSB := &strings.Builder{}
		newPrefixSB.WriteString(prefix)
		if isTail {
			newPrefixSB.WriteString("│   ")
		} else {
			newPrefixSB.WriteString("    ")
		}
		o.lr[1].draw(treeSB, newPrefixSB, false)
	}
	treeSB.WriteString(prefix)
	if isTail {
		treeSB.WriteString("└── ")
	} else {
		treeSB.WriteString("┌── ")
	}
	treeSB.WriteString(o.String())
	treeSB.WriteByte('\n')
	if o.lr[0] != nil {
		newPrefixSB := &strings.Builder{}
		newPrefixSB.WriteString(prefix)
		if isTail {
			newPrefixSB.WriteString("    ")
		} else {
			newPrefixSB.WriteString("│   ")
		}
		o.lr[0].draw(treeSB, newPrefixSB, true)
	}
}

func (t *treap) String() string {
	if t.root == nil {
		return "Empty\n"
	}
	treeSB := &strings.Builder{}
	treeSB.WriteString("Root\n")
	t.root.draw(treeSB, &strings.Builder{}, true)
	return treeSB.String()
}

// 占位用，实际代码见 treap 文件夹

type nodeM[K comparable, V any] struct {
	son      [2]*nodeM[K, V]
	priority uint
	key      K
	value    V
	subSize  int
}

func (o *nodeM[K, V]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

type treapM[K comparable, V any] struct {
	rd         uint
	root       *nodeM[K, V]
	comparator func(a, b K) int
}

func (t *treapM[K, V]) size() int                       { return t.root.size() }
func (t *treapM[K, V]) empty() bool                     { return t.size() == 0 }
func (t *treapM[K, V]) put(key K, value V)              {}
func (t *treapM[K, V]) delete(key K)                    {}
func (t *treapM[K, V]) min() *nodeM[K, V]               { return t.kth(0) }
func (t *treapM[K, V]) max() *nodeM[K, V]               { return t.kth(t.size() - 1) }
func (t *treapM[K, V]) lowerBoundIndex(key K) (kth int) { return }
func (t *treapM[K, V]) upperBoundIndex(key K) (kth int) { return }
func (t *treapM[K, V]) kth(k int) (o *nodeM[K, V])      { return }
func (t *treapM[K, V]) prev(key K) *nodeM[K, V]         { return t.kth(t.lowerBoundIndex(key) - 1) }
func (t *treapM[K, V]) next(key K) *nodeM[K, V]         { return t.kth(t.upperBoundIndex(key)) }
func (t *treapM[K, V]) find(key K) (o *nodeM[K, V])     { return }

func newMap[K cmp.Ordered, V any]() *treapM[K, V]                         { return nil }
func newMapWith[K comparable, V any](comp func(a, b K) int) *treapM[K, V] { return nil }

//

type treapS[K comparable] struct{}

func (t *treapS[K]) lowerBoundPreSum(lowS int) (cnt, sum int) { return }
func (t *treapS[K]) put(key K, num int)                       {}
func newTreapWith[K comparable](comp func(a, b K) int, keyToInt func(key K) int) *treapS[K] {
	return nil
}
