package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF762A(in io.Reader, out io.Writer) {
	var n int64
	var k int
	Fscan(in, &n, &k)
	ds := []int64{1}
	for i := int64(2); i*i <= n; i++ {
		for p, d := i, ds; n%i == 0; n /= i {
			for _, d := range d {
				ds = append(ds, d*p)
			}
			p *= i
		}
	}
	if n > 1 {
		for _, d := range ds {
			ds = append(ds, d*n)
		}
	}
	sort.Slice(ds, func(i, j int) bool { return ds[i] < ds[j] })
	if k > len(ds) {
		Fprint(out, -1)
	} else {
		Fprint(out, ds[k-1])
	}
}

//func main() { CF762A(os.Stdin, os.Stdout) }
