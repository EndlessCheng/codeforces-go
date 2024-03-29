package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1037D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n+1)
	g[1] = []int{0}
	for m := n - 1; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	g2 := make([][]int, n+1)
	g2[1] = []int{0}
	Fscan(in, &v)
	for queue := []int{v}; len(queue) > 0; {
		v, queue = queue[0], queue[1:]
		for m := len(g[v]) - 1; m > 0; m-- {
			Fscan(in, &w)
			g2[v] = append(g2[v], w)
			g2[w] = append(g2[w], v)
			queue = append(queue, w)
		}
	}

	// 判断 g 和 g2 是否同构
	for i, e := range g {
		e2 := g2[i]
		if len(e) != len(e2) {
			Fprint(out, "No")
			return
		}
		sort.Ints(e)
		sort.Ints(e2)
		for j, w := range e {
			if w != g2[i][j] {
				Fprint(out, "No")
				return
			}
		}
	}
	Fprint(out, "Yes")
}

//func main() {
//	Sol1037D(os.Stdin, os.Stdout)
//}
