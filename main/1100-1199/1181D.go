package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1181D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, k int64
	Fscan(in, &n, &m, &q)
	a := make([]int64, n)
	c := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
		a[i]--
		c[a[i]]++
		a[i] += int64(c[a[i]]-1) * m
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	for i := range a {
		a[i] -= int64(i) // 减去比举办时间小的城市个数，排序后直接减去下标即可
	}
	for ; q > 0; q-- {
		Fscan(in, &k)
		k += int64(sort.Search(int(n), func(i int) bool { return a[i] >= k-n })) - n
		k %= m
		if k == 0 {
			k = m
		}
		Fprintln(out, k)
	}
}

//func main() { CF1181D(os.Stdin, os.Stdout) }
