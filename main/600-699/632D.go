package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf632D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, v int
	Fscan(in, &n, &m)
	ps := make([][]int, m+1)
	for i := 1; i <= n; i++ {
		if Fscan(in, &v); v <= m {
			ps[v] = append(ps[v], i)
		}
	}
	sz, mxL := make([]int, m+1), 1
	for i := 1; i <= m; i++ {
		for j := i; j <= m; j += i {
			if sz[j] += len(ps[i]); sz[j] > sz[mxL] {
				mxL = j
			}
		}
	}
	ids := []int{}
	for d := 1; d*d <= mxL; d++ {
		if mxL%d == 0 {
			ids = append(ids, ps[d]...)
			if d*d < mxL {
				ids = append(ids, ps[mxL/d]...)
			}
		}
	}
	sort.Ints(ids)
	Fprintln(out, mxL, len(ids))
	for _, v := range ids {
		Fprint(out, v, " ")
	}
}

//func main() { cf632D(bufio.NewReader(os.Stdin), os.Stdout) }
