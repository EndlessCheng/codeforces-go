package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, r, up int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type pair struct{ up, i int }
	qs := make([][]pair, n)
	for i := 0; i < q; i++ {
		Fscan(in, &r, &up)
		qs[r-1] = append(qs[r-1], pair{up, i})
	}

	ans := make([]int, q)
	f := []int{}
	for i, v := range a {
		j := sort.SearchInts(f, v)
		if j < len(f) {
			f[j] = v
		} else {
			f = append(f, v)
		}
		for _, p := range qs[i] {
			ans[p.i] = sort.SearchInts(f, p.up+1)
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
