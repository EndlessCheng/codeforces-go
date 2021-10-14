package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1539C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var k, x int64
	Fscan(in, &n, &k, &x)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	gap := []int64{}
	for i := 1; i < n; i++ {
		if d := a[i] - a[i-1]; d > x {
			gap = append(gap, d)
		}
	}
	sort.Slice(gap, func(i, j int) bool { return gap[i] < gap[j] })

	for len(gap) > 0 {
		c := (gap[0] - 1) / x
		if c > k {
			break
		}
		k -= c
		gap = gap[1:]
	}
	Fprint(out, len(gap)+1)
}

//func main() { CF1539C(os.Stdin, os.Stdout) }
