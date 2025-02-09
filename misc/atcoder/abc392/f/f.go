package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
type node struct {
	ch  [2]*node
	sz  int
	key int
}

// 设置如下返回值是为了方便使用 node 中的 ch 数组
func (o *node) cmpKth(k int) int {
	d := k - o.ch[0].size() - 1
	switch {
	case d < 0:
		return 0 // 左儿子
	case d > 0:
		return 1 // 右儿子
	default:
		return -1
	}
}

func (o *node) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *node) maintain() {
	o.sz = 1 + o.ch[0].size() + o.ch[1].size()
}

// 旋转，并维护子树大小
// d=0：左旋，返回 o 的右儿子
// d=1：右旋，返回 o 的左儿子
func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	o.maintain()
	x.maintain()
	return x
}

// 将子树 o（中序遍历）的第 k 个节点伸展到 o，并返回该节点
// 1 <= k <= o.size()
func (o *node) splay(k int) (kth *node) {
	d := o.cmpKth(k)
	if d < 0 {
		return o
	}
	k -= d * (o.ch[0].size() + 1)
	c := o.ch[d]
	if d2 := c.cmpKth(k); d2 >= 0 {
		c.ch[d2] = c.ch[d2].splay(k - d2*(c.ch[0].size()+1))
		if d2 == d {
			o = o.rotate(d ^ 1)
		} else {
			o.ch[d] = c.rotate(d)
		}
	}
	return o.rotate(d ^ 1)
}

// 分裂子树 o，把 o（中序遍历）的前 k 个节点放在 lo 子树，其余放在 ro 子树
// 返回的 lo 节点为 o（中序遍历）的第 k 个节点
// 1 <= k <= o.size()
// 特别地，k = o.size() 时 ro 为 nil
func (o *node) split(k int) (lo, ro *node) {
	lo = o.splay(k)
	ro = lo.ch[1]
	lo.ch[1] = nil
	lo.maintain()
	return
}

// 把子树 ro 合并进子树 o，返回合并前 o（中序遍历）的最后一个节点
// 相当于把 ro 的中序遍历 append 到 o 的中序遍历之后
// ro 可以为 nil，但 o 不能为 nil
func (o *node) merge(ro *node) *node {
	// 把最大节点伸展上来，这样会空出一个右儿子用来合并 ro
	o = o.splay(o.size())
	o.ch[1] = ro
	o.maintain()
	return o
}

var root *node

func add(i, v int) {
	o := &node{key: v, sz: 1}
	if i == 0 {
		root = o.merge(root)
	} else {
		lo, ro := root.split(i)
		root = lo.merge(o).merge(ro)
	}
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, p int
	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &p)
		add(p-1, i)
	}
	var inorder func(*node)
	inorder = func(o *node) {
		if o == nil {
			return
		}
		inorder(o.ch[0])
		Fprint(out, o.key, " ")
		inorder(o.ch[1])
	}
	inorder(root)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
