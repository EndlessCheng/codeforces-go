package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1547G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]int, n)
		for ; m > 0; m-- {
			var v, w int
			Fscan(in, &v, &w)
			g[v-1] = append(g[v-1], w-1)
		}

		ans := make([]int, n)
		inStack := make([]bool, n)
		var dfs func(int, bool)
		dfs = func(v int, inCycle bool) {
			inStack[v] = true
			if inCycle {
				ans[v] = -1
			} else {
				// 首次访问：ans[v] = 1
				// 再次访问：ans[v] = 2
				ans[v]++
			}
			for _, w := range g[v] {
				if ans[w] < 0 {
					continue
				}
				if inCycle || inStack[w] { // w 在环上，再访问一次（w 至多访问三次）
					dfs(w, true) // 从 w 出发能到达的点都在环上
				} else if ans[w] < 2 { // 除非后面发现 w 在环上，否则至多访问 w 两次
					dfs(w, false)
				}
			}
			inStack[v] = false
		}
		dfs(0, false)
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1547G(os.Stdin, os.Stdout) }
