package copypasta

// "Old Driver Tree" "珂朵莉树"
// 一种可以动态合并与分裂的分块结构，在随机数据下有高效性能 O(nloglogn)
// 一般情况可以用于区间更新单点查询
// https://oi-wiki.org/ds/odt/
// todo 随机数据下的复杂度证明 https://zhuanlan.zhihu.com/p/102786071

// 模板题 LC715 https://leetcode-cn.com/problems/range-module/
// LC699 https://leetcode.cn/problems/falling-squares/
// https://atcoder.jp/contests/abl/tasks/abl_e 代码【比较全面】https://atcoder.jp/contests/abl/submissions/36147029
// https://codeforces.com/problemset/problem/292/E 代码 https://codeforces.com/contest/292/submission/173666674
// https://codeforces.com/problemset/problem/558/E 代码 https://codeforces.com/problemset/submission/558/117163317
// https://codeforces.com/problemset/problem/915/E 代码【比较全面】https://codeforces.com/problemset/submission/915/117158161
// https://codeforces.com/problemset/problem/817/F（数据水）代码 https://codeforces.com/contest/817/submission/118365591
// todo https://www.luogu.com.cn/problem/P5350
//      https://www.luogu.com.cn/problem/P5586

// 使用时，为简化判断，可在初始时插入一段 [1,n] 区间（或 [0,2e9] 等）
// 直接复制上面的代码
type odtNode struct {
	tpNode
	l, r int
}

func (t *treap) put1(l, r int, val tpValueType) {}
func (t *treap) floor(int) (_ *odtNode)         { return }
func (t *treap) next(int) (_ *odtNode)          { return }

func (t *treap) split(mid int) {
	if o := t.floor(mid); o.l < mid && mid <= o.r {
		r, val := o.r, o.val
		o.r = mid - 1
		t.put1(mid, r, val)
	}
}

func (t *treap) prepare(l, r int) {
	t.split(l)
	t.split(r + 1)
}

func (t *treap) setRange(l, r int, val tpValueType) {
	t.prepare(l, r)
	// 保留 l，后面直接修改，从而代替删除+插入操作
	for o := t.next(l); o != nil && o.l <= r; o = t.next(o.l) {
		t.delete(tpKeyType(o.l))
	}
	o := t.floor(l)
	o.r = r
	o.val = val
}

// https://codeforces.com/problemset/problem/558/E
func (t *treap) sortBytes(l, r int, inc bool) {
	t.prepare(l, r)
	// 整段删除重建
	cnt := [26]int{}
	for o := t.floor(l); o != nil && o.l <= r; o = t.next(o.l) {
		cnt[o.val] += o.r - o.l + 1
		t.delete(tpKeyType(o.l))
	}
	if inc {
		for i, c := range cnt {
			if c > 0 {
				t.put1(l, l+c-1, tpValueType(i))
				l += c
			}
		}
	} else {
		for i, c := range cnt {
			if c > 0 {
				t.put1(r-c+1, r, tpValueType(i))
				r -= c
			}
		}
	}
}
