package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, r, q int
	Fscan(in, &n, &m)
	ls := make([][]int, n+1)
	rs := make([][]int, n+1)
	for range m {
		Fscan(in, &l, &r)
		ls[r] = append(ls[r], l)
		rs[l] = append(rs[l], r)
	}

	maxL2 := make([]int, n+1)
	mx, mx2 := 0, 0
	for r, l := range ls {
		slices.Sort(l)
		for _, l := range l {
			if l > mx {
				mx2 = mx
				mx = l
			} else if l > mx2 {
				mx2 = l
			}
		}
		maxL2[r] = mx2
		slices.Sort(rs[r])
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &l, &r)
		i := sort.SearchInts(rs[l], r+1) - 1
		if i < 0 {
			Fprintln(out, "No")
			continue
		}
		r1 := rs[l][i]
		if r1 == r {
			if maxL2[r] >= l {
				Fprintln(out, "Yes")
			} else {
				Fprintln(out, "No")
			}
		} else {
			j := sort.SearchInts(ls[r], l)
			if j < len(ls[r]) && ls[r][j] <= r1+1 {
				Fprintln(out, "Yes")
			} else {
				Fprintln(out, "No")
			}
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
