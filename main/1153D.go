package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1153D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, leafCnt int
	Fscan(in, &n)
	tp := make([]int8, n)
	for i := range tp {
		Fscan(in, &tp[i])
	}
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &v)
		g[v-1] = append(g[v-1], w)
	}

	var f func(int) int
	f = func(v int) (cnt int) {
		if g[v] == nil {
			leafCnt++
			tp[v] = 1
			return 1
		}
		if tp[v] == 0 {
			for _, w := range g[v] {
				cnt += f(w)
			}
		} else {
			cnt = 1e9
			for _, w := range g[v] {
				if c := f(w); c < cnt {
					cnt = c
				}
			}
		}
		return
	}
	Fprint(_w, leafCnt+1-f(0))
}

//func main() { CF1153D(os.Stdin, os.Stdout) }
