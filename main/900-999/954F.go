package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type matrix54 [][]int64

func newMatrix54(n, m int) matrix54 {
	a := make(matrix54, n)
	for i := range a {
		a[i] = make([]int64, m)
	}
	return a
}

func newIdentityMatrix54(n int) matrix54 {
	a := make(matrix54, n)
	for i := range a {
		a[i] = make([]int64, n)
		a[i][i] = 1
	}
	return a
}

func (a matrix54) mul(b matrix54) matrix54 {
	c := newMatrix54(len(a), len(b[0]))
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				c[i][j] = (c[i][j] + v*b[k][j]) % (1e9 + 7)
			}
		}
	}
	return c
}

func CF954F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, row int
	var m, l, r int64
	Fscan(in, &n, &m)
	type pair struct {
		row   int
		col   int64
		delta int8
	}
	block := make([]pair, 0, n*2+2)
	block = append(block, pair{0, 2, 0}, pair{0, m + 1, 0})
	for ; n > 0; n-- {
		Fscan(in, &row, &l, &r)
		row--
		block = append(block, pair{row, l, 1}, pair{row, r + 1, -1})
	}
	sort.Slice(block, func(i, j int) bool { return block[i].col < block[j].col })

	ans := newIdentityMatrix54(3)
	cntB := [3]int{}
	for i, b := range block[:len(block)-1] {
		cntB[b.row] += int(b.delta)
		m := newMatrix54(3, 3)
		if cntB[0] == 0 {
			m[0][0] = 1
			m[1][0] = 1
		}
		if cntB[1] == 0 {
			m[0][1] = 1
			m[1][1] = 1
			m[2][1] = 1
		}
		if cntB[2] == 0 {
			m[1][2] = 1
			m[2][2] = 1
		}
		for n := block[i+1].col - b.col; n > 0; n >>= 1 {
			if n&1 > 0 {
				ans = ans.mul(m)
			}
			m = m.mul(m)
		}
	}
	Fprint(out, ans[1][1])
}

//func main() { CF954F(os.Stdin, os.Stdout) }
