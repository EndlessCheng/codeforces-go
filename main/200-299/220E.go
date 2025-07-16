package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type fenwick20 []int

func (f fenwick20) update(i, v int) {
	for ; i < len(f); i += i & -i {
		f[i] += v
	}
}

func (f fenwick20) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func cf220E(in io.Reader, out io.Writer) {
	var n, k, l, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	b := slices.Clone(a)
	slices.Sort(b)
	b = slices.Compact(b)
	m := len(b)

	tr := make(fenwick20, m+1)
	for i := n - 1; i >= 0; i-- {
		a[i] = sort.SearchInts(b, a[i]) + 1
		k -= tr.pre(a[i] - 1)
		tr.update(a[i], 1)
	}

	tl := make(fenwick20, m+1)
	for r := 1; r < n; r++ {
		tr.update(a[r-1], -1)
		k += l - tl.pre(a[r-1]) + tr.pre(a[r-1]-1)
		for l < r {
			inv := l - tl.pre(a[l]) + tr.pre(a[l]-1)
			if k < inv {
				break
			}
			k -= inv
			tl.update(a[l], 1)
			l++
		}
		ans += l
	}
	Fprint(out, ans)
}

//func main() { cf220E(bufio.NewReader(os.Stdin), os.Stdout) }
