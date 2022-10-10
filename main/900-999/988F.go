package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF988F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var n, m int
	Fscan(in, &n, &n, &m)
	seg := make([]struct{ l, r int }, n)
	for i := range seg {
		Fscan(in, &seg[i].l, &seg[i].r)
	}
	sort.Slice(seg, func(i, j int) bool { return seg[i].l < seg[j].r })
	type pair struct{ x, w int }
	a := make([]pair, m, m+1)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].w)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
	if a[0].x > seg[0].l {
		Fprint(out, -1)
		return
	}

	a = append(a, pair{seg[n-1].r, 0}) // 方便计算答案（也可以用 dst，但其实完全不需要这个值）
	f := make([]int, m+1)
	j := -1
	for i, p := range a {
		x := p.x
		for j+1 < n && seg[j+1].l < x { // 寻找 x 左侧最近区间左端点
			j++
		}
		if j < 0 {
			continue
		}
		x = min(x, seg[j].r) // 如果 x 在区间外（x 与左端点重合也算），那么移动到左侧区间的右端点时就应该丢伞了
		if a[i-1].x >= x {   // 如果上面没有 continue 那么这里必然 i>0
			f[i] = f[i-1]
			continue
		}
		f[i] = 1e9
		for j, q := range a[:i] {
			f[i] = min(f[i], f[j]+(x-q.x)*q.w)
		}
	}
	Fprint(out, f[m])
}

//func main() { CF988F(os.Stdin, os.Stdout) }
