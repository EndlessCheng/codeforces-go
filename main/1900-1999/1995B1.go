package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1995B1(in io.Reader, out io.Writer) {
	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		cnt := map[int]int{}
		for range n {
			Fscan(in, &v)
			cnt[v]++
		}

		minLeft := m
		for x, c1 := range cnt {
			c2 := cnt[x+1]
			k1 := min(m/x, c1)
			left := m - x*k1 // 先买价值为 x 的物品
			k2 := min(left/(x+1), c2)
			left -= (x + 1) * k2   // 再买价值为 x+1 的物品
			left -= min(k1, c2-k2) // 最后把价值为 x 的物品替换成价值为 x+1 的物品
			minLeft = min(minLeft, max(left, 0))
		}
		Fprintln(out, m-minLeft)
	}
}

//func main() { cf1995B1(bufio.NewReader(os.Stdin), os.Stdout) }
