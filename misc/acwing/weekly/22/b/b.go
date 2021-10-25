package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, x, y int
	Fscan(in, &n, &k)
	r := make([]int, n)
	rs := make([][]int, n)
	for i := range r {
		Fscan(in, &r[i])
	}
	for ; k > 0; k-- {
		Fscan(in, &x, &y)
		x--
		y--
		rs[x] = append(rs[x], r[y]) // 存储每个战士的矛盾战士的战斗力
		rs[y] = append(rs[y], r[x])
	}

	// 排序
	for _, r := range rs {
		sort.Ints(r)
	}
	sortedR := append([]int(nil), r...)
	sort.Ints(sortedR)

	// 二分得到小于该战士战斗力的战士个数，以及其中与该战士矛盾的战士个数，相减即为答案
	for i, v := range r {
		Fprint(out, sort.SearchInts(sortedR, v)-sort.SearchInts(rs[i], v), " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
