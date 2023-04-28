package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1808D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	pos := [2e5 + 1][2][]int{}
	for i := range a {
		Fscan(in, &a[i])
		pos[a[i]][i%2] = append(pos[a[i]][i%2], i)
	}
	ans := int64(0)
	for i, v := range a[:n-1-k/2] {
		ps := pos[v][i%2]
		l := max(i-k/2+1, 0)
		r := min(i+k-1, n-1)
		ans += int64(r - l - k + 2 - sort.SearchInts(ps, (r*2-k+1-i)+1) + sort.SearchInts(ps, l*2+k-1-i))
	}
	Fprint(out, ans)
}

//func main() { CF1808D(os.Stdin, os.Stdout) }
