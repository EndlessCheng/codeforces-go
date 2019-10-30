package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

type odtBlock struct {
	l, r int
	val  int64
}

type odt []odtBlock

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

func (t odt) add(l, r int, val int64) {
	for i, b := range t {
		if b.l > r {
			break
		}
		if l <= b.l && b.r <= r {
			t[i].val += val
		}
	}
}

func (t *odt) set(l, r int, val int64) {
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
			res += int64(b.r-b.l+1) * t.quickPow(b.val, n, mod)
		}
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

	t := make(odt, n)
	for i := range t {
		t[i] = odtBlock{i, i, int64(rand(vMax) + 1)}
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
		t.prepare(l, r)
		switch op {
		case 1:
			t.add(l, r, int64(x))
		case 2:
			t.set(l, r, int64(x))
		case 3:
			Fprintln(out, t.kth(l, r, x))
		default:
			y := int64(rand(vMax) + 1)
			Fprintln(out, t.powSum(l, r, x, y))
		}
	}
}

//func main() {
//	Sol896C(os.Stdin, os.Stdout)
//}
