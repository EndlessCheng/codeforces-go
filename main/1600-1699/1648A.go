package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1648A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, ans int
	Fscan(in, &n, &m)
	pos := [100001][2][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fscan(in, &v)
			pos[v][0] = append(pos[v][0], i)
			pos[v][1] = append(pos[v][1], j)
		}
	}
	f := func(p []int) {
		sj := 0
		for i, j := range p {
			ans += i*j - sj
			sj += j
		}
	}
	for _, p := range pos {
		f(p[0])
		sort.Ints(p[1])
		f(p[1])
	}
	Fprint(out, ans)
}

//func main() { CF1648A(os.Stdin, os.Stdout) }
