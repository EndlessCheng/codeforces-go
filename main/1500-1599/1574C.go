package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1574C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	var n, m int
	var tot, def, atk int64
	Fscan(in, &n)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		tot += a[i]
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &def, &atk)
		i := sort.Search(n, func(i int) bool { return a[i] >= def })
		res := int64(2e18)
		if i < n {
			res = atk - (tot - a[i])
		}
		if i > 0 {
			if cost := def - a[i-1] + max(atk-(tot-a[i-1]), 0); cost < res {
				res = cost
			}
		}
		Fprintln(out, max(res, 0))
	}
}

//func main() { CF1574C(os.Stdin, os.Stdout) }
