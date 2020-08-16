package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1276B(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	r := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}
	subSize := func(g [][]int, a, b int) (sz int) {
		ids := make([]int, len(g))
		idCnt := 0
		var f func(int)
		f = func(v int) {
			ids[v] = idCnt
			for _, w := range g[v] {
				if w != a && ids[w] == 0 {
					f(w)
				}
			}
		}
		for i, id := range ids {
			if i != a && id == 0 {
				idCnt++
				f(i)
			}
		}
		if idCnt == 1 {
			return
		}
		for i, id := range ids {
			if i != a && id != ids[b] {
				sz++
			}
		}
		return
	}

	for t := r(); t > 0; t-- {
		n, m, a, b := r(), r(), r()-1, r()-1
		g := make([][]int, n)
		for ; m > 0; m-- {
			v, w := r()-1, r()-1
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		Fprintln(out, int64(subSize(g, a, b))*int64(subSize(g, b, a)))
	}
}

//func main() { CF1276B(os.Stdin, os.Stdout) }
