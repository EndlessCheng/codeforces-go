package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, r, c, ans int
	Fscan(in, &n, &n, &n)
	type pair struct{ r, c int }
	grid := make(map[pair]int, n)
	rowS := map[int]int{}
	colS := map[int]int{}
	for ; n > 0; n-- {
		Fscan(in, &r, &c)
		grid[pair{r, c}] = 1
		rowS[r]++
		colS[c]++
	}

	type cs struct{ c, s int }
	colList := make([]cs, 0, len(colS))
	for x, s := range colS {
		colList = append(colList, cs{x, s})
	}
	sort.Slice(colList, func(i, j int) bool { return colList[i].s > colList[j].s })

	for r, s := range rowS {
		for _, cs := range colList {
			v, ok := grid[pair{r, cs.c}] // 每个点至多访问一次
			if !ok {
				ans = max(ans, s+cs.s)
				break // 保证时间复杂度是 O(n) 的关键：colList 从大到小排序，后面的只会更小，无需遍历
			}
			ans = max(ans, s+cs.s-v)
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a }
