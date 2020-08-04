package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF954D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, start, end, ans int
	Fscan(in, &n, &m, &start, &end)
	start--
	end--
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	bfs := func(st int) []int {
		depths := make([]int, n)
		vis := make([]bool, n)
		vis[st] = true
		queue := []int{st}
		for len(queue) > 0 {
			v, queue = queue[0], queue[1:]
			for _, w := range g[v] {
				if !vis[w] {
					vis[w] = true
					depths[w] = depths[v] + 1
					queue = append(queue, w)
				}
			}
		}
		return depths
	}
	stDepths, edDepths := bfs(start), bfs(end)
	minDis := stDepths[end]
	for v, gi := range g {
		sort.Ints(gi)
		j := 0
		for w := range g {
			if j == len(gi) || w < gi[j] {
				if w > v && stDepths[v]+edDepths[w]+1 >= minDis && stDepths[w]+edDepths[v]+1 >= minDis {
					ans++
				}
			} else {
				j++
			}
		}
	}
	Fprint(out, ans)
}

//func main() {
//	CF954D(os.Stdin, os.Stdout)
//}
