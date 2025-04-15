package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
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
	lcmSize := make([]int, m+1)
	tar := 1
	for i := 1; i <= m; i++ {
		for j := i; j <= m; j += i {
			lcmSize[j] += len(ps[i])
			if lcmSize[j] > lcmSize[tar] {
				tar = j
			}
		}
	}
	ids := []int{}
	for d := 1; d*d <= tar; d++ {
		if tar%d == 0 {
			ids = append(ids, ps[d]...)
			if d*d < tar {
				ids = append(ids, ps[tar/d]...)
			}
		}
	}
	slices.Sort(ids)
	Fprintln(out, tar, len(ids))
	for _, v := range ids {
		Fprint(out, v, " ")
	}
}

//func main() { cf632D(bufio.NewReader(os.Stdin), os.Stdout) }
