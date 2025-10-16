package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// 我的题解 https://www.luogu.com.cn/article/kdouydty

// https://github.com/EndlessCheng
func cf1028H(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 5032108
	lpf := [mx]int{}
	for i := 2; i < mx; i++ {
		if lpf[i] == 0 {
			for j := i; j < mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	var n, q int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type query struct{ l, i int }
	qs := make([][]query, n)
	for i := range q {
		var l, r int
		Fscan(in, &l, &r)
		qs[r-1] = append(qs[r-1], query{l, i})
	}

	ans := make([]int, q)
	const mxW = 7
	opToI := make([]int, mxW*2+1)
	maxI := [mx][mxW + 1]int{}
	ps := []int{}
	mul := [1 << mxW]int{1}
	for i, v := range a {
		ps = ps[:0]
		for v > 1 {
			p := lpf[v]
			e := 1
			for v /= p; v%p == 0; v /= p {
				e ^= 1
			}
			if e > 0 {
				ps = append(ps, p)
			}
		}
		w := len(ps)

		for w2 := range mxW + 1 {
			op := w + w2
			opToI[op] = max(opToI[op], maxI[1][w2])
		}
		maxI[1][w] = i + 1

		for j, p := range ps {
			b := 1 << j
			for k, m := range mul[:b] {
				m *= p
				mul[b|k] = m

				common := bits.OnesCount8(uint8(b | k))
				for w2 := common; w2 <= mxW; w2++ {
					op := w + w2 - common*2
					opToI[op] = max(opToI[op], maxI[m][w2])
				}
				maxI[m][w] = i + 1
			}
		}

		for _, p := range qs[i] {
			for op, j := range opToI {
				if j >= p.l {
					ans[p.i] = op
					break
				}
			}
		}
	}

	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf1028H(bufio.NewReader(os.Stdin), os.Stdout) }
