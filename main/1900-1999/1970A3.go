package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1970A3(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var s string
	Fscan(in, &s)
	n := len(s)
	g := make([][]int, n)
	c0, c1 := 0, 0
	for i := range n {
		if s[i] == '(' {
			c1++
			g[c0] = append(g[c0], c1)
		} else {
			c0++
		}
	}

	var dfs func(int)
	dfs = func(x int) {
		slices.Reverse(g[x])
		for _, y := range g[x] {
			Fprint(out, "(")
			dfs(y)
			Fprint(out, ")")
		}
	}
	dfs(0)
}

//func main() { cf1970A3(bufio.NewReader(os.Stdin), os.Stdout) }
