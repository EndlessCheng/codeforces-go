package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf67D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	pos := make([]int, n+1)
	for i := n; i > 0; i-- {
		Fscan(in, &v)
		pos[v] = i
	}
	g := []int{}
	for ; n > 0; n-- {
		Fscan(in, &v)
		v = pos[v]
		j := sort.SearchInts(g, v)
		if j < len(g) {
			g[j] = v
		} else {
			g = append(g, v)
		}
	}
	Fprint(out, len(g))
}

//func main() { cf67D(os.Stdin, os.Stdout) }
