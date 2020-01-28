package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF722C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	var n int
	Fscan(in, &n)
	arr := make([]int, n)
	for i := range arr {
		Fscan(in, &arr[i])
	}
	pos := make([]int, n)
	for i := range pos {
		Fscan(in, &pos[i])
	}

	ans := make([]int64, n)
	vis := make([]bool, n+2)
	sum := make([]int64, n+1)
	initFa(n + 1)
	max := int64(0)
	for i := n - 1; i >= 0; i-- {
		ans[i] = max
		p := pos[i]
		vis[p] = true
		sum[p] += int64(arr[p-1])
		if vis[p-1] {
			f := find(p - 1)
			sum[p] += sum[f]
			fa[f] = p
		}
		if vis[p+1] {
			f := find(p + 1)
			sum[p] += sum[f]
			fa[f] = p
		}
		if sum[p] > max {
			max = sum[p]
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() {
//	CF722C(os.Stdin, os.Stdout)
//}
