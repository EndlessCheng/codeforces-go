package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// https://github.com/EndlessCheng
type st74 [][17]int

func newST74(a []int) st74 {
	n := len(a)
	st := make(st74, n)
	for i, v := range a {
		st[i][0] = v
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = gcd74(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}

func (st st74) query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	return gcd74(st[l][k], st[r-1<<k][k])
}

func cf474F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, l, r int
	Fscan(in, &n)
	a := make([]int, n)
	pos := map[int][]int{}
	for i := range a {
		Fscan(in, &a[i])
		pos[a[i]] = append(pos[a[i]], i)
	}

	t := newST74(a)
	Fscan(in, &q)
	for range q {
		Fscan(in, &l, &r)
		l--
		p := pos[t.query(l, r)]
		Fprintln(out, r-l-sort.SearchInts(p, r)+sort.SearchInts(p, l))
	}
}

//func main() { cf474F(bufio.NewReader(os.Stdin), os.Stdout) }
func gcd74(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
