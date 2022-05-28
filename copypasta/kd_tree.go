package copypasta

import (
	"math/rand"
	"time"
)

/* k-d tree: k-dimensional tree; k 维树
https://en.wikipedia.org/wiki/K-d_tree

推荐 https://www.luogu.com.cn/blog/command-block/kdt-xiao-ji
https://www.luogu.com.cn/blog/lc-2018-Canton/solution-p4148
https://oi-wiki.org/ds/kdt/

todo 题单 https://www.luogu.com.cn/training/4295
模板题 https://www.luogu.com.cn/problem/P4148
todo https://codeforces.com/problemset/problem/44/G
*/

type kdNode struct {
	lr          [2]*kdNode
	p, mi, mx   [2]int // 0 为 x，1 为 y
	sz, val, sm int
}

func (kdNode) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (kdNode) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (o *kdNode) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *kdNode) sum() int {
	if o != nil {
		return o.sm
	}
	return 0
}

func (o *kdNode) maintain() {
	o.sz = o.lr[0].size() + o.lr[1].size() + 1
	o.sm = o.lr[0].sum() + o.lr[1].sum() + o.val
	for i := 0; i < 2; i++ {
		o.mi[i] = o.p[i]
		o.mx[i] = o.p[i]
		for _, ch := range o.lr {
			if ch != nil {
				o.mi[i] = o.min(o.mi[i], ch.mi[i])
				o.mx[i] = o.max(o.mx[i], ch.mx[i])
			}
		}
	}
}

func (o *kdNode) nodes() []*kdNode {
	nodes := make([]*kdNode, 0, o.size())
	var f func(*kdNode)
	f = func(o *kdNode) {
		if o != nil {
			nodes = append(nodes, o)
			f(o.lr[0])
			f(o.lr[1])
		}
	}
	f(o)
	rand.Shuffle(len(nodes), func(i, j int) { nodes[i], nodes[j] = nodes[j], nodes[i] })
	return nodes
}

func divideKDT(a []*kdNode, k, dim int) {
	for l, r := 0, len(a)-1; l < r; {
		v := a[l].p[dim]
		i, j := l, r+1
		for {
			for i++; i < r && a[i].p[dim] < v; i++ {
			}
			for j--; j > l && a[j].p[dim] > v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[l], a[j] = a[j], a[l]
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
}

// 另一种实现是选择的维度要满足其内部点的分布的差异度最大，见 https://oi-wiki.org/ds/kdt/
func buildKDT(nodes []*kdNode, dim int) *kdNode {
	if len(nodes) == 0 {
		return nil
	}
	m := len(nodes) / 2
	divideKDT(nodes, m, dim)
	o := nodes[m]
	o.lr[0] = buildKDT(nodes[:m], dim^1)
	o.lr[1] = buildKDT(nodes[m+1:], dim^1)
	o.maintain()
	return o
}

func (o *kdNode) rebuild(dim int) *kdNode { return buildKDT(o.nodes(), dim) }

func (o *kdNode) put(p [2]int, val, dim int) *kdNode {
	if o == nil {
		o = &kdNode{p: p, val: val}
		o.maintain()
		return o
	}
	if p[dim] < o.p[dim] {
		o.lr[0] = o.lr[0].put(p, val, dim^1)
	} else {
		o.lr[1] = o.lr[1].put(p, val, dim^1)
	}
	o.maintain()
	if sz := o.size() * 3; o.lr[0].size()*4 > sz || o.lr[1].size()*4 > sz { // alpha=3/4
		return o.rebuild(dim)
	}
	return o
}

// 矩形 X-Y 在矩形 x-y 内
func inRect(x1, y1, x2, y2, X1, Y1, X2, Y2 int) bool {
	return x1 <= X1 && X2 <= x2 && y1 <= Y1 && Y2 <= y2
}

// 矩形 X-Y 在矩形 x-y 外
func outRect(x1, y1, x2, y2, X1, Y1, X2, Y2 int) bool {
	return X2 < x1 || X1 > x2 || Y2 < y1 || Y1 > y2
}

func (o *kdNode) query(x1, y1, x2, y2 int) (res int) {
	if o == nil || outRect(x1, y1, x2, y2, o.mi[0], o.mi[1], o.mx[0], o.mx[1]) {
		return
	}
	if inRect(x1, y1, x2, y2, o.mi[0], o.mi[1], o.mx[0], o.mx[1]) {
		return o.sm
	}
	if inRect(x1, y1, x2, y2, o.p[0], o.p[1], o.p[0], o.p[1]) { // 根在询问矩形内
		res = o.val
	}
	res += o.lr[0].query(x1, y1, x2, y2) + o.lr[1].query(x1, y1, x2, y2)
	return
}

type kdTree struct {
	root *kdNode
}

func newKdTree() *kdTree {
	rand.Seed(time.Now().UnixNano())
	return &kdTree{}
}

func (t *kdTree) put(p [2]int, val int) { t.root = t.root.put(p, val, 0) }

func (t *kdTree) query(x1, y1, x2, y2 int) int { return t.root.query(x1, y1, x2, y2) }
