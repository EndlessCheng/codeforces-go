package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func mergeCount1966(a []int) int {
	n := len(a)
	if n <= 1 {
		return 0
	}
	left := slices.Clone(a[:n/2])
	right := slices.Clone(a[n/2:])
	cnt := mergeCount1966(left) + mergeCount1966(right)
	l, r := 0, 0
	for i := range a {
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			a[i] = left[l]
			l++
		} else {
			cnt += n/2 - l
			a[i] = right[r]
			r++
		}
	}
	return cnt
}

func p1966(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	f := func() []int {
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := slices.Clone(a)
		slices.Sort(b)
		for i, v := range a {
			a[i] = sort.SearchInts(b, v)
		}
		return a
	}
	// 注意不能排序，而是要做置换
	// 元素位置不能变，只是值变了
	a := f()
	p := make([]int, n)
	for i, x := range a {
		p[x] = i
	}
	b := f()
	for i, v := range b {
		b[i] = p[v]
	}
	Fprint(out, mergeCount1966(b)%(1e8-3))
}

//func main() { p1966(bufio.NewReader(os.Stdin), os.Stdout) }
