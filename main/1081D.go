package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1081D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, v int
	Fscan(in, &n, &m, &k)
	sp := make([]bool, n)
	for i := 0; i < k; i++ {
		Fscan(in, &v)
		sp[v-1] = true
	}
	es := make([][3]int, m)
	for i := range es {
		Fscan(in, &es[i][0], &es[i][1], &es[i][2])
	}
	sort.Slice(es, func(i, j int) bool { return es[i][2] < es[j][2] })

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	cnt := 1
	merge := func(x, y int) {
		x, y = find(x), find(y)
		if x == y {
			return
		}
		if sp[x] && sp[y] {
			cnt++
			fa[y] = x
		} else if sp[x] {
			fa[y] = x
		} else {
			fa[x] = y
		}
	}
	for i := 0; ; {
		w := es[i][2]
		for ; i < m && es[i][2] == w; i++ {
			merge(es[i][0]-1, es[i][1]-1)
		}
		if cnt == k {
			Fprint(_w, strings.Repeat(strconv.Itoa(w)+" ", k))
			break
		}
	}
}

//func main() { CF1081D(os.Stdin, os.Stdout) }
