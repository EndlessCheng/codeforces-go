package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF547B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	left := make([]int, n)
	st := []int{-1}
	for i := range a {
		Fscan(in, &a[i])
		for len(st) > 1 && a[st[len(st)-1]] >= a[i] {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	right := make([]int, n)
	st = []int{n}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 1 && a[st[len(st)-1]] >= a[i] {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}

	ans := make([]int, n+1)
	for i, v := range a {
		size := right[i] - left[i] - 1
		ans[size] = max(ans[size], v)
	}
	for i := n - 1; i > 0; i-- {
		ans[i] = max(ans[i], ans[i+1])
	}
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF547B(os.Stdin, os.Stdout) }
