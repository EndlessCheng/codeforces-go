package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1000C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var l, r, prev int64
	Fscan(in, &n)
	a := make([]int64, 0, n*2)
	for i := 0; i < n; i++ {
		Fscan(in, &l, &r)
		a = append(a, l<<1|1, (r+1)<<1)
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	ans := make([]int64, n+1)
	for i, c := 0, 0; i < n*2; {
		v := a[i] >> 1
		ans[c] += v - prev
		for ; i < n*2 && a[i]>>1 == v; i++ {
			c += int(a[i]&1)<<1 - 1
		}
		prev = v
	}
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF1000C(os.Stdin, os.Stdout) }
