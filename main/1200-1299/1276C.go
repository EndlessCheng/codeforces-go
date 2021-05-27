package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1276C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, mxH, mxW int
	Fscan(in, &n)
	cnt := map[int]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		cnt[v]++
	}
	type pair struct{ c, v int }
	cv := make([]pair, 0, len(cnt))
	for v, c := range cnt {
		cv = append(cv, pair{c, v})
	}
	sort.Slice(cv, func(i, j int) bool { return cv[i].c < cv[j].c })

	for h, s, i := 1, 0, 0; ; h++ {
		s += len(cv) - i
		w := s / h
		if w < h {
			break
		}
		if h*w > mxH*mxW {
			mxH, mxW = h, w
		}
		for ; i < len(cv) && cv[i].c <= h; i++ {
		}
	}

	ans := make([][]int, mxH)
	for i := range ans {
		ans[i] = make([]int, mxW)
	}
	for rc, i := 0, len(cv)-1; rc < mxH*mxW; i-- {
		c, v := cv[i].c, cv[i].v
		if c > mxH {
			c = mxH
		}
		for ; c > 0; c-- {
			col, row := rc/mxH, rc%mxH
			ans[row][(col+row)%mxW] = v
			rc++
		}
	}
	Fprintln(out, mxH*mxW)
	Fprintln(out, mxH, mxW)
	for _, r := range ans {
		for _, v := range r {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1276C(os.Stdin, os.Stdout) }
