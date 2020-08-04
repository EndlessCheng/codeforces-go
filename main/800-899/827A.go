package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF827A(_r io.Reader, _w io.Writer) {
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n)
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
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	ans := make([]byte, 2e6-1)
	end := 0
	initFa(2e6)
	var n, k, st int
	var s string
	for Fscan(in, &n); n > 0; n-- {
		for Fscan(in, &s, &k); k > 0; k-- {
			Fscan(in, &st)
			st--
			e := st + len(s)
			if e > end {
				end = e
			}
			for i := find(st); i < e; i = find(i + 1) {
				ans[i] = s[i-st]
				fa[i] = e
			}
		}
	}
	ans = ans[:end]
	for i, c := range ans {
		if c == 0 {
			ans[i] = 'a'
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() {
//	CF827A(os.Stdin, os.Stdout)
//}
