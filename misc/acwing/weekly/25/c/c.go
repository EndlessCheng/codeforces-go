package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, l, r, save int
	Fscan(in, &n, &m, &l)
	c := make([]int, n)
	for i := range c {
		Fscan(in, &c[i])
	}
	uf := newUnionFind(n)
	for ; m > 0; m-- {
		Fscan(in, &l, &r)
		uf.merge(l-1, r-1) // 用并查集将询问合并
	}

	groups := make([][]int, n)
	for i, c := range c {
		i = uf.find(i)
		groups[i] = append(groups[i], c) // 统计同一组询问的气球颜色
	}

	for _, cs := range groups {
		// 组与组之间互不影响，对于每一组，保留该组中出现次数最多的气球颜色
		maxCnt := 0
		cnt := map[int]int{}
		for _, c := range cs {
			if cnt[c]++; cnt[c] > maxCnt {
				maxCnt = cnt[c]
			}
		}
		save += maxCnt
	}
	Fprint(out, n-save) // 剩下即为要染色的
}

func main() { run(os.Stdin, os.Stdout) }

type uf struct {
	fa []int
}

func newUnionFind(n int) uf {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return uf{fa}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u uf) merge(from, to int) {
	x, y := u.find(from), u.find(to)
	if x != y {
		u.fa[x] = y
	}
}
