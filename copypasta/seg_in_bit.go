package copypasta

import "math"

// 树套树：树状数组套动态开点权值线段树
// https://oi-wiki.org/ds/seg-in-bit/
// 三维偏序 https://www.luogu.com.cn/problem/P3810 https://www.luogu.com.cn/record/136178821
// https://www.luogu.com.cn/problem/P3380 https://www.luogu.com.cn/record/136286395
// https://www.luogu.com.cn/problem/P2617 Dynamic Rankings
// 树状数组在这里就是纯纯工具人，只用来拆分区间
type fenwickWithSeg []*stNode // 见 segment_tree.go

func newFenwickTreeWithSeg(n, mx int) fenwickWithSeg {
	t := make(fenwickWithSeg, n+1)
	for i := range t {
		t[i] = newStRoot(0, mx) // 注意下界
	}
	return t
}

// 二维单点更新：位置 (i,j) 用 val 更新
func (f fenwickWithSeg) update(i, j, val int) {
	for ; i < len(f); i += i & -i {
		f[i].update(j, val)
	}
}

// 二维前缀和：累加所有 x <= i 且 y <= j 的值
func (f fenwickWithSeg) pre(i, j int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i].query(0, j) // 注意下界
	}
	return
}

// 返回 [l,r] 内 v 的排名，即小于 v 的元素个数 + 1
func (f fenwickWithSeg) rank(l, r, v int) int {
	return f.pre(r, v-1) - f.pre(l-1, v-1) + 1
}

// 返回 [l,r] 内的第 k 个数（设为 v），即有 k-1 个数小于 v
// 需要保证 1 <= k <= r-l+1
func (f fenwickWithSeg) kth(l, r, k int) int {
	var ar, al []*stNode // 如果禁止 GC，需要把这行放在外面，然后用下面注释的两行代码（复用内存）
	//ar = ar[:0]
	//al = al[:0]
	for ; r > 0; r &= r - 1 {
		ar = append(ar, f[r])
	}
	for l--; l > 0; l &= l - 1 {
		al = append(al, f[l])
	}
	left, right := 0, f[0].r
	for left < right {
		s := 0
		for _, o := range ar {
			s += o.lo.val
		}
		for _, o := range al {
			s -= o.lo.val
		}
		mid := (left + right) >> 1
		if s >= k {
			for i, o := range ar {
				ar[i] = o.lo
			}
			for i, o := range al {
				al[i] = o.lo
			}
			right = mid
		} else {
			k -= s
			for i, o := range ar {
				ar[i] = o.ro
			}
			for i, o := range al {
				al[i] = o.ro
			}
			left = mid + 1
		}
	}
	return left
}

// 返回 [l,r] 内严格小于 v 的最大的数
// 如果不存在，返回 math.MinInt
func (f fenwickWithSeg) prev(l, r, v int) int {
	rk := f.rank(l, r, v)
	if rk == 1 { // 没有比 v 小的数
		return math.MinInt
	}
	return f.kth(l, r, rk-1)
}

// 返回 [l,r] 内严格大于 v 的最小的数
// 如果不存在，返回 math.MaxInt
func (f fenwickWithSeg) next(l, r, v int) int {
	rk := f.rank(l, r, v+1)
	if rk > r-l+1 { // [l,r] 内的所有数都 <= v
		return math.MaxInt
	}
	return f.kth(l, r, rk)
}
