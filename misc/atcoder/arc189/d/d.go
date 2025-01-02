package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	sum := make([]int, n+1)
	leftGE := make([]int, n)
	right := make([]int, n)
	st := []int{-1}
	for i := range a {
		Fscan(in, &a[i])
		sum[i+1] = sum[i] + a[i]
		for len(st) > 1 && a[st[len(st)-1]] < a[i] {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		leftGE[i] = st[len(st)-1]
		st = append(st, i)
	}
	for _, i := range st[1:] {
		right[i] = n
	}

	rightGE := make([]int, n)
	left := make([]int, n)
	st = []int{n}
	for i, v := range slices.Backward(a) {
		for len(st) > 1 && a[st[len(st)-1]] < v {
			left[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		rightGE[i] = st[len(st)-1]
		st = append(st, i)
	}
	for _, i := range st[1:] {
		left[i] = -1
	}

	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return a[id[i]] > a[id[j]] })

	ans := make([]int, n)
	for _, i := range id {
		if rightGE[i]-leftGE[i] == 2 {
			ans[i] = a[i]
			continue
		}
		l, r := left[i], right[i]
		s := sum[r] - sum[l+1]
		if l >= 0 && s > a[l] {
			ans[i] = ans[l]
		} else if r < n && s > a[r] {
			ans[i] = ans[r]
		} else {
			ans[i] = s
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
