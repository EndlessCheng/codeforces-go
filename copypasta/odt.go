package copypasta

import "sort"

// "Old Driver Tree"
// 一种可以动态合并与分裂的分块结构，在随机数据下有高效性能
// https://oi-wiki.org/ds/odt/

// 这里用 slice 实现
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

func (t *odt) split(mid int) {
	ot := *t
	for i, b := range ot {
		if b.l <= mid && mid < b.r {
			*t = append(ot[:i+1], append(odt{{mid + 1, b.r, b.val}}, ot[i+1:]...)...)
			ot[i].r = mid
			break
		}
	}
}

func (t *odt) prepare(l, r int) {
	t.split(l - 1)
	t.split(r)
}

// All op below should first call `prepare(l, r)`

func (t *odt) merge(l, r int, val int64) {
	ot := *t
	for i, b := range ot {
		if b.l == l {
			ot[i] = odtBlock{b.l, r, val}
			j := i + 1
			for ; j < len(ot) && ot[j].l <= r; j++ {
			}
			if j > i+1 {
				*t = append(ot[:i+1], ot[j:]...)
			}
			break
		}
	}
}

func (t odt) add(l, r int, val int64) {
	for i, b := range t {
		if l <= b.l && b.r <= r {
			t[i].val += val
		}
	}
}

func (t odt) kth(l, r, k int) int64 {
	blocks := []odtBlock{}
	for _, b := range t {
		if l <= b.l && b.r <= r {
			blocks = append(blocks, b)
		}
	}
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

func (odt) quickPow(x int64, n int, mod int64) int64 {
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

func (t odt) powSum(l, r int, n int, mod int64) (res int64) {
	for _, b := range t {
		if l <= b.l && b.r <= r {
			// 总和能溢出的话这里要额外取模
			res += int64(b.r-b.l+1) * t.quickPow(b.val, n, mod)
		}
	}
	return res % mod
}
