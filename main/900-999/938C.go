package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF938C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, v int
	Fscan(in, &q)
	qs := map[int][]int{}
	ans := make([][2]int, q)
	for i := range ans {
		if Fscan(in, &v); v > 0 {
			qs[v] = append(qs[v], i)
		} else {
			ans[i] = [2]int{1, 1}
		}
	}

	for n := 1; n*n-(n/2)*(n/2) <= 1e9; n++ {
		for l := 2; l <= n; {
			h := n / l
			v := n*n - h*h
			if v > 1e9 {
				break
			}
			for _, i := range qs[v] {
				ans[i] = [2]int{n, l}
			}
			l = n/h + 1
		}
	}
	for _, p := range ans {
		if p[0] > 0 {
			Fprintln(out, p[0], p[1])
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { CF938C(os.Stdin, os.Stdout) }
