package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf91B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := make([]any, n)
	ans[n-1] = -1
	st := []int{n - 1}
	for i := n - 2; i >= 0; i-- {
		v := a[i]
		j := sort.Search(len(st), func(j int) bool { return a[st[j]] < v })
		if j < len(st) {
			ans[i] = st[j] - i - 1
		} else {
			ans[i] = -1
		}
		if v < a[st[len(st)-1]] {
			st = append(st, i)
		}
	}
	Fprintln(out, ans...)
}

//func main() { cf91B(os.Stdin, os.Stdout) }
