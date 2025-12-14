package main

import "math/bits"

// https://space.bilibili.com/206214
type data struct {
	lc  byte // 区间左端点字母
	rc  byte // 区间右端点字母
	del int  // 区间删除次数
}

// 线段树有两个下标，一个是线段树节点的下标，另一个是线段树维护的区间的下标
// 节点的下标：从 1 开始，如果你想改成从 0 开始，需要把左右儿子下标分别改成 node*2+1 和 node*2+2
// 区间的下标：从 0 开始
type seg []data

// 合并左右子树的信息
func merge(l, r data) data {
	ans := l.del + r.del
	if l.rc == r.lc {
		ans++ // 多删一个字母
	}
	return data{l.lc, r.rc, ans}
}

func newSegmentTree(a string) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

// 合并左右儿子的信息到当前节点
func (t seg) maintain(node int) {
	t[node] = merge(t[node*2], t[node*2+1])
}

// 用 a 初始化线段树
func (t seg) build(a string, node, l, r int) {
	if l == r { // 叶子
		t[node].lc = a[l] - 'A'
		t[node].rc = t[node].lc
		return
	}
	m := (l + r) / 2
	t.build(a, node*2, l, m)     // 初始化左子树
	t.build(a, node*2+1, m+1, r) // 初始化右子树
	t.maintain(node)
}

func (t seg) flip(node, l, r, i int) {
	if l == r { // 叶子（到达目标）
		t[node].lc ^= 1
		t[node].rc ^= 1
		return
	}
	m := (l + r) / 2
	if i <= m { // i 在左子树
		t.flip(node*2, l, m, i)
	} else { // i 在右子树
		t.flip(node*2+1, m+1, r, i)
	}
	t.maintain(node)
}

func (t seg) query(node, l, r, ql, qr int) data {
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		return t[node]
	}
	m := (l + r) / 2
	if qr <= m { // [ql, qr] 在左子树
		return t.query(node*2, l, m, ql, qr)
	}
	if ql > m { // [ql, qr] 在右子树
		return t.query(node*2+1, m+1, r, ql, qr)
	}
	return merge(t.query(node*2, l, m, ql, qr), t.query(node*2+1, m+1, r, ql, qr))
}

func minDeletions1(s string, queries [][]int) (ans []int) {
	n := len(s)
	t := newSegmentTree(s)
	for _, q := range queries {
		if q[0] == 1 {
			t.flip(1, 0, n-1, q[1])
		} else {
			ans = append(ans, t.query(1, 0, n-1, q[1], q[2]).del)
		}
	}
	return
}

//

type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) update(i int, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// 求区间和 a[l] + ... + a[r]
// 1 <= l <= r <= n
// 时间复杂度 O(log n)
func (f fenwick) query(l, r int) int {
	if r < l {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

func minDeletions(s string, queries [][]int) (ans []int) {
	n := len(s)
	t := newFenwickTree(n - 1)
	for i := 1; i < n; i++ {
		if s[i-1] == s[i] { // 删除 i
			t.update(i, 1)
		}
	}

	bs := []byte(s)
	for _, q := range queries {
		if q[0] == 2 {
			ans = append(ans, t.query(q[1]+1, q[2]))
			continue
		}

		i := q[1]

		if i > 0 {
			val := 1
			if bs[i-1] == bs[i] {
				val = -1
			}
			t.update(i, val)
		}

		if i < n-1 {
			val := 1
			if bs[i] == bs[i+1] {
				val = -1
			}
			t.update(i+1, val)
		}

		bs[i] ^= 'A' ^ 'B' // A 变成 B，B 变成 A
	}
	return
}
