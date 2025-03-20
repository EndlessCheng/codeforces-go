package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func p8776(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	pre := []int{}
	type pair struct{ i, v int }
	rollback := make([]pair, n-k)
	for i, v := range a[:n-k] {
		j := sort.SearchInts(pre, v+1)
		if j < len(pre) {
			rollback[i] = pair{j, pre[j]}
			pre[j] = v
		} else {
			rollback[i] = pair{j, -1}
			pre = append(pre, v)
		}
	}
	ans := len(pre) + k

	suf := []int{}
	for i := n - 1; i >= k; i-- {
		p := rollback[i-k]
		if p.v < 0 {
			pre = pre[:len(pre)-1]
		} else {
			pre[p.i] = p.v
		}
		v := -a[i]
		j := sort.SearchInts(suf, v+1)
		if j < len(suf) {
			suf[j] = v
		} else {
			suf = append(suf, v)
		}
		ans = max(ans, j+1+k+sort.SearchInts(pre, -v+1))
	}
	Fprint(out, ans)
}

//func main() { p8776(bufio.NewReader(os.Stdin), os.Stdout) }
