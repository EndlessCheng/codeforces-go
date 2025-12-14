package main

import (
	"bufio"
	. "fmt"
	"io"
	"maps"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick755 []int

func (t fenwick755) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick755) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return res
}

func (t fenwick755) query(l, r int) int {
	if r < l {
		return 0
	}
	return t.pre(r) - t.pre(l-1)
}

func p3755(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	type pair struct{ v, power int }
	xMap := map[int][]pair{}
	ys := make([]int, n)
	for i := range ys {
		var x, p int
		Fscan(in, &x, &ys[i], &p)
		xMap[x] = append(xMap[x], pair{ys[i], p})
	}

	xs := slices.Sorted(maps.Keys(xMap))
	slices.Sort(ys)
	ys = slices.Compact(ys)

	type data struct{ qid, sign, y1, y2 int }
	qs := make([][]data, len(xs))
	for i := range m {
		var x1, y1, x2, y2 int
		Fscan(in, &x1, &y1, &x2, &y2)
		x1 = sort.SearchInts(xs, x1)
		x2 = sort.SearchInts(xs, x2+1) - 1
		if x1 > x2 {
			continue
		}
		y1 = sort.SearchInts(ys, y1)
		y2 = sort.SearchInts(ys, y2+1) - 1
		if y1 > y2 {
			continue
		}
		if x1 > 0 {
			qs[x1-1] = append(qs[x1-1], data{i, -1, y1, y2})
		}
		qs[x2] = append(qs[x2], data{i, 1, y1, y2})
	}

	ans := make([]int, m)
	t := make(fenwick755, len(ys)+1)
	for i, x := range xs {
		// 把横坐标为 x 的所有点都加到树状数组中
		for _, p := range xMap[x] {
			t.update(sort.SearchInts(ys, p.v)+1, p.power)
		}
		for _, q := range qs[i] {
			// 查询横坐标 <= x（已满足）且纵坐标在 [y1,y2] 中的点的个数
			ans[q.qid] += q.sign * t.query(q.y1+1, q.y2+1)
		}
	}

	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { p3755(bufio.NewReader(os.Stdin), os.Stdout) }
