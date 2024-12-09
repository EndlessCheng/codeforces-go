package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type fenwick []int

func (f fenwick) add(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func p2163(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, x, x1, y1, x2, y2 int
	Fscan(in, &n, &m)
	xMap := map[int][]int{}
	ys := make([]int, n)
	for i := range ys {
		Fscan(in, &x, &ys[i])
		xMap[x] = append(xMap[x], ys[i])
	}

	xs := make([]int, 0, len(xMap))
	for k := range xMap {
		xs = append(xs, k)
	}
	slices.Sort(xs)
	slices.Sort(ys)
	ys = slices.Compact(ys)

	// 离线询问
	type data struct{ qid, sign, y1, y2 int }
	qs := make([][]data, len(xs))
	for i := 0; i < m; i++ {
		Fscan(in, &x1, &y1, &x2, &y2)
		x1 = sort.SearchInts(xs, x1) // 离散化，下标从 0 开始
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

	// 回答询问
	ans := make([]int, m)
	t := make(fenwick, len(ys)+1)
	for i, x := range xs { // 从小到大枚举 x
		// 把横坐标为 x 的所有点都加到树状数组中
		for _, y := range xMap[x] {
			t.add(sort.SearchInts(ys, y) + 1) // 离散化，并且下标从 1 开始
		}
		for _, q := range qs[i] {
			// 查询横坐标 <= x（已满足）且纵坐标在 [y1,y2] 中的点的个数
			ans[q.qid] += q.sign * t.query(q.y1+1, q.y2+1) // 下标从 1 开始
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { p2163(bufio.NewReader(os.Stdin), os.Stdout) }
