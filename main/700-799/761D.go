package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf761D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, l, r, v int
	Fscan(in, &n, &l, &r)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		pos[v-1] = i
	}

	pre := int(-2e9)
	for _, p := range pos {
		pre = max(pre+1, l-a[p])
		a[p] += pre
		if a[p] > r {
			Fprint(out, -1)
			return
		}
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { cf761D(bufio.NewReader(os.Stdin), os.Stdout) }
