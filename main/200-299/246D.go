package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF246D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w int
	Fscan(in, &n, &m)
	ans := int(1e9)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < ans {
			ans = a[i]
		}
	}
	cnt := [1e5 + 1]map[int]bool{}
	for i := range cnt[:] {
		cnt[i] = map[int]bool{}
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v, w = a[v-1], a[w-1]
		cnt[v][w] = true
		cnt[w][v] = true
	}
	for i, c := range cnt[:] {
		delete(c, i)
	}
	for i, c := range cnt[:] {
		if len(c) > len(cnt[ans]) {
			ans = i
		}
	}
	Fprint(out, ans)
}

//func main() { CF246D(os.Stdin, os.Stdout) }
