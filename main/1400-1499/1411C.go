package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1411C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		fa := make([]int, n+1)
		for i := range fa {
			fa[i] = i
		}
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				fa[x] = find(fa[x])
			}
			return fa[x]
		}

		ans := 0
		for ; m > 0; m-- {
			Fscan(in, &x, &y)
			if x == y {
				continue
			}
			x = find(x)
			y = find(y)
			if x != y {
				fa[x] = y
				ans++
			} else {
				ans += 2
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1411C(os.Stdin, os.Stdout) }
