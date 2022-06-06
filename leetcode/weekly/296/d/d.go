package main

// https://space.bilibili.com/206214/dynamic
type node struct {
	ch  [2]*node
	sz  int
	key byte
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

// 构建一颗中序遍历为 [l,r] 的 splay 树
// 比如，给你一个序列和一些修改操作，每次取出一段子区间，cut 掉然后 append 到末尾，输出完成所有操作后的最终序列：
//     我们可以 buildSplay(1,n)，每次操作调用两次 split 来 cut 区间，得到三颗子树 a b c
//     append 之后应该是 a c b，那么我们可以 a.merge(c.merge(b)) 来完成这一操作
//     注意 merge 后可能就不满足搜索树的性质了，但是没有关系，中序遍历的结果仍然是正确的，我们只要保证这一点成立，就能正确得到完成所有操作后的最终序列
func buildSplay(s string) *node {
	if s == "" {
		return nil
	}
	m := len(s) / 2
	o := &node{key: s[m]}
	o.ch[0] = buildSplay(s[:m])
	o.ch[1] = buildSplay(s[m+1:])
	o.maintain()
	return o
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

type TextEditor struct {
	root *node
	cur  int
}

func Constructor() TextEditor { return TextEditor{} }

func (t *TextEditor) AddText(text string) {
	if t.cur == 0 {
		t.root = buildSplay(text).merge(t.root)
	} else {
		lo, ro := t.root.split(t.cur)
		t.root = lo.merge(buildSplay(text)).merge(ro)
	}
	t.cur += len(text)
}

func (t *TextEditor) DeleteText(k int) int {
	if t.cur == 0 {
		return 0
	}
	if t.cur <= k { // 左边全部删除
		_, t.root = t.root.split(t.cur)
		ans := t.cur
		t.cur = 0
		return ans
	} else {
		lo, ro := t.root.split(t.cur)
		t.cur -= k
		lo, _ = lo.split(t.cur) // 删除中间 k 个
		t.root = lo.merge(ro)
		return k
	}
}

func (t *TextEditor) text() string {
	if t.cur == 0 {
		return ""
	}
	k := max(t.cur-10, 0)
	t.root = t.root.splay(k + 1)
	ans := make([]byte, 1, t.cur-k)
	ans[0] = t.root.key
	var inorder func(*node) bool
	inorder = func(o *node) bool {
		if o == nil {
			return false
		}
		if inorder(o.ch[0]) || len(ans) == cap(ans) {
			return true
		}
		ans = append(ans, o.key)
		return inorder(o.ch[1])
	}
	inorder(t.root.ch[1])
	return string(ans)
}

func (t *TextEditor) CursorLeft(k int) string {
	t.cur = max(t.cur-k, 0)
	return t.text()
}

func (t *TextEditor) CursorRight(k int) string {
	t.cur = min(t.cur+k, t.root.size())
	return t.text()
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
