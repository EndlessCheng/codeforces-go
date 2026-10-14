package main

import (
	. "fmt"
	"io"
	"maps"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick []int

func (t fenwick) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func (t fenwick) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

func p2448(in io.Reader, out io.Writer) {
	var k, p, q, ans int
	Fscan(in, &k)
	s := map[int]int{}
	for range k {
		Fscan(in, &p, &q)
		if s[p] == 0 {
			s[p] = p
		}
		if s[q] == 0 {
			s[q] = q
		}
		s[p], s[q] = s[q], s[p]
	}

	idx := slices.Sorted(maps.Keys(s))
	m := len(idx)
	t := make(fenwick, m+1)
	for p, i := range idx {
		if p > 0 {
			cnt := i - idx[p-1] - 1
			ans += t.query(p+1, m) * cnt
			t.update(p, cnt)
		}
		v := sort.SearchInts(idx, s[i]) + 1
		ans += t.query(v, m)
		t.update(v, 1)
	}
	Fprint(out, ans)
}

//func main() { p2448(bufio.NewReader(os.Stdin), os.Stdout) }
