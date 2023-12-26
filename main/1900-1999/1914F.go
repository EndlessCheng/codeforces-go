package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1914F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for w := 1; w < n; w++ {
			Fscan(in, &v)
			v--
			g[v] = append(g[v], w)
		}

		size := make([]int, n)
		var dfs func(int)
		dfs = func(x int) {
			size[x] = 1
			for i, y := range g[x] {
				dfs(y)
				size[x] += size[y]
				if size[y] > size[g[x][0]] {
					g[x][0], g[x][i] = g[x][i], g[x][0]
				}
			}
		}
		dfs(0)

		ans := 0
		other := 0
		x := 0
		for {
			if other > 0 {
				ans++ // 其它点和 v 匹配
				other--
			}
			if len(g[x]) == 0 {
				break
			}

			s := size[x] - 1
			y := g[x][0]

			// 最大子树大小 <= 其它点个数
			// 祖先节点已经在上面判断了， other 是不含祖先节点的
			if size[y]*2 <= s+other {
				ans += (s + other) / 2
				break
			}

			other += s - size[y]
			x = y
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1914F(os.Stdin, os.Stdout) }
