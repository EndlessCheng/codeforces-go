package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF425A(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	sum := make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		sum[i+1] = sum[i] + a[i]
	}
	ans := a[0]
	for r, s := range sum {
		for l, sl := range sum[:r] {
			b := append([]int{}, a[l:r]...)
			sort.Ints(b)
			c := append(append(sort.IntSlice{}, a[:l]...), a[r:]...)
			sort.Sort(sort.Reverse(c))
			s := s - sl
			for i := 0; i < k && i < len(b) && i < len(c) && b[i] < c[i]; i++ {
				s += c[i] - b[i]
			}
			if s > ans {
				ans = s
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF425A(os.Stdin, os.Stdout) }
