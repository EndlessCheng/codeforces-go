package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

type odtBlock896C struct {
	l, r int
	val  int64
}

type odt896C []odtBlock896C

func (t *odt896C) split(mid int) int {
	ot := *t
	for i, b := range ot {
		if b.l == mid+1 {
			return i
		}
		if b.l <= mid && mid < b.r {
			*t = append(ot[:i+1], append(odt896C{{mid + 1, b.r, b.val}}, ot[i+1:]...)...)
			ot[i].r = mid
			return i + 1
		}
	}
	return len(ot)
}

func (t *odt896C) prepare(l, r int) (begin, end int) {
	begin = t.split(l - 1)
	end = t.split(r)
	return
}

func (t *odt896C) merge(begin, end, r int, val int64) {
	ot := *t
	ot[begin].r = r
	ot[begin].val = val
	if begin+1 < end {
		*t = append(ot[:begin+1], ot[end:]...)
	}
}

func (t odt896C) add(begin, end int, val int64) {
	for i := begin; i < end; i++ {
		t[i].val += val
	}
}

func (t odt896C) kth(begin, end, k int) int64 {
	blocks := make(odt896C, end-begin)
	copy(blocks, t[begin:end])
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

func (odt896C) quickPow(x int64, n int, mod int64) int64 {
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

func (t odt896C) powSum(begin, end int, n int, mod int64) (res int64) {
	for _, b := range t[begin:end] {
		res += int64(b.r-b.l+1) * t.quickPow(b.val, n, mod)
	}
	return res % mod
}

// github.com/EndlessCheng/codeforces-go
func Sol896C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m, vMax int
	var seed int64
	Fscan(in, &n, &m, &seed, &vMax)
	rand := func(_n int) int {
		const mod int64 = 1e9 + 7
		ret := seed
		seed = (seed*7 + 13) % mod
		return int(ret) % _n
	}

	t := make(odt896C, n)
	for i := range t {
		t[i] = odtBlock896C{i, i, int64(rand(vMax) + 1)}
	}
	for ; m > 0; m-- {
		op := rand(4) + 1
		l, r := rand(n), rand(n)
		if l > r {
			l, r = r, l
		}
		var x int
		if op == 3 {
			x = rand(r-l+1) + 1
		} else {
			x = rand(vMax) + 1
		}
		begin, end := t.prepare(l, r)
		switch op {
		case 1:
			t.add(begin, end, int64(x))
		case 2:
			t.merge(begin, end, r, int64(x))
		case 3:
			Fprintln(out, t.kth(begin, end, x))
		default:
			y := int64(rand(vMax) + 1)
			Fprintln(out, t.powSum(begin, end, x, y))
		}
	}
}

//func main() {
//	Sol896C(os.Stdin, os.Stdout)
//}