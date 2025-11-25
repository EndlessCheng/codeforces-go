package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf383B(in io.Reader, out io.Writer) {
	var sz, n int
	Fscan(in, &sz, &n)
	type point struct{ row, col int }
	a := make([]point, n)
	for i := range a {
		Fscan(in, &a[i].row, &a[i].col)
	}
	slices.SortFunc(a, func(a, b point) int { return cmp.Or(a.row-b.row, a.col-b.col) })

	type pair struct{ l, r int }
	ps := []pair{{1, 1}}
	preRow := 0
	for i := 0; i < n; {
		if len(ps) > 0 && a[i].row > preRow+1 {
			ps = []pair{{ps[0].l, sz}}
		}

		st := i
		row := a[i].row
		for i++; i < n && a[i].row == row; i++ {
		}

		newPs := []pair{}
		k, l := 0, 1
		for j := st; j <= i; j++ {
			r := sz
			if j < i {
				r = a[j].col - 1
			}
			for ; k < len(ps) && ps[k].r < l; k++ {
			}
			if k < len(ps) && max(ps[k].l, l) <= r {
				newPs = append(newPs, pair{max(ps[k].l, l), r})
			}
			l = r + 2
		}

		ps = newPs
		preRow = row
	}

	if len(ps) > 0 && (preRow < sz || ps[len(ps)-1].r == sz) {
		Fprint(out, sz*2-2)
	} else {
		Fprint(out, -1)
	}
}

//func main() { cf383B(bufio.NewReader(os.Stdin), os.Stdout) }
