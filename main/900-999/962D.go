package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// 堆的写法见 https://codeforces.com/problemset/submission/962/141866417

// github.com/EndlessCheng/codeforces-go
func CF962D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var v int64
	Fscan(in, &n)
	pos := map[int64]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		for pos[v] > 0 {
			delete(pos, v)
			v *= 2
		}
		pos[v] = i
	}
	a := make([]int64, 0, len(pos))
	for v := range pos {
		a = append(a, v)
	}
	sort.Slice(a, func(i, j int) bool { return pos[a[i]] < pos[a[j]] })
	Fprintln(out, len(a))
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF962D(os.Stdin, os.Stdout) }
