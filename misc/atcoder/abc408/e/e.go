package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	es := make([]struct{ x, y, w int }, m)
	for i := range es {
		Fscan(in, &es[i].x, &es[i].y, &es[i].w)
	}

	fa := make([]int, n+1)
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	ans := 0
	for i := 29; i >= 0; i-- {
		for j := range fa {
			fa[j] = j
		}
		tar := ans >> i
		for _, e := range es {
			if e.w>>i|tar == tar {
				fa[find(e.x)] = find(e.y)
			}
		}
		if find(1) != find(n) {
			ans |= 1 << i
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
