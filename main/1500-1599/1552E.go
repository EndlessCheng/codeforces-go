package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1552E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, v, c int
	Fscan(in, &n, &k)
	lim := (n-1)/(k-1) + 1
	ans := make([][]int, n)
	inAns := make([]bool, n)
	pos := make([][]int, n)
	for i := 1; i <= n*k; i++ {
		Fscan(in, &v)
		v--
		if inAns[v] {
			continue
		}
		pos[v] = append(pos[v], i)
		if len(pos[v]) != 2 {
			continue
		}
		if c++; c == lim {
			c = 0
			for j, p := range pos {
				if len(p) > 1 {
					ans[j] = p[:2]
					inAns[j] = true
				}
				pos[j] = nil
			}
		}
	}
	for j, p := range pos {
		if len(p) > 1 {
			ans[j] = p[:2]
		}
	}
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

//func main() { CF1552E(os.Stdin, os.Stdout) }
