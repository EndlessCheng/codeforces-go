package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1019C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for range m {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
	}

	conn := make([]bool, n)
	for v, nb := range g {
		for _, w := range nb {
			if !conn[v] && v < w {
				conn[w] = true
			}
		}
	}
	for v := n - 1; v >= 0; v-- {
		for _, w := range g[v] {
			if !conn[v] {
				conn[w] = true
			}
		}
	}

	ans := []int{}
	for i, b := range conn {
		if !b {
			ans = append(ans, i+1)
		}
	}
	Fprintln(out, len(ans))
	for _, x := range ans {
		Fprint(out, x, " ")
	}
}

//func main() { cf1019C(bufio.NewReader(os.Stdin), os.Stdout) }
