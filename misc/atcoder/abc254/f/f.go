package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type sparseTable [][]int

func newSparseTable(a []int) sparseTable {
	n := len(a)
	if n == 0 {
		return nil
	}
	w := bits.Len(uint(n))
	st := make([][]int, w)
	for i := range st {
		st[i] = make([]int, n)
	}
	copy(st[0], a)
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = gcd(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return st
}

// [l, r) 下标从 0 开始
func (s sparseTable) query(l, r int) int {
	if l >= r {
		return 0
	}
	k := bits.Len(uint(r-l)) - 1
	return gcd(s[k][l], s[k][r-1<<k])
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, la, ra, lb, rb int
	Fscan(in, &n, &q)
	f := func() ([]int, sparseTable) {
		a := make([]int, n)
		d := make([]int, n-1)
		for i := range a {
			Fscan(in, &a[i])
			if i > 0 {
				d[i-1] = abs(a[i] - a[i-1])
			}
		}
		return a, newSparseTable(d)
	}
	a, sta := f()
	b, stb := f()

	for range q {
		Fscan(in, &la, &ra, &lb, &rb)
		Fprintln(out, gcd(a[la-1]+b[lb-1], gcd(sta.query(la-1, ra-1), stb.query(lb-1, rb-1))))
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
func abs(x int) int { if x < 0 { return -x }; return x }
func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
