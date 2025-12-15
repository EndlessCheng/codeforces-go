package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type sparseTable9[T any] struct {
	st [][]T
	op func(T, T) T
}

func newSparseTable9[T any](nums []T, op func(T, T) T) sparseTable9[T] {
	n := len(nums)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], nums)
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = op(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return sparseTable9[T]{st, op}
}

// [l,r) 0-index
func (s sparseTable9[T]) query(l, r int) T {
	k := bits.Len(uint(r-l)) - 1
	return s.op(s.st[k][l], s.st[k][r-1<<k])
}

func cf1709D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, r1, c1, r2, c2, k int
	Fscan(in, &n, &m)
	a := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := newSparseTable9(a, func(a, b int) int { return max(a, b) })

	Fscan(in, &q)
	for range q {
		Fscan(in, &r1, &c1, &r2, &c2, &k)
		if (r1-r2)%k == 0 && (c1-c2)%k == 0 && t.query(min(c1, c2)-1, max(c1, c2)) < n-(n-r1)%k {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1709D(bufio.NewReader(os.Stdin), os.Stdout) }
