package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf551E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r, v int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	type block struct {
		l, r, todo int
		lp, rp     map[int]int
	}
	blockSize := int(math.Ceil(math.Sqrt(float64(n) / 5)))
	blocks := make([]*block, (n-1)/blockSize+1)
	calc := func(l, r int) (map[int]int, map[int]int) {
		lp := map[int]int{}
		for j := r - 1; j >= l; j-- {
			lp[a[j]] = j
		}
		rp := map[int]int{}
		for j := l; j < r; j++ {
			rp[a[j]] = j
		}
		return lp, rp
	}
	for i := 0; i < n; i += blockSize {
		r := min(i+blockSize, n)
		lp, rp := calc(i, r)
		blocks[i/blockSize] = &block{i, r, 0, lp, rp}
	}

	for range q {
		Fscan(in, &op, &l)
		if op == 1 {
			Fscan(in, &r, &v)
			l--
			for _, b := range blocks {
				if b.r <= l {
					continue
				}
				if b.l >= r {
					break
				}
				if l <= b.l && b.r <= r {
					b.todo += v
					continue
				}
				for j := b.l; j < b.r; j++ {
					a[j] += b.todo
					if l <= j && j < r {
						a[j] += v
					}
				}
				b.todo = 0
				b.lp, b.rp = calc(b.l, b.r)
			}
		} else {
			posL := n
			for _, b := range blocks {
				if i, ok := b.lp[l-b.todo]; ok {
					posL = i
					break
				}
			}
			if posL == n {
				Fprintln(out, -1)
				continue
			}
			for i := len(blocks) - 1; ; i-- {
				b := blocks[i]
				if j, ok := b.rp[l-b.todo]; ok {
					Fprintln(out, j-posL)
					break
				}
			}
		}
	}
}

//func main() { cf551E(bufio.NewReader(os.Stdin), os.Stdout) }
