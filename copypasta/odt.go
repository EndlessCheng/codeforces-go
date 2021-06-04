package copypasta

import "sort"

// 说明见 odt_bst.go

// 这里用 slice 实现 O(nlogn)
// 对应题目 https://codeforces.com/problemset/problem/896/C
// 我的题解 https://www.luogu.com.cn/blog/endlesscheng/solution-cf896c
type odtBlock struct {
	l, r int
	val  int64
}

type odt []odtBlock

func newODT(arr []int64) odt {
	n := len(arr)
	t := make(odt, n)
	for i := range t {
		t[i] = odtBlock{i, i, arr[i]}
	}
	return t
}

// [l, r] => [l, mid-1] [mid, r]
// return index of [mid, r]
// return len(t) if not found
func (t *odt) split(mid int) int {
	ot := *t
	for i, b := range ot {
		if b.l == mid {
			return i
		}
		if b.l < mid && mid <= b.r { // b.l <= mid-1
			*t = append(ot[:i+1], append(odt{{mid, b.r, b.val}}, ot[i+1:]...)...)
			ot[i].r = mid - 1
			return i + 1
		}
	}
	return len(ot)
}

func (t *odt) prepare(l, r int) (begin, end int) {
	begin = t.split(l)
	end = t.split(r + 1)
	return
}

// 以下方法传入的 begin, end 来自事先计算的 t.prepare

func (t *odt) merge(begin, end, r int, val int64) {
	ot := *t
	ot[begin].r = r
	ot[begin].val = val
	if begin+1 < end {
		*t = append(ot[:begin+1], ot[end:]...)
	}
}

func (t odt) add(begin, end int, val int64) {
	for i := begin; i < end; i++ {
		t[i].val += val
	}
}

func (t odt) kth(begin, end, k int) int64 {
	blocks := append(odt(nil), t[begin:end]...)
	sort.Slice(blocks, func(i, j int) bool { return blocks[i].val < blocks[j].val })
	k--
	for _, b := range blocks {
		if cnt := b.r - b.l + 1; k >= cnt {
			k -= cnt
		} else {
			return b.val
		}
	}
	panic(k)
}

func (odt) pow(x int64, n int, mod int64) int64 {
	x %= mod
	res := int64(1) % mod
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func (t odt) powSum(begin, end int, n int, mod int64) (res int64) {
	for _, b := range t[begin:end] {
		// 总和能溢出的话这里要额外取模
		res += int64(b.r-b.l+1) * t.pow(b.val, n, mod)
	}
	return res % mod
}
