package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1209D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n+1)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }
	same := func(x, y int) bool { return find(x) == find(y) }

	var n, m, v, w, ans int
	Fscan(in, &n, &m)
	initFa(n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		if same(v, w) {
			ans++
		} else {
			merge(v, w)
		}
	}
	Fprint(out, ans)
}

//func main() {
//	Sol1209D(os.Stdin, os.Stdout)
//}
